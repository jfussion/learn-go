package iteration

func Repeat(s string, count int) (repeated string) {
	for i := 0; i < count; i++ {
		repeated += s
	}

	return
}
