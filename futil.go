package futil

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"
)

func InSQL[T int | int64 | int32 | int16 | int8 | string](IDs []T) string {
	if reflect.TypeOf(IDs[0]) == reflect.TypeOf("") {
		var strs []string
		for i := 0; i < len(IDs); i++ {
			str := fmt.Sprintf(`"%v"`, IDs[i])
			strs = append(strs, str)
		}
		return fmt.Sprintf("(%s)", strings.Join(strs, ","))
	}
	jIDS := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(IDs)), ","), "[]")
	return fmt.Sprintf("(%s)", jIDS)
}

func Benchmark() func() int64 {
	start := time.Now()
	return func() int64 {
		elapsed := time.Since(start)
		log.Printf("\tBenchmark: %s\n", elapsed)
		return int64(elapsed / time.Millisecond)
	}
}
