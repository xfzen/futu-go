package utils

import (
	"fmt"

	"github.com/hokaccha/go-prettyjson"
)

func PrettyJson(v interface{}) string {
	s, _ := prettyjson.Marshal(v)
	return string(s)
}

func DumpSlice(title string, list []float64) {
	fmt.Printf("%v:", title)

	nlen := len(list)
	for i := (nlen - 10); i < len(list); i++ {
		fmt.Printf("%10.5f", list[i])
	}

	fmt.Printf("\n")
}
