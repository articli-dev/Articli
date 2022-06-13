package segmentfault

import (
	"os"
	"testing"
)

func TestUploadImage(t *testing.T) {
	client, err := NewClient(os.Getenv(TestEnvSegmentfaultToken))
	if err != nil {
		t.Fatal(err)
	}
	url, err := client.UploadImage("/Users/hsowan/workspace/articli/images/go.png")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
}
