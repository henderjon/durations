package durations

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInterval(t *testing.T) {

	mockPingChan := make(chan string)

	go NewRunner(Seconds(3), []Runner{
		func() {
			mockPingChan <- "BOOM"
			close(mockPingChan)
		},
	})

	var mockPingResult strings.Builder
	for s := range mockPingChan {
		// log.Println("interval")
		mockPingResult.WriteString(s)
	}

	if diff := cmp.Diff(mockPingResult.String(), "BOOM"); diff != "" {
		t.Errorf("cron: (-got +want)\n%s", diff)
	}

}
