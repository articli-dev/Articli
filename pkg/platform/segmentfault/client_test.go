package segmentfault

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient(os.Getenv(TestEnvSegmentfaultToken))
	if err != nil {
		t.Fatal(err)
	}
	s, _ := json.Marshal(client)
	fmt.Printf("%s\n", s)
}
