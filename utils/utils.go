package utils

import "log"

func CheckError(err error) {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	if err != nil {
		log.Fatal(err)
	}
}
