package segmentfault

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestCreateQuestion(t *testing.T) {
	client, err := NewClient(os.Getenv(TestEnvSegmentfaultToken))
	if err != nil {
		t.Fatal(err)
	}

	d := &Draft{
		Tags:  []int64{1040000000311191}, // Rust
		Type:  DraftTypeQuestion,
		Text:  "Rust async 用法",
		Title: fmt.Sprintf("Rust async"),
	}
	err = client.createDraft(d)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.CreateQuestion(d)
	if err != nil {
		t.Fatal(err)
	}
	s, _ := json.Marshal(resp)
	fmt.Printf("%s\n", s)
}

func TestListQuestions(t *testing.T) {
	client, err := NewClient(os.Getenv(TestEnvSegmentfaultToken))
	if err != nil {
		t.Fatal(err)
	}

	opts := &ListOptions{}
	resp, err := client.ListQuestions(opts, true)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("pagination: %+v", resp.Pagination)
	for _, row := range resp.Rows {
		t.Logf("%+v", row)
	}

	opts = &ListOptions{
		Query: QueryTypeNewest,
	}
}

func TestDeleteQuestion(t *testing.T) {
	client, err := NewClient(os.Getenv(TestEnvSegmentfaultToken))
	if err != nil {
		t.Fatal(err)
	}

	err = client.DeleteQuestion(1010000041881484)
	if err != nil {
		t.Fatal(err)
	}
}
