/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/spf13/cobra"
)

// hanyoCmd represents the hanyo command
var hanyoCmd = &cobra.Command{
	Use:  "hanyo",
	Long: `汎用店舗のCSVファイルを作成します`,
	Run: func(cmd *cobra.Command, args []string) {
		main()
	},
}

func init() {
	rootCmd.AddCommand(hanyoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hanyoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hanyoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var header = []string{
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

var value = []string{}

var CSV_PATH = getHomeDir() + "/Desktop/hanyo.csv"

func main() {
	num := question("何件のCSVを作成しますか？(数字だけ入力):")
	fmt.Printf("%v", num)
	createHeader()
	createValue()
}

func getHomeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

func question(q string) string {
	fmt.Print(q)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	return text
}

func createHeader() {
	file, err := os.Create(CSV_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Write(header)
	writer.Flush()
}

func createValue() {
	file, err := os.OpenFile(CSV_PATH, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, "hgoehogheoge")
}
