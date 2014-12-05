package golar

import (
	"fmt"
)

var (
	black_code = "30"
	red_code = "31"
	green_code = "32"
	yellow_code = "33"
	blue_code = "34"
	magenta_code = "35"
	syan_code = "36"
	white_code = "37"
)

func colorlize(str, color string) {
	color_prefix := "\x1b[" + color + "m"
	color_suffix := "\x1b[0m"
	fmt.Println(color_prefix + str + color_suffix)
}

