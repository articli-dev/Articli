package segmentfault

import (
	"os"
	"testing"
)

func TestListNotes(t *testing.T) {
	client, err := NewClient(os.Getenv(TestEnvSegmentfaultToken))
	if err != nil {
		t.Fatal(err)
	}

	opts := &ListNotesOptions{
		Page: 1,
		Size: 10,
		Sort: NoteSortCreated,
		Q:    "test",
	}
	resp, err := client.ListNotes(opts)
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range resp.Rows {
		t.Logf("%+v", row)
	}
}

func TestUpdateNotePublic(t *testing.T) {
	client, err := NewClient(os.Getenv(TestEnvSegmentfaultToken))
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		noteID   int64
		isPublic bool
	}{
		{1330000041880250, true},
		{1330000041880250, false},
	}
	for _, c := range cases {
		err = client.UpdateNotePublic(c.noteID, c.isPublic)
		if err != nil {
			t.Fatal(err)
		}
	}
}
