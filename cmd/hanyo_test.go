package cmd

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/Rasukarusan/go-jyuchucsv/util"
)

func TestCreateRecord(t *testing.T) {
	dir, _ := os.Getwd()
	var tc = util.TemplateCsv{}
	tc.SetTemplateCsv(dir + "/../" + util.PathTemplateCsvDir + "hanyo.csv")
	actual := createRecord(tc, 1)
	expected := []string{""}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(result(actual, expected))
	}
}

func result(actual interface{}, expected interface{}) string {
	return fmt.Sprintf("\nactual: (%T) %#v\nexpected: (%T) %#v", actual, actual, expected, expected)
}
