package durations

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

//2021-12-23T09:47:12-06:00
// fmt.Println(time.Unix(1640274432, 937853*1000).UnixMicro())

func TestFloatTimeUnmarshal(t *testing.T) {
	given := []byte(`1640274432.937853`)
	tm := &FloatTime{}
	tm.UnmarshalJSON(given)
	if diff := cmp.Diff(tm.Format(time.RFC3339), `2021-12-23T09:47:12-06:00`); diff != "" {
		t.Errorf("Nanosecond: (-got +want)\n%s", diff)
	}
}
func TestFloatTimeMarshal(t *testing.T) {
	f := &FloatTime{
		Time: time.Unix(1640274432, 937853000),
	}

	f.Precision = time.Nanosecond
	s, e := json.Marshal(f)
	if diff := cmp.Diff(string(s), `1640274432937853000`); diff != "" || e != nil {
		t.Errorf("Nanosecond: (-got +want)\n%s; %s", diff, e)
	}

	f.Precision = time.Microsecond
	s, e = json.Marshal(f)
	if diff := cmp.Diff(string(s), `1640274432937853`); diff != "" || e != nil {
		t.Errorf("Microsecond: (-got +want)\n%s; %s", diff, e)
	}

	f.Precision = time.Millisecond
	s, e = json.Marshal(f)
	if diff := cmp.Diff(string(s), `1640274432937`); diff != "" || e != nil {
		t.Errorf("Millisecond: (-got +want)\n%s; %s", diff, e)
	}
}
func TestNormalFloatTimeUnmarshal(t *testing.T) {
	given := []byte(`1640274432`)
	tm := &FloatTime{}
	tm.UnmarshalJSON(given)
	if diff := cmp.Diff(tm.Format(time.RFC3339), `2021-12-23T09:47:12-06:00`); diff != "" {
		t.Errorf("Nanosecond: (-got +want)\n%s", diff)
	}
}
func TestNormalFloatTimeMarshal(t *testing.T) {
	f := &FloatTime{
		Time: time.Unix(1640274432, 0),
	}

	f.Precision = time.Nanosecond
	s, e := json.Marshal(f)
	if diff := cmp.Diff(string(s), `1640274432000000000`); diff != "" || e != nil {
		t.Errorf("Nanosecond: (-got +want)\n%s; %s", diff, e)
	}

	f.Precision = time.Microsecond
	s, e = json.Marshal(f)
	if diff := cmp.Diff(string(s), `1640274432000000`); diff != "" || e != nil {
		t.Errorf("Microsecond: (-got +want)\n%s; %s", diff, e)
	}

	f.Precision = time.Millisecond
	s, e = json.Marshal(f)
	if diff := cmp.Diff(string(s), `1640274432000`); diff != "" || e != nil {
		t.Errorf("Millisecond: (-got +want)\n%s; %s", diff, e)
	}
}

// func TestFloatTimeMarshalNew(t *testing.T) {
// 	f := NewFloatTime()
// 	s, e := json.Marshal(f)
// 	if diff := cmp.Diff(string(s), `1640709402`); diff != "" || e != nil {
// 		t.Errorf("Millisecond: (-got +want)\n%s; %s", diff, e)
// 	}

// 	if diff := cmp.Diff(f.Format(time.RFC3339), `2021-12-28T16:36:42Z`); diff != "" {
// 		t.Errorf("Nanosecond: (-got +want)\n%s", diff)
// 	}
// }
