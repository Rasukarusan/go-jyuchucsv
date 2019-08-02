package util

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestGetDate(t *testing.T) {
	time, _ := time.Parse("2006-01-02 15:04", "2014-12-31 18:10")
	actual := Date(time)
	expected := "2014/12/31 18:10"
	if actual != expected {
		t.Errorf(result(actual, expected))
	}
}

func TestGetRandomStr(t *testing.T) {
	time, _ := time.Parse("2006-01-02 15:04", "2014-12-31 18:10")
	actual := UnixNanoStr(time)
	expected := "1420049400000000000"
	if actual != expected {
		t.Errorf(result(actual, expected))
	}
}

func TestReadStdin(t *testing.T) {
	expected := "100"
	var stdin bytes.Buffer
	stdin.Write([]byte(expected + "\n"))
	actual := ReadStdin(&stdin)
	if actual != expected {
		t.Errorf(result(actual, expected))
	}
}

func result(actual interface{}, expected interface{}) string {
	return fmt.Sprintf("\nactual: (%T) %#v\nexpected: (%T) %#v", actual, actual, expected, expected)
}
