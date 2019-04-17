package durations

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestWeekof(t *testing.T) {

	// notice that the date is the 31st of Dec, 2014, but the Monday of that date's week
	// is the 29th but both of those dates are considered the first week of 2015 because 4
	// of the 7 days of that week (beginning on Monday) fall in the year 2015. In other words
	// the first week of the year is the week of the first Thursday of the year.
	ts, e := time.Parse("2006-01-02", "2014-12-31")
	if e != nil {
		t.Errorf(e.Error())
	}

	ts = Weekof(ts)

	if diff := cmp.Diff("2014-12-29 Monday", ts.Format("2006-01-02 Monday")); diff != "" {
		t.Errorf("Format: (+want/-got)\n%s", diff)
	}

	y, w := ts.ISOWeek()
	got := fmt.Sprintf(FmtYearWeek, y, w)

	if diff := cmp.Diff("201501", got); diff != "" {
		t.Errorf("year: (+want/-got)\n%s", diff)
	}

}
