package durations

import (
	"fmt"
	"time"
)

// FmtYearWeek is the format string used by fmt.Println to create a 6 digit string of the yearweek ('YYYYWW').
// e.g. `y, w := t.ISOWeek(); fmt.Sprintf(FmtYearWeek, y, w)`
const FmtYearWeek = "%d%02d"

// YearWeek give you the YearWeek (YYYYWW) of the given date
func YearWeek(date time.Time) string {
	tmpY, tmpW := date.ISOWeek()
	return fmt.Sprintf(FmtYearWeek, tmpY, tmpW)
}
