package segmentfault

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/go-querystring/query"
	"github.com/juju/errors"
)

type Draft struct {
	// Common fields
	Language string    `json:"language,omitempty"`
	Text     string    `json:"text"`
	Title    string    `json:"title"`
	Type     DraftType `json:"type"`

	// 编辑已发布的内容
	ObjectID int64 `json:"object_id,omitempty"`

	// 提问/文章
	Tags []int64 `json:"tags,omitempty"`

	// 文章
	Cover string `json:"cover,omitempty"`

	// 更新时才有
	ID int64 `json:"id,omitempty"`
}

func (d *Draft) GetURL() string {
	if d.ObjectID == 0 {
		return fmt.Sprintf("https://segmentfault.com/draft/%d/edit", d.ID)
	}
	return fmt.Sprintf("https://segmentfault.com/a/%d/edit?draftId=%d", d.ObjectID, d.ID)
}

func (d *Draft) ValidateNote() error {
	d.Type = DraftTypeNote
	return nil
}

func (d *Draft) ValidateArticle() error {
	d.Type = DraftTypeArticle
	return nil
}

func (d *Draft) ValidateQuestion() error {
	d.Type = DraftTypeQuestion
	return nil
}

func (d *Draft) IntoCreateQuestionRequest() *CreateQuestionRequest {
	return &CreateQuestionRequest{
		DraftID: d.ID,
		Title:   d.Title,
		Text:    d.Text,
		Tags:    d.Tags,
	}
}

const TimeFormat = "2006-01-02T15:04:05.000Z"

func (d *Draft) IntoCreateArticleRequest(opts *CreateArticleOptions) *CreateArticleRequest {
	req := &CreateArticleRequest{
		DraftID: d.ID,
		Cover:   d.Cover,
		Text:    d.Text,
		Title:   d.Title,
		Tags:    d.Tags,
	}
	req.Type = opts.ArticleType
	req.URL = opts.URL
	if opts.License {
		req.License = 1
	}
	if opts.Created != nil {
		req.Created = opts.Created.UTC().Format(TimeFormat)
	}
	return req
}

func (d *Draft) IntoCreateNoteRequest() *CreateNoteRequest {
	return &CreateNoteRequest{
		Text:     d.Text,
		Title:    d.Title,
		Language: d.Language,
	}
}

type DraftType string

const (
	DraftTypeArticle  DraftType = "article"
	DraftTypeQuestion DraftType = "question"
	DraftTypeNote     DraftType = "note"
)

func (c *Client) DeleteDraft(id int64) error {
	endpoint := "/draft/" + strconv.FormatInt(id, 10)
	err := c.Request(http.MethodDelete, endpoint, nil, nil, nil)
	return errors.Trace(err)
}

type CreateDraftResponse struct {
	ID       int64 `json:"id"`
	ObjectID int64 `json:"object_id"`
}

func (c *Client) createDraft(d *Draft) error {
	endpoint := "/draft"
	var resp *CreateDraftResponse
	err := c.Request(http.MethodPost, endpoint, nil, d, &resp)
	if err != nil {
		return errors.Trace(err)
	}
	d.ID = resp.ID
	return nil
}

func (c *Client) CreateNoteDraft(d *Draft) error {
	err := d.ValidateNote()
	if err != nil {
		return errors.Trace(err)
	}
	err = c.createDraft(d)
	return errors.Trace(err)
}

func (c *Client) CreateArticleDraft(d *Draft) error {
	err := d.ValidateArticle()
	if err != nil {
		return errors.Trace(err)
	}
	err = c.createDraft(d)
	return errors.Trace(err)
}

func (c *Client) CreateQuestionDraft(d *Draft) error {
	err := d.ValidateQuestion()
	if err != nil {
		return errors.Trace(err)
	}
	err = c.createDraft(d)
	return errors.Trace(err)
}

func (c *Client) UpdateDraft(d *Draft) error {
	endpoint := "/draft/" + strconv.FormatInt(d.ID, 10)
	err := c.Request(http.MethodPut, endpoint, nil, d, nil)
	return errors.Trace(err)
}

type DraftRow struct {
	ID       int64  `json:"id"`
	ObjectID int64  `json:"object_id"`
	Title    string `json:"title"`
	Modified int64  `json:"modified"`
	TypeStr  string `json:"type_str"`
	TypeName string `json:"type_name"`
}

func (d *DraftRow) GetURL() string {
	switch d.TypeStr {
	case string(DraftTypeArticle):
		if d.ObjectID == 0 {
			return fmt.Sprintf("https://segmentfault.com/draft/%d/edit", d.ID)
		}
		return fmt.Sprintf("https://segmentfault.com/a/%d/edit?draftId=%d", d.ObjectID, d.ID)
	case string(DraftTypeQuestion):
		return fmt.Sprintf("https://segmentfault.com/ask?draftId=%d", d.ID)
	case string(DraftTypeNote):
		return fmt.Sprintf("https://segmentfault.com/record?draftId=%d", d.ID)
	default:
		return ""
	}
}

type ListDraftsResponse struct {
	Pagination
	Rows []*DraftRow `json:"rows"`
}

type ListDraftsOptions struct {
	Page int `url:"page,omitempty"`
}

func (c *Client) ListDrafts(opts *ListDraftsOptions) (resp *ListDraftsResponse, err error) {
	var params url.Values
	params, err = query.Values(opts)
	if err != nil {
		err = errors.Trace(err)
		return
	}

	endpoint := "/drafts/@me"
	err = c.Get(endpoint, params, &resp)
	err = errors.Trace(err)
	return
}
