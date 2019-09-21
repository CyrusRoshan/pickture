package utils

import (
	"fmt"
	"log"
	"math"
	"time"
)

func PanicIfErr(err error, messages ...string) {
	if err != nil {
		if len(messages) > 0 {
			fmt.Println(messages)
		}
		panic(err)
	}
}

func LogTimeSpent(fn func(), messages ...string) {
	start := time.Now()
	fn()
	elapsed := time.Now().Sub(start)
	log.Printf("Time elapsed %d ms %q", elapsed.Nanoseconds()/int64(math.Pow10(6)), messages)
}
