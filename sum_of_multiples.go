package summultiples

import (
	"fmt"
	"strconv"
)

// SumMultiples given limit, find the sum of all the unique multiples of particular numbers up to but not including that number
func SumMultiples(limit int, divisors ...int) (sum int) {

	// limit of 20 that are multiples of 3 or 5
	// we get 3, 5, 6, 9, 10, 12, 15, and 18.
	// The sum of these is 78.

	fmt.Println("limit", limit)
	fmt.Println("divisors", divisors)

	sum = 0
	mults := ""

	for i := 0; i < limit; i++ {
		for _, num := range divisors {
			if num == 0 {
				// avoid a divide by zero
				continue
			}
			if i%num == 0 {
				mults += strconv.Itoa(i) + " "
				sum += i
				break
			}
		}
	}

	fmt.Println("mults", mults)
	fmt.Println("sum", sum)
	return sum
}

func main() {

	twenty := SumMultiples(20, 3, 5)
	if twenty != 78 {
		panic("20 -> 78")
	}
}
