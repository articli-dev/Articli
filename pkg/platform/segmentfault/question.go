package segmentfault

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/juju/errors"
)

type QuestionData struct {
	ID               int64  `json:"id"`
	UserID           int64  `json:"user_id"`
	AcceptedAnswerID int64  `json:"accepted_answer_id"`
	LastAnswerID     int64  `json:"last_answer_id"`
	RevisionID       int64  `json:"revision_id"`
	SiteID           int64  `json:"site_id"`
	CategoryID       int64  `json:"category_id"`
	IsEdited         int    `json:"is_edited"`
	IsSite           int    `json:"is_site"`
	Status           int    `json:"status"`
	Answers          int    `json:"answers"`
	Comments         int    `json:"comments"`
	Votes            int    `json:"votes"`
	Views            int    `json:"views"`
	UniqueViews      int    `json:"unique_views"`
	Title            string `json:"title"`
	OriginalText     string `json:"original_text"`
	ParsedText       string `json:"parsed_text"`
	TagsList         string `json:"tags_list"`
	Created          int64  `json:"created"`
	Modified         int64  `json:"modified"`
}

type CreateQuestionRequest struct {
	DraftID int64   `json:"draft_id"`
	Title   string  `json:"title"`
	Text    string  `json:"text"`
	Log     string  `json:"log"`
	Tags    []int64 `json:"tags"`
}

type CreateQuestionResponse struct {
	Data    *QuestionData `json:"data"`
	Message string        `json:"msg"`
}

func (c *Client) CreateQuestion(d *Draft) (resp *CreateQuestionResponse, err error) {
	endpoint := "/question"
	req := d.IntoCreateQuestionRequest()
	err = c.Request(http.MethodPost, endpoint, nil, req, &resp)
	err = errors.Trace(err)
	return
}

func (c *Client) DeleteQuestion(id int64) error {
	endpoint := "/question/" + strconv.FormatInt(id, 10)
	err := c.Request(http.MethodDelete, endpoint, nil, nil, nil)
	return errors.Trace(err)
}

type QuestionRow struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	Votes        int    `json:"votes"`
	Comments     int    `json:"comments"`
	Created      int64  `json:"created"`
	IsAccepted   bool   `json:"is_accepted"`
	StatusString string `json:"status_string"`
	RealViews    int    `json:"real_views"`
}

func (r *QuestionRow) GetURL() string {
	return BuildURL(r.URL)
}

type ListQuestionsResponse struct {
	Pagination
	Rows []*QuestionRow `json:"rows"`
}

func (c *Client) ListQuestions(opts *ListOptions, owned bool) (resp *ListQuestionsResponse, err error) {
	var endpoint string
	if owned {
		endpoint = "/homepage/" + c.User.Slug + "/questions"
	} else {
		endpoint = "/questions"
	}

	var params url.Values
	if opts != nil {
		params = opts.IntoParams()
	}
	err = c.Get(endpoint, params, &resp)
	err = errors.Trace(err)
	return
}
