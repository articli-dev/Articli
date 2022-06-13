package segmentfault

import (
	"os"
	"testing"
)

func TestListAnswers(t *testing.T) {
	client, err := NewClient(os.Getenv(TestEnvSegmentfaultToken))
	if err != nil {
		t.Fatal(err)
	}

	opts := &ListAnswersOptions{
		Page: 1,
		Size: 10,
		Sort: AnswerSortNewest,
	}
	resp, err := client.ListAnswers(opts)
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range resp.Answers {
		t.Logf("%+v", row)
	}
}
