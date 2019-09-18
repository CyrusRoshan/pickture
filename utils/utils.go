package utils

import "fmt"

func PanicIfErr(err error, messages ...string) {
	if err != nil {
		if len(messages) > 0 {
			fmt.Println(messages)
		}
		panic(err)
	}
}
