package durations

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestYearWeek(t *testing.T) {

	expected := "201039"
	ts, _ := time.Parse("2006-01-02 15:04:05", "2010-10-01 11:03:34")
	actual := YearWeek(ts)
	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("Nanoseconds: (-got +want)\n%s", diff)
	}

}
