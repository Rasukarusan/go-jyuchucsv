package util

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strings"

	_ "github.com/Rasukarusan/go-jyuchucsv/statik"

	"github.com/rakyll/statik/fs"
)

// Recordは動的に置換したいため配列ではなく文字列にする
type TemplateCsv struct {
	Header []string
	Record string
}

// テンプレートCSVを構造体にセット
func (tc *TemplateCsv) SetTemplateCsv(templateCsvPath string) {
	FS, _ := fs.New()
	fp, err := FS.Open(templateCsvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	scanner.Scan()
	tc.Header = strings.Split(scanner.Text(), ",")
	scanner.Scan()
	tc.Record = scanner.Text()
}

// 出力用CSVにヘッダーを書き込む
func (tc *TemplateCsv) WriteHeader(resultCsvPath string) {
	file, err := os.Create(resultCsvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Write(tc.Header)
	writer.Flush()
}
