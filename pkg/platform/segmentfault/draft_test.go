package segmentfault

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestDeleteDraft(t *testing.T) {
	client, err := NewClient(os.Getenv(TestEnvSegmentfaultToken))
	if err != nil {
		t.Fatal(err)
	}
	err = client.DeleteDraft(1220000041875865)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateDraft(t *testing.T) {
	client, err := NewClient(os.Getenv(TestEnvSegmentfaultToken))
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 20; i++ {
		d := &Draft{
			Text:  "123",
			Title: fmt.Sprintf("title-%d", i),
			Type:  DraftTypeArticle,
		}
		err = client.createDraft(d)
		if err != nil {
			t.Fatal(err)
		}
		s, _ := json.Marshal(d)
		fmt.Printf("%s\n", s)
		time.Sleep(5)
	}
}
