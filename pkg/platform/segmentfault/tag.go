package segmentfault

import (
	"net/http"
	"net/url"

	"github.com/juju/errors"
)

type Tag struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	IconURL string `json:"icon_url"`
	URL     string `json:"url"`
}

func (t *Tag) GetURL() string {
	return DefaultSiteURL + t.URL
}

type ListTagsResponse struct {
	Size int               `json:"size"`
	Rows map[string][]*Tag `json:"rows"`
}

func (c *Client) ListTags() (resp *ListTagsResponse, err error) {
	endpoint := "/tags"
	err = c.Request(http.MethodGet, endpoint, nil, nil, &resp)
	err = errors.Trace(err)
	return
}

type SearchTagRow struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	IconURL      string `json:"icon_url"`
	Excerpt      string `json:"excerpt"`
	ThumbnailURL string `json:"thumbnail_url"`
	IsFollowed   bool   `json:"is_followed"`
	ContentCount int    `json:"content_count"`
}

func (t *SearchTagRow) GetURL() string {
	return DefaultSiteURL + t.URL
}

type SearchTagsResponse struct {
	Rows []*SearchTagRow `json:"rows"`
	Size int             `json:"size"`
}

func (c *Client) SearchTags(q string) (resp *SearchTagsResponse, err error) {
	endpoint := "/tags"
	params := url.Values{
		"query": {string(QueryTypeSearch)},
		"q":     {q},
	}
	err = c.Get(endpoint, params, &resp)
	err = errors.Trace(err)
	return
}
