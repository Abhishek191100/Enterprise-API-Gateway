package utils

import (
	"fmt"
)

func CheckError(e error) {
	if e != nil {
		fmt.Printf("Unable to proceed because of following error: %v",e)
		panic(" exiting now")
	}
}