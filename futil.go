package futil

import (
	"fmt"
	"strings"
)

func InSqlInt(length int) string {
	var IDsStr []string
	for i := 0; i < length; i++ {
		IDsStr = append(IDsStr, "?")
	}
	return fmt.Sprintf("(%s)", strings.Join(IDsStr, ","))
}
