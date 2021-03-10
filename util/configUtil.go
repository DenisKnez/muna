package util

import (
	"fmt"

	ini "gopkg.in/ini.v1"
)

//GetConfig returns a pointer to the configuraiton file
func GetConfig() *ini.File {
	config, err := ini.Load("config.ini")

	if err != nil {
		fmt.Println(err)
	}

	return config
}
