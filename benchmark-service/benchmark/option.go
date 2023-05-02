package benchmark

import "time"

type option struct {
	threadNum int
	timeout   time.Duration
}

type optionFunc func(o *option)

func OptThreadNum(num int) optionFunc {
	return func(o *option) {
		o.threadNum = num
	}
}

func OptTimeout(timeout time.Duration) optionFunc {
	return func(o *option) {
		o.timeout = timeout
	}
}
