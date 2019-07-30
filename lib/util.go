package util

import "os/user"

func GetHomeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
}
