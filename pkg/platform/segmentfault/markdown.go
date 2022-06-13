package segmentfault

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/juju/errors"

	"github.com/k8scat/articli/pkg/markdown"
)

type SaveType string

const (
	SaveTypeArticle SaveType = "article"
	SaveTypeDraft   SaveType = "draft"
)

type MarkdownMeta struct {
	ArticleID  string
	DraftID    string
	Title      string
	Brief      string
	Content    string
	CoverImage string
	TagIDs     []int64
}

func (m *MarkdownMeta) IntoDraft() (*Draft, error) {
	draftID, err := strconv.ParseInt(m.DraftID, 10, 64)
	if err != nil {
		return nil, errors.Trace(err)
	}
	objectID, err := strconv.ParseInt(m.ArticleID, 10, 64)
	if err != nil {
		return nil, errors.Trace(err)
	}
	d := &Draft{
		Title:    m.Title,
		Text:     m.Content,
		Type:     DraftTypeArticle,
		ID:       draftID,
		ObjectID: objectID,
		Cover:    m.CoverImage,
		Tags:     m.TagIDs,
	}
	return d, nil
}

func ParseMark(mark *markdown.Mark) (m *MarkdownMeta, err error) {
	v := mark.Meta.Get("segmentfault")
	if v == nil {
		err = errors.New("segmentfault meta not found")
		return
	}
	meta, ok := v.(markdown.Meta)
	if !ok {
		err = errors.New("segmentfault meta not found")
		return
	}

	m = new(MarkdownMeta)
	m.Title = meta.GetString("title")
	if m.Title == "" {
		m.Title = mark.Meta.GetString("title")
		if m.Title == "" {
			err = errors.New("title is required")
			return
		}
	}

	m.CoverImage = meta.GetString("cover_image")
	if m.CoverImage == "" {
		coverImages := mark.Meta.GetStringSlice("cover_images")
		if len(coverImages) > 0 {
			m.CoverImage = coverImages[0]
		}
	}

	m.Content = mark.Content
	m.Brief = meta.GetString("brief_content")
	if m.Brief == "" {
		m.Brief = mark.Brief
	}
	briefContentLen := len([]rune(m.Brief))
	if briefContentLen > 100 {
		s := compressContent(m.Brief)
		m.Brief = string([]rune(s)[:80])
	} else if briefContentLen < 50 {
		s := compressContent(m.Content)
		m.Brief = string([]rune(s)[:80])
	}

	prefixContent := meta.GetString("prefix_content")
	if prefixContent != "" {
		m.Content = fmt.Sprintf("%s\n\n%s", prefixContent, m.Content)
	}
	suffixContent := meta.GetString("suffix_content")
	if suffixContent != "" {
		m.Content = fmt.Sprintf("%s\n\n%s", m.Content, suffixContent)
	}

	tags := meta.GetStringSlice("tags")
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}
		parts := strings.Split(tag, ",")
		id, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return nil, errors.Trace(err)
		}
		m.TagIDs = append(m.TagIDs, id)
	}

	m.ArticleID = meta.GetString("article_id")
	m.DraftID = meta.GetString("draft_id")
	return
}

func WriteMarkdownMeta(saveType SaveType, mark *markdown.Mark, m *MarkdownMeta, isCreate bool) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	v := mark.Meta.Get("segmentfault")
	meta, _ := v.(markdown.Meta)
	if isCreate {
		meta = meta.Set(fmt.Sprintf("%s_create_time", saveType), now)
	} else {
		meta = meta.Set(fmt.Sprintf("%s_update_time", saveType), now)
	}

	if m.DraftID != "" {
		meta = meta.Set("draft_id", m.DraftID)
	}
	if m.ArticleID != "" {
		meta = meta.Set("article_id", m.ArticleID)
	}
	mark.Meta = mark.Meta.Set("segmentfault", meta)
	err := mark.WriteFile(mark.File)
	return errors.Trace(err)
}

func compressContent(s string) string {
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	return s
}
