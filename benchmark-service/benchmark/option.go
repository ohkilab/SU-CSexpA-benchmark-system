package benchmark

type option struct {
	threadNum   int
	attmptCount int
}

type optionFunc func(o *option)

func OptThreadNum(num int) optionFunc {
	return func(o *option) {
		o.threadNum = num
	}
}

func OptAttemptCount(count int) optionFunc {
	return func(o *option) {
		o.attmptCount = count
	}
}
