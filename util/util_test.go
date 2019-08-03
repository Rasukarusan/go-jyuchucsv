package util

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestGetDate(t *testing.T) {
	time, _ := time.Parse("2006-01-02 15:04", "2014-12-31 18:10")
	got := Date(time)
	want := "2014/12/31 18:10"
	if got != want {
		t.Errorf(result(got, want))
	}
}

func TestGetRandomStr(t *testing.T) {
	time, _ := time.Parse("2006-01-02 15:04", "2014-12-31 18:10")
	got := UnixNanoStr(time)
	want := "1420049400000000000"
	if got != want {
		t.Errorf(result(got, want))
	}
}

func TestReadStdin(t *testing.T) {
	want := "100"
	var stdin bytes.Buffer
	stdin.Write([]byte(want + "\n"))
	got := ReadStdin(&stdin)
	if got != want {
		t.Errorf(result(got, want))
	}
}

func result(got interface{}, want interface{}) string {
	return fmt.Sprintf("\ngot: (%T) %#v\nwant: (%T) %#v", got, got, want, want)
}
