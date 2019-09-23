package utils

import (
	"log"
	"math"
	"time"
)

func PanicIfErr(err error, messages ...string) {
	if err != nil {
		if len(messages) > 0 {
			log.Printf("%q", messages)
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
