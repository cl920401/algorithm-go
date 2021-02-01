package math

func countDigitOne(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	bit := -1
	for i := 1; n/i > 0; i *= 10 {
		bit++
	}
	// 10 ^ (bit-1) - 1 => sub
	sub := getBitCount(bit)
	head := n / tenPow(bit)
	if head == 1 {
		return sub + countDigitOne(n%tenPow(bit)) + n%tenPow(bit) + 1
	} else {
		return sub*head + countDigitOne(n%tenPow(bit)) + tenPow(bit)
	}
}

// 10的n次方-1 共有几个1
func getBitCount(n int) int {
	if n <= 0 {
		return 0
	}
	return tenPow(n-1) + getBitCount(n-1)*10
}

func tenPow(x int) int {
	ret := 1
	for i := 0; i < x; i++ {
		ret *= 10
	}
	return ret
}

func countDigitOne1(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		for i > 0 {
			if i%10 == 1 {
				count++
			}
			i = i / 10
		}
	}
	return count
}
