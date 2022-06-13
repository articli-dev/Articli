package segmentfault

import (
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/juju/errors"
)

type Blog struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnail_url"`
}

type ArticleRow struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Created     int64  `json:"created"`
	Votes       int    `json:"votes"`
	IsLinked    bool   `json:"is_linked"`
	Comments    int    `json:"comments"`
	Cover       string `json:"cover"`
	IsShowCover bool   `json:"is_show_cover"`
	RealViews   int    `json:"real_views"`
	User        *User  `json:"user"`
	Blog        *Blog  `json:"blog"`
}

func (a *ArticleRow) GetURL() string {
	return DefaultSiteURL + a.URL
}

type ListArticlesResponse struct {
	Pagination
	Rows []*ArticleRow `json:"rows"`
}

type ListArticlesOptions struct {
	Size int         `url:"size,omitempty"`
	Page int         `url:"page,omitempty"`
	Sort ArticleSort `url:"sort,omitempty"`
}

type ArticleSort string

const (
	ArticleSortNewest ArticleSort = "newest"
	ArticleSortVotes  ArticleSort = "votes"
)

func (c *Client) ListArticles(opts *ListArticlesOptions) (resp *ListArticlesResponse, err error) {
	var params url.Values
	params, err = query.Values(opts)
	if err != nil {
		err = errors.Trace(err)
		return
	}

	endpoint := "/homepage/" + c.User.Slug + "/articles"
	err = c.Get(endpoint, params, &resp)
	err = errors.Trace(err)
	return
}

type CreateArticleRequest struct {
	BlogID  int64       `json:"blog_id"`
	Cover   string      `json:"cover"`
	DraftID int64       `json:"draft_id"`
	License int         `json:"license,omitempty"` // 1 表示注明版权
	Log     string      `json:"log"`
	Tags    []int64     `json:"tags"`
	Text    string      `json:"text"`
	Title   string      `json:"title"`
	Type    ArticleType `json:"type"`
	URL     string      `json:"url"`
	Created string      `json:"created,omitempty"` // 定时发布，例如在 2022-05-26 00:00 发布：2022-05-25T16:00:00.000Z
}

func (req *CreateArticleRequest) ValidateWithOptions() error {
	if (req.Type == ArticleTypeRepost || req.Type == ArticleTypeTranslate) && req.URL == "" {
		return errors.New("url is required")
	}
	if len(req.Tags) == 0 {
		return errors.New("tags is required")
	}
	return nil
}

type ArticleType int

const (
	ArticleTypeOriginal  ArticleType = 1 // 原创
	ArticleTypeRepost    ArticleType = 2 // 转载
	ArticleTypeTranslate ArticleType = 3 // 翻译
)

type CreateArticleOptions struct {
	ArticleType ArticleType
	URL         string
	License     bool
	Created     *time.Time
}

type ArticleData struct {
	ID           int64       `json:"id"`
	BlogID       int64       `json:"blog_id"`
	UserID       int64       `json:"user_id"`
	RevisionID   int64       `json:"revision_id"`
	Title        string      `json:"title"`
	OriginalText string      `json:"original_text"`
	ParsedText   string      `json:"parsed_text"`
	TagsList     string      `json:"tags_list"`
	Status       int         `json:"status"`
	Votes        int         `json:"votes"`
	Views        int         `json:"views"`
	UniqueViews  int         `json:"unique_views"`
	Comments     int         `json:"comments"`
	ArticleType  ArticleType `json:"article_type"`
	OriginURL    string      `json:"origin_url"`
	Cover        string      `json:"cover"`
	HideSetting  string      `json:"hide_setting"`
	Created      int64       `json:"created"`
	Modified     int64       `json:"modified"`
	Recommended  int64       `json:"recommended"`
	IsArchived   int         `json:"is_archived"`
}

type CreateArticleResponse struct {
	Data    *ArticleData `json:"data"`
	Message string       `json:"msg"`
}

func (c *Client) CreateArticle(d *Draft, opts *CreateArticleOptions) (resp *CreateArticleResponse, err error) {
	endpoint := "/article"
	req := d.IntoCreateArticleRequest(opts)
	err = c.Request(http.MethodPost, endpoint, nil, req, &resp)
	err = errors.Trace(err)
	return
}

func (c *Client) DeleteArticle(id string) (err error) {
	endpoint := "/article/" + id
	err = c.Request(http.MethodDelete, endpoint, nil, nil, nil)
	err = errors.Trace(err)
	return
}
