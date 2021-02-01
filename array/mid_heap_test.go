package array

import (
	"fmt"
	"testing"
)

func TestConstructorMedianFinder(t *testing.T) {
	mid := ConstructorMedianFinder()
	mid.AddNum(-1)
	fmt.Println(mid.FindMedian())
	mid.AddNum(-2)
	fmt.Println(mid.FindMedian())
	mid.AddNum(-3)
	fmt.Println(mid.FindMedian())
	mid.AddNum(-4)
	fmt.Println(mid.FindMedian())
	mid.AddNum(-5)
	fmt.Println(mid.FindMedian())
}
