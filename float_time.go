package durations

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type FloatTime struct {
	Precision time.Duration
	time.Time
}

func NewFloatTime() *FloatTime {
	return &FloatTime{
		Precision: time.Second,
		Time:      time.Now().UTC(),
	}
}

func (f *FloatTime) MarshalJSON() ([]byte, error) {
	var i int64
	switch true {
	case f.Precision == time.Nanosecond:
		i = f.UnixNano()
	case f.Precision == time.Microsecond:
		i = f.UnixMicro()
	case f.Precision == time.Millisecond:
		i = f.UnixMilli()
	default:
		i = f.Unix()
	}

	str := strconv.FormatInt(i, 10)
	return []byte(str), nil
}

func (f *FloatTime) UnmarshalJSON(b []byte) error {
	if string(b) == `null` {
		return nil
	}
	parts := strings.Split(string(b), ".")
	if len(parts) > 2 {
		return errors.New("invalid format, too many decimals")
	}

	if len(parts) == 1 {
		parts = append(parts, "0")
	}

	sec, e := strconv.ParseInt(parts[0], 10, 64)
	if e != nil {
		return fmt.Errorf("error parsing seconds value; %s", e)
	}

	nsec, e := strconv.ParseInt(parts[1], 10, 64)
	if e != nil {
		return fmt.Errorf("error parsing nanoseconds value; %s", e)
	}

	switch true {
	case f.Precision == time.Nanosecond:
		// nsec = nsec
	case f.Precision == time.Microsecond:
		nsec = nsec * int64(time.Microsecond)
	case f.Precision == time.Millisecond:
		nsec = nsec * int64(time.Millisecond)
	default:
		nsec = 0
	}

	t := time.Unix(sec, nsec)
	f.Time = t
	return nil
}

func (f *FloatTime) Unix() int64 {
	return f.Time.Unix()
}
