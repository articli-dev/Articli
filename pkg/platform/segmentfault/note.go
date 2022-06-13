package segmentfault

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/go-querystring/query"
	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
)

type NoteRow struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	CoverURL  string `json:"cover_url"`
	Excerpt   string `json:"excerpt"`
	Created   int64  `json:"created"`
	Modified  int64  `json:"modified"`
	IsSticky  int    `json:"is_sticky"`
	IsPrivate int    `json:"is_private"`
}

type ListNotesResponse struct {
	Pagination
	Rows []*NoteRow `json:"rows"`
}

type ListNotesOptions struct {
	Sort  NoteSort  `url:"sort,omitempty"`
	Page  int       `url:"page,omitempty"`
	Size  int       `url:"size,omitempty"`
	Query QueryType `url:"query"` // Always set to "mine"
}

func (o *ListNotesOptions) Validate() error {
	o.Query = QueryTypeMine
	return nil
}

type NoteSort string

const (
	NoteSortCreated  NoteSort = "create"
	NoteSortModified NoteSort = "modified"
)

func (c *Client) ListNotes(opts *ListNotesOptions) (resp *ListNotesResponse, err error) {
	if err = opts.Validate(); err != nil {
		err = errors.Trace(err)
		return
	}

	var params url.Values
	params, err = query.Values(opts)
	if err != nil {
		err = errors.Trace(err)
		return
	}
	endpoint := "/notes"
	err = c.Get(endpoint, params, &resp)
	err = errors.Trace(err)
	return
}

func (c *Client) UpdateNotePublic(noteID int64, isPublic bool) error {
	endpoint := "/note/" + strconv.FormatInt(noteID, 10) + "/public"
	payload := map[string]interface{}{
		"is_public": isPublic,
	}
	req, err := c.NewRequest(http.MethodPut, endpoint, nil, payload)
	if err != nil {
		return errors.Trace(err)
	}
	var result string
	err = c.Do(req, &result)
	log.WithField("result", result).Debug("update note public")
	return errors.Trace(err)
}

type CreateNoteResponse struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Language     string `json:"language"`
	Created      int64  `json:"created"`
	Modified     int64  `json:"modified"`
	IsBookmarked bool   `json:"is_bookmarked"`
	IsSticky     int    `json:"is_sticky"`
	IsPrivate    int    `json:"is_private"`
	User         User   `json:"user"`
	OriginalText string `json:"original_text"`
	ParsedText   string `json:"parsed_text"`
}

type CreateNoteRequest struct {
	Title    string `json:"title"`
	Text     string `json:"text"`
	Language string `json:"language"`
}

func (c *Client) CreateNote(req *CreateNoteRequest) (resp *CreateNoteResponse, err error) {
	endpoint := "/draft"
	err = c.Request(http.MethodPost, endpoint, nil, req, &resp)
	err = errors.Trace(err)
	return
}
