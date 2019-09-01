package strconvuint

import "strconv"

const maxUint64 = 1<<64 - 1

func Atouint(s string) (uint, error) {
	ui64, err := strconv.ParseUint(s, 10, 0)
	return uint(ui64), err
}
