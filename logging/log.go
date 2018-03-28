package logging

import (
	"fmt"
)

func Debugf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf("[DEBUG]"+format, a)
}
