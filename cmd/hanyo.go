package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Rasukarusan/go-jyuchucsv/util"
	"github.com/spf13/cobra"
)

var hanyoCmd = &cobra.Command{
	Use:  "hanyo",
	Long: `汎用店舗のCSVファイルを作成します`,
	Run: func(cmd *cobra.Command, args []string) {
		main()
	},
}

func init() {
	rootCmd.AddCommand(hanyoCmd)
}

var now = time.Now
var resultCsvPath = util.HomeDir() + "/Desktop/hanyo.csv"
var tc = util.TemplateCsv{}

func main() {
	tc.SetTemplateCsv(util.PathTemplateCsvDir + "hanyo.csv")
	for {
		fmt.Print("何件の受注伝票を作成しますか？(数字だけ入力):")
		text := util.ReadStdin(os.Stdin)
		if text == "" {
			return
		}
		tc.WriteHeader(resultCsvPath)
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("\x1b[31mERROR\x1b[0m：整数を入力してください")
			continue
		}
		writeRecords(num)
		fmt.Printf("%d件のCSVを作成しました。\n", num)
		break
	}
}

func writeRecords(num int) {
	file, err := os.OpenFile(resultCsvPath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	for i := 0; i < num; i++ {
		record := createRecord(tc, i)
		writer.Write(record)
	}
	writer.Flush()
}

func createRecord(tc util.TemplateCsv, count int) []string {
	t := now()
	date := util.Date(t)
	s := util.UnixNanoStr(t)
	// 動的に変更しなければならないものだけ置換する
	record := tc.Record
	record = strings.Replace(record, "TENPO_DENPYO_NO", fmt.Sprintf("HANYO-%s", s), 1)
	record = strings.Replace(record, "JYUCHU_BI", fmt.Sprintf("%s", date), 1)
	record = strings.Replace(record, "JYUCHU_JYUSYO_2", fmt.Sprintf("%s", s), 1)
	record = strings.Replace(record, "HASOU_JYUSYO_2", fmt.Sprintf("%s", s), 1)
	return strings.Split(record, ",")
}
