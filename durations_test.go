package durations

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDurations(t *testing.T) {

	real := Hours(1)

	if diff := cmp.Diff(int64(3600000000000), real.Nanoseconds()); diff != "" {
		t.Errorf("Nanoseconds: (+want/-got)\n%s", diff)
	}

	if diff := cmp.Diff(float64(3600), real.Seconds()); diff != "" {
		t.Errorf("Seconds: (+want/-got)\n%s", diff)
	}

	if diff := cmp.Diff(float64(60), real.Minutes()); diff != "" {
		t.Errorf("Minutes: (+want/-got)\n%s", diff)
	}

	if diff := cmp.Diff(float64(1), real.Hours()); diff != "" {
		t.Errorf("Hours: (+want/-got)\n%s", diff)
	}

}
