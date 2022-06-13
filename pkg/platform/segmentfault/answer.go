package segmentfault

import (
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/juju/errors"
)

type Answer struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	Votes      int    `json:"votes"`
	Comments   int    `json:"comments"`
	IsAccepted bool   `json:"is_accepted"`
	Created    int64  `json:"created"`
}

type ListAnswersResponse struct {
	Pagination
	Answers []*Answer `json:"rows"`
}

type ListAnswersOptions struct {
	Size int        `url:"size,omitempty"`
	Page int        `url:"page,omitempty"`
	Sort AnswerSort `url:"sort,omitempty"`
}

type AnswerSort string

const (
	AnswerSortNewest AnswerSort = "newest"
	AnswerSortVotes  AnswerSort = "votes"
)

func (c *Client) ListAnswers(opts *ListAnswersOptions) (resp *ListAnswersResponse, err error) {
	var params url.Values
	params, err = query.Values(opts)
	if err != nil {
		err = errors.Trace(err)
		return
	}

	endpoint := "/homepage/" + c.User.Slug + "/answers"
	err = c.Get(endpoint, params, &resp)
	err = errors.Trace(err)
	return
}
