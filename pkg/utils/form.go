package utils

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
)

// https://stackoverflow.com/questions/20205796/post-data-using-the-content-type-multipart-form-data
type Form struct {
	data        map[string]io.Reader
	orderedKeys []string
}

func NewForm() *Form {
	return &Form{
		data:        make(map[string]io.Reader),
		orderedKeys: make([]string, 0),
	}
}

func (f *Form) Set(key string, value io.Reader) {
	f.data[key] = value
	f.orderedKeys = append(f.orderedKeys, key)
}

func (f *Form) SetString(key, value string) {
	f.data[key] = strings.NewReader(value)
	f.orderedKeys = append(f.orderedKeys, key)
}

func (f *Form) SetFile(key, filepath string) {
	f.data[key] = MustOpen(filepath)
	f.orderedKeys = append(f.orderedKeys, key)
}

func (f *Form) Get(key string) io.Reader {
	return f.data[key]
}

// Encode encodes the form into a multipart/form-data body.
func (f *Form) Encode() (buf *bytes.Buffer, contentType string, err error) {
	buf = new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	for _, k := range f.orderedKeys {
		r := f.data[k]
		var fw io.Writer
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(k, x.Name()); err != nil {
				err = errors.Trace(err)
				return
			}
		} else {
			if fw, err = w.CreateFormField(k); err != nil {
				err = errors.Trace(err)
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			err = errors.Trace(err)
			return
		}
		if x, ok := r.(io.Closer); ok {
			x.Close()
		}
	}
	// 显式关闭
	w.Close()
	contentType = w.FormDataContentType()
	return
}

func MustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	return r
}
