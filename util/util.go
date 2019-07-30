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

func GetDate() string {
	const format = "2006/01/02 15:04"
	t := time.Now()
	return t.Format(format)
}

func GetRandomStr() string {
	t := time.Now()
	return strconv.FormatInt(t.UnixNano(), 10)
}
