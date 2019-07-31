package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
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

var csvPath = util.GetHomeDir() + "/Desktop/hanyo.csv"

var HEADER = []string{
	"店舗伝票番号",
	"受注日",
	"受注郵便番号",
	"受注住所１",
	"受注住所２",
	"受注名",
	"受注名カナ",
	"受注電話番号",
	"受注メールアドレス",
	"発送郵便番号",
	"発送先住所１",
	"発送先住所２",
	"発送先名",
	"発送先カナ",
	"発送電話番号",
	"支払方法",
	"発送方法",
	"商品計",
	"税金",
	"発送料",
	"手数料",
	"ポイント",
	"その他費用",
	"合計金額",
	"ギフトフラグ",
	"時間帯指定",
	"日付指定",
	"作業者欄",
	"備考",
	"商品名",
	"商品コード",
	"商品価格",
	"受注数量",
	"商品オプション",
	"出荷済フラグ",
	"顧客区分",
	"顧客コード",
	"ピッキング指示内容",
	"出荷予定日",
}

var recordValue = []string{
	"",
	"",
	"2500011",
	"神奈川県小田原市栄町",
	"",
	"山田太郎",
	"ヤマダタロウ",
	"1234567",
	"test.tarou@test",
	"2500011",
	"神奈川県小田原市栄町",
	"",
	"山田太郎",
	"ヤマダタロウ",
	"1234567",
	"代金引換",
	"佐川急便",
	"2400",
	"120",
	"350",
	"200",
	"0",
	"0",
	"3070",
	"0",
	"",
	"",
	"",
	"",
	"テスト項目選択肢",
	"item-select-l-blue",
	"2400",
	"2",
	"",
	"0",
	"9",
	"",
	"",
	"",
}

func main() {
	for {
		fmt.Print("何件の受注伝票を作成しますか？(数字だけ入力):")
		text := util.ReadStdin(os.Stdin)
		if text == "" {
			return
		}
		writeHeader()
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

func writeHeader() {
	file, err := os.Create(csvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Write(HEADER)
	writer.Flush()
}

func writeRecords(num int) {
	file, err := os.OpenFile(csvPath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	for i := 0; i < num; i++ {
		record := createRecord(i)
		writer.Write(record)
	}
	writer.Flush()
}

func createRecord(count int) []string {
	t := now()
	date := util.GetDate(t)
	s := util.GetRandomStr(t)
	// 動的に変更しなければならないものだけ改めて定義する
	idxDenpyoNo := 0
	idxJyuchuBi := 1
	idxJyuchuJyusyo2 := 4
	idxHasouJyusyo2 := 11
	recordValue[idxDenpyoNo] = fmt.Sprintf("HANYO-%s", s)
	recordValue[idxJyuchuBi] = fmt.Sprintf("%s", date)
	recordValue[idxJyuchuJyusyo2] = fmt.Sprintf("%s", s)
	recordValue[idxHasouJyusyo2] = fmt.Sprintf("%s", s)
	return recordValue
}
