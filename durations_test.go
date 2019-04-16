package durations

import (
	"fmt"
	"testing"
	"time"

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

func TestWeekof(t *testing.T) {

	// notice that the date is the 31st of Dec, 2014, but the Monday of that date's week
	// is the 29th but those dates are considered the first week of 2015.
	ts, e := time.Parse("2006-01-02", "2014-12-31")
	if e != nil {
		t.Errorf(e.Error())
	}

	ts = Weekof(ts)

	if diff := cmp.Diff("2014-12-29 Monday", ts.Format("2006-01-02 Monday")); diff != "" {
		t.Errorf("Format: (+want/-got)\n%s", diff)
	}

	y, w := ts.ISOWeek()
	got := fmt.Sprintf("%d%02d", y, w)

	if diff := cmp.Diff("201501", got); diff != "" {
		t.Errorf("year: (+want/-got)\n%s", diff)
	}

}
