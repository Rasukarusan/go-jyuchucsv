package util

import (
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
