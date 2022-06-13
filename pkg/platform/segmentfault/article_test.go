package segmentfault

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestListArticles(t *testing.T) {
	client, err := NewClient(os.Getenv(TestEnvSegmentfaultToken))
	if err != nil {
		t.Fatal(err)
	}
	opts := &ListArticlesOptions{
		Page: 1,
		Size: 10,
		Sort: ArticleSortNewest,
	}
	articles, err := client.ListArticles(opts)
	if err != nil {
		t.Fatal(err)
	}
	s, _ := json.Marshal(articles)
	fmt.Printf("%s\n", s)
}
