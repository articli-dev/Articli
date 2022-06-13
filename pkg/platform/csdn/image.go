package csdn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/juju/errors"

	"github.com/k8scat/articli/pkg/utils"
)

type UploadData struct {
	AccessID    string `json:"accessId"`
	CallbackURL string `json:"callbackUrl"`
	Dir         string `json:"dir"`
	Expire      string `json:"expire"`
	FilePath    string `json:"filePath"`
	Host        string `json:"host"`
	Policy      string `json:"policy"`
	Signature   string `json:"signature"`
}

type RequestUploadResponse struct {
	Data *UploadData `json:"data"`
	BaseResponse
}

func (c *Client) requestUpload(filename string) (*UploadData, error) {
	ext := strings.TrimLeft(filepath.Ext(filename), ".")
	if err := validateExt(ext); err != nil {
		return nil, errors.Trace(err)
	}

	rawurl := "https://imgservice.csdn.net/direct/v1.0/image/upload"
	req, err := http.NewRequest(http.MethodGet, rawurl, nil)
	if err != nil {
		return nil, errors.Trace(err)
	}

	query := make(url.Values)
	query.Set("watermark", "")
	query.Set("type", "blog")
	query.Set("rtype", "markdown")
	req.URL.RawQuery = query.Encode()

	req.Header.Set("Cookie", c.Cookie)
	req.Header.Set("x-image-app", "direct_blog")
	req.Header.Set("x-image-suffix", ext)
	req.Header.Set("x-image-dir", "direct")

	resp, err := c.Request(req)
	if err != nil {
		return nil, errors.Trace(err)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var result *RequestUploadResponse
	if err = json.Unmarshal(b, &result); err != nil {
		return nil, errors.Trace(err)
	}
	if result.Code != 200 {
		return nil, errors.New(result.Message)
	}
	return result.Data, nil
}

type UploadResponse struct {
	Data struct {
		ImageURL string `json:"imageUrl"`
	} `json:"data"`
	BaseResponse
}

// UploadImage uploads image to Aliyun OSS.
// 仅对文件名的后缀进行验证，不对真正的文件内容格式进行验证，所以可以将一个格式不支持的文件进行重命名即可上传
func (c *Client) UploadImage(path string) (string, error) {
	uploadData, err := c.requestUpload(path)
	if err != nil {
		return "", errors.Trace(err)
	}

	form := utils.NewForm()
	form.SetString("key", uploadData.FilePath)
	form.SetString("policy", uploadData.Policy)
	form.SetString("OSSAccessKeyId", uploadData.AccessID)
	form.SetString("success_action_status", "200")
	form.SetString("signature", uploadData.Signature)
	form.SetString("callback", uploadData.CallbackURL)
	form.SetFile("file", path)
	buf, contentType, err := form.Encode()
	if err != nil {
		return "", errors.Trace(err)
	}

	rawurl := uploadData.Host
	req, err := http.NewRequest(http.MethodPost, rawurl, buf)
	if err != nil {
		return "", errors.Trace(err)
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("User-Agent", browser.Computer())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.Trace(err)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Trace(err)
	}
	var result *UploadResponse
	if err = json.Unmarshal(b, &result); err != nil {
		return "", errors.Trace(err)
	}
	if result.Code != 200 {
		return "", errors.New(result.Message)
	}
	return result.Data.ImageURL, nil
}

func validateExt(ext string) error {
	switch strings.ToLower(ext) {
	case "jpg", "jpeg", "png", "gif":
		return nil
	default:
		return errors.New("invalid image suffix, only support jpg, jpeg, png, gif")
	}
}
