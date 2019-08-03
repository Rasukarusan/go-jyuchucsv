package cmd

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/Rasukarusan/go-jyuchucsv/util"
)

func TestCreateRecord(t *testing.T) {
	tc := util.TemplateCsv{}
	tc.SetTemplateCsv("test-csv-template/hanyo.csv")
	now = func() time.Time { return time.Date(2019, 12, 31, 18, 10, 0, 0, time.Local) }
	got := createRecord(tc)
	want := []string{"HANYO-1577783400000000000", "2019/12/31 18:10", "1577783400000000000", "1577783400000000000"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(result(got, want))
	}
	now = time.Now
}

func result(got interface{}, want interface{}) string {
	return fmt.Sprintf("\ngot: (%T) %#v\nwant: (%T) %#v", got, got, want, want)
}
