package libs

import (
	"fmt"
	"os"
)

func ErrorExit(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
