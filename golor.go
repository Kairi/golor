package golor

import (
	"fmt"
)

type ColorCode string
const (
	BLACK = "30"
	RED = "31"
	GREEN = "32"
	YELLOW = "33"
	BLUE = "34"
	MAGENTA = "35"
	SYAN = "36"
	WHITE = "37"
)

func Colorlize(str string , color ColorCode) {
	color_prefix := "\x1b[" + string(color) + "m"
	color_suffix := "\x1b[0m"
	fmt.Println(color_prefix + str + color_suffix)
}

