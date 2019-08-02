package util

import (
	"bufio"
	"io"
	"os/user"
	"strconv"
	"time"
)

func HomeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

func Date(t time.Time) string {
	const format = "2006/01/02 15:04"
	return t.Format(format)
}

func UnixNanoStr(t time.Time) string {
	return strconv.FormatInt(t.UnixNano(), 10)
}

func ReadStdin(stdin io.Reader) string {
	scanner := bufio.NewScanner(stdin)
	scanner.Scan()
	text := scanner.Text()
	return text
}
