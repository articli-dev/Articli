package segmentfault

import (
	"net/http"

	"github.com/juju/errors"

	"github.com/k8scat/articli/pkg/utils"
)

type UploadResponse struct {
	URL string `json:"url"`
}

func (c *Client) UploadImage(file string) (string, error) {
	form := utils.NewForm()
	form.SetFile("image", file)
	buf, contentType, err := form.Encode()
	if err != nil {
		return "", errors.Trace(err)
	}

	endpoint := "/image"
	req, err := c.NewRequest(http.MethodPost, endpoint, nil, buf)
	if err != nil {
		return "", errors.Trace(err)
	}
	req.Header.Set("Content-Type", contentType)

	var resp *UploadResponse
	if err = c.Do(req, &resp); err != nil {
		return "", errors.Trace(err)
	}
	return DefaultSiteURL + resp.URL, nil
}
