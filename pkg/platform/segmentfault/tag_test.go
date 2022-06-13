package segmentfault

import (
	"os"
	"testing"
)

func TestListTags(t *testing.T) {
	client, err := NewClient(os.Getenv(TestEnvSegmentfaultToken))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.ListTags()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("size: %d", resp.Size)
	for category, tags := range resp.Rows {
		t.Logf("%s tags count: %d\n", category, len(tags))
		for _, tag := range tags {
			t.Logf("category: %s, tag: %+v", category, tag)
		}
	}
}

func TestSearchTags(t *testing.T) {
	client, err := NewClient(os.Getenv(TestEnvSegmentfaultToken))
	if err != nil {
		t.Fatal(err)
	}
	q := "go"
	resp, err := client.SearchTags(q)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("size: %d", resp.Size)
	for _, row := range resp.Rows {
		t.Logf("%+v", row)
	}
}
