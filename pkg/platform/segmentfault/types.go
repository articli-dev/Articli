package segmentfault

import (
	"net/url"
	"strconv"
)

type Pagination struct {
	Page      int `json:"page"`
	Size      int `json:"size"`
	TotalPage int `json:"total_page"`
}

type SortType string

const (
	SortTypeNewest   SortType = "newest"
	SortTypeVotes    SortType = "votes"
	SortTypeModified SortType = "modified"
	SortTypeCreated  SortType = "created"
)

type QueryType string

const (
	QueryTypeMine         QueryType = "mine"
	QueryTypeSearch       QueryType = "search"
	QueryTypeNewest       QueryType = "newest"
	QueryTypeRecommend    QueryType = "recommend"
	QueryTypeUnanswered   QueryType = "unanswered"
	QueryTypeFollowing    QueryType = "following"
	QueryTypeDailyHottest QueryType = "daily_hottest"
	QueryTypeWeekHottest  QueryType = "week_hottest"
	QueryTypeMonthHottest QueryType = "month_hottest"
)

type ObjectType string

const (
	ObjectTypeArticle  ObjectType = "article"
	ObjectTypeNote     ObjectType = "note"
	ObjectTypeQuestion ObjectType = "question"
	ObjectTypeAnswer   ObjectType = "answer"
)

type ListOptions struct {
	ObjectType ObjectType
	Page       int
	Size       int
	Sort       SortType
	Query      QueryType
	Q          string
}

const (
	PageSizeMax = 20
	PageSizeMin = 3

	PageSizeMaxDraft = 20
	PageSizeMinDraft = 3
)

func (opts *ListOptions) IntoParams() url.Values {
	values := make(url.Values)
	if opts.Page < 1 {
		opts.Page = 1
	}
	values.Set("page", strconv.Itoa(opts.Page))

	if opts.Size < PageSizeMin || opts.Size > PageSizeMax {
		opts.Size = PageSizeMax
	}
	values.Set("size", strconv.Itoa(opts.Size))

	if opts.Sort != "" {
		values.Set("sort", string(opts.Sort))
	}
	if opts.Q != "" {
		values.Set("q", opts.Q)
	}

	switch opts.ObjectType {
	case ObjectTypeNote:
		values.Set("query", string(QueryTypeMine))
	}
	return values
}
