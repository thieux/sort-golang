package sort

import "testing"
import "fmt"

func sort(a []int) []int {
	sorted := a

	for i := 1; len(sorted) > i; i++ {
		for j := i; j > 0 && sorted[j-1] > sorted[j]; j-- {
			sorted[j-1], sorted[j] = sorted[j], sorted[j-1]
		}
	}

	return sorted
}

func TestEmpty(t *testing.T) {
	a := slice()

	eq(slice(), sort(a), t)
}

func TestSingletonNull(t *testing.T) {
	orig := slice(0)
	exp := slice(0)

	eq(exp, sort(orig), t)
}

func Test2Reversed(t *testing.T) {
	orig := slice(2, 1)
	exp := slice(1, 2)

	eq(exp, sort(orig), t)
}

func Test2Ordered(t *testing.T) {
	orig := slice(1, 2)
	exp := slice(1, 2)

	eq(exp, sort(orig), t)
}

func Test3With2FirstsReversed(t *testing.T) {
	orig := slice(2, 1, 3)
	exp := slice(1, 2, 3)

	eq(exp, sort(orig), t)
}

func Test3With2LastsReversed(t *testing.T) {
	orig := slice(1, 3, 2)
	exp := slice(1, 2, 3)

	eq(exp, sort(orig), t)
}

func Test4With2MiddlesReversed(t *testing.T) {
	orig := slice(1, 3, 2, 4)
	exp := slice(1, 2, 3, 4)

	eq(exp, sort(orig), t)
}

func Test3With1DistantPermutation(t *testing.T) {
	orig := slice(3, 2, 1)
	exp := slice(1, 2, 3)

	eq(exp, sort(orig), t)
}

func Test5With2DistantPermutation(t *testing.T) {
	orig := slice(5, 4, 3, 2, 1)
	exp := slice(1, 2, 3, 4, 5)

	eq(exp, sort(orig), t)
}

func eq(expected []int, actual []int, t *testing.T) {
	if len(expected) != len(actual) {
		logDiff(expected, actual, t)
		t.Fail()
	} else {
		for i := 0; i < len(actual); i++ {
			if expected[i] != actual[i] {
				logDiff(expected, actual, t)
				t.FailNow()
			}
		}
	}
}

func logDiff(expected []int, actual []int, t *testing.T) {
	t.Log(fmt.Sprintf("expected: %v", expected))
	t.Log(fmt.Sprintf("actual: %v", actual))
}

func slice(numbers ...int) []int {
	return numbers
}
