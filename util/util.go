package util

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"time"
)

func GetHomeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

func GetDate(t time.Time) string {
	const format = "2006/01/02 15:04"
	return t.Format(format)
}

func GetRandomStr(t time.Time) string {
	return strconv.FormatInt(t.UnixNano(), 10)
}

func ReadStdin(s string) string {
	fmt.Print(s)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	return text
}
