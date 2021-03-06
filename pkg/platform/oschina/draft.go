package oschina

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/google/go-querystring/query"
	"github.com/juju/errors"
	"github.com/tidwall/gjson"
	"net/url"
	"path/filepath"
	"strings"
)

// SaveDraft create a new draft if id is empty, otherwise update draft
func (c *Client) SaveDraft(params *ContentParams) error {
	if err := params.Validate(); err != nil {
		return errors.Trace(err)
	}

	rawurl := c.BuildURL("/blog/save_draft")
	values, err := query.Values(params)
	if err != nil {
		return errors.Trace(err)
	}
	raw, err := c.Post(rawurl, values, DefaultHandler)
	if err != nil {
		return errors.Trace(err)
	}
	if params.DraftID == "" {
		params.DraftID = gjson.Get(raw, "result.draft").String()
		if params.DraftID == "" {
			return errors.New("failed to get draft id")
		}
	}
	return nil
}

func (c *Client) DeleteDraft(id string) error {
	rawurl := c.BuildURL("/blog/delete_draft")
	values := url.Values{
		"id": {id},
	}
	_, err := c.Post(rawurl, values, DefaultHandler)
	return errors.Trace(err)
}

type Draft struct {
	ID    string
	Title string
	URL   string
}

func (c *Client) ListDrafts(page int) (drafts []*Draft, hasNext bool, err error) {
	if page < 1 {
		page = 1
	}
	path := fmt.Sprintf("%s%s", c.BaseURL, fmt.Sprintf("/admin/drafts?p=%d", page))
	raw, err := c.Get(path, nil, nil)
	if err != nil {
		err = errors.Trace(err)
		return
	}
	doc, err := htmlquery.Parse(strings.NewReader(raw))
	if err != nil {
		err = errors.Trace(err)
		return
	}
	q := `//div[@class="ui relaxed divided items list-container"]//a[@class="header"]`
	nodes, err := htmlquery.QueryAll(doc, q)
	if err != nil {
		err = errors.Trace(err)
		return
	}

	if len(nodes) == 0 {
		return
	}

	for _, node := range nodes {
		draft := &Draft{
			Title: node.FirstChild.Data,
		}
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				draft.URL = attr.Val
				draft.ID = filepath.Base(attr.Val)
			}
		}
		drafts = append(drafts, draft)
	}

	q = `//div[@class="ui pagination menu "]/a[@class="item next-item"]`
	nodes, err = htmlquery.QueryAll(doc, q)
	if err != nil {
		err = errors.Trace(err)
		return
	}
	hasNext = len(nodes) > 0
	return
}

func (c *Client) GetDraftDetail(id string) (*ContentParams, error) {
	path := fmt.Sprintf("%s%s", c.BaseURL, fmt.Sprintf("/blog/write/draft/%s", id))
	raw, err := c.Get(path, nil, nil)
	if err != nil {
		return nil, errors.Trace(err)
	}

	doc, err := htmlquery.Parse(strings.NewReader(raw))
	if err != nil {
		return nil, errors.Trace(err)
	}

	result := new(ContentParams)

	// Title
	q := `//form[@class="ui write-article form"]//input[@name="title" and @type="text"]`
	nodes, err := htmlquery.QueryAll(doc, q)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if len(nodes) != 1 {
		return nil, errors.New("found more none or than one title node")
	}
	node := nodes[0]
	for _, attr := range node.Attr {
		if attr.Key == "value" {
			result.Title = attr.Val
		}
	}

	// Content
	q = `//form[@class="ui write-article form"]//textarea[@name="body"]`
	nodes, err = htmlquery.QueryAll(doc, q)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if len(nodes) != 1 {
		return nil, errors.New("found more none or than one body node")
	}
	result.Content = nodes[0].LastChild.Data

	// ????????????
	q = `//form[@class="ui write-article form"]//input[@name="origin_url" and @type="text"]`
	nodes, err = htmlquery.QueryAll(doc, q)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if len(nodes) != 1 {
		return nil, errors.New("found more none or than one origin_url node")
	}
	node = nodes[0]
	for _, attr := range node.Attr {
		if attr.Key == "value" {
			result.OriginalURL = attr.Val
		}
	}

	// ???????????????
	q = `//form[@class="ui write-article form"]//input[@name="privacy" and @type="checkbox"]`
	nodes, err = htmlquery.QueryAll(doc, q)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if len(nodes) != 1 {
		return nil, errors.New("found more none or than one privacy node")
	}
	node = nodes[0]
	for _, attr := range node.Attr {
		if attr.Key == "value" {
			if attr.Val == "1" {
				result.Privacy = 1
			}
		}
	}

	// ??????
	q = `//form[@class="ui write-article form"]//input[@name="as_top" and @type="checkbox"]`
	nodes, err = htmlquery.QueryAll(doc, q)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if len(nodes) != 1 {
		return nil, errors.New("found more none or than one as_top node")
	}
	node = nodes[0]
	for _, attr := range node.Attr {
		if attr.Key == "value" {
			if attr.Val == "1" {
				result.Top = 1
			}
		}
	}

	// ????????????
	q = `//form[@class="ui write-article form"]//input[@name="deny_comment" and @type="checkbox"]`
	nodes, err = htmlquery.QueryAll(doc, q)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if len(nodes) != 1 {
		return nil, errors.New("found more none or than one deny_comment node")
	}
	node = nodes[0]
	for _, attr := range node.Attr {
		if attr.Key == "value" {
			if attr.Val == "1" {
				result.DenyComment = 1
			}
		}
	}

	// ???????????????????????????
	q = `//form[@class="ui write-article form"]//input[@name="downloadImg" and @type="checkbox"]`
	nodes, err = htmlquery.QueryAll(doc, q)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if len(nodes) != 1 {
		return nil, errors.New("found more none or than one downloadImg node")
	}
	node = nodes[0]
	for _, attr := range node.Attr {
		if attr.Key == "value" {
			if attr.Val == "1" {
				result.DownloadImage = 1
			}
		}
	}

	// ????????????
	q = `//form[@class="ui write-article form"]//input[@name="type"]`
	nodes, err = htmlquery.QueryAll(doc, q)
	if err != nil {
		return nil, errors.Trace(err)
	}
	checked := false
	for _, node := range nodes {
		for _, attr := range node.Attr {
			if attr.Key == "checked" {
				checked = true
			}
			if attr.Key == "value" {
				result.Type = ArticleType(attr.Val)
			}
		}
		if checked {
			break
		}
	}

	// ???????????? Category
	q = `//select[@id="catalogDropdown"]/option`
	nodes, err = htmlquery.QueryAll(doc, q)
	if err != nil {
		return nil, errors.Trace(err)
	}
	for _, node := range nodes {
		for _, attr := range node.Attr {
			if attr.Key == "selected" {
				categoryName := node.LastChild.Data
				category, err := c.GetCategoryByName(categoryName)
				if err != nil {
					return nil, errors.Trace(err)
				}
				if category != nil {
					result.Category = category.ID
				}
			}
		}
	}
	return result, nil
}

func (c *Client) PublishDraft(id string) (articleID string, err error) {
	var params *ContentParams
	params, err = c.GetDraftDetail(id)
	if err != nil {
		err = errors.Trace(err)
		return
	}
	err = c.SaveArticle(params)
	err = errors.Trace(err)
	return
}
