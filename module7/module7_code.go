package module7

import (
	"fmt"
	"runtime"
)

//go:generate goimports -w module7_code.go

func content() {
	fmt.Println(runtime.GOOS)
}
