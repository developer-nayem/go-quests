package functions
import (
	"errors"
)

// Divide returns a / b or an error if b == 0
func Divide(a, b int) (int, error) {
	// TODO: implement
	// Read README.md for instructions on returning proper errors
	if b == 0{
		return 0, errors.New("b cannot be zero")
	} else {
		return a/b, nil 
	}
	
}

// SumAll returns the sum of all provided integers
func SumAll(nums ...int) int {
	// TODO: implement
	// Read README.md for instructions on handling variadic arguments
	sum := 0
	if len(nums) == 0{
		return sum
	} else {
		for _,value := range nums{
			sum += value
		}
	}
	return sum 
	
}

// MaxMin returns the max and min of all provided integers
// Returns an error if no numbers are provided
func MaxMin(nums ...int) (int, int, error) {
	// TODO: implement
	// Read README.md for instructions on proper error handling
	if len(nums) == 0{
		return 0,0,errors.New("no parameters passed")
	}

	min := nums[0]
	max := nums[0]

	for _, value := range nums {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	return max,min,nil
}

// ConcatAll joins all strings using the provided separator
func ConcatAll(sep string, strs ...string) string {
	// TODO: implement
	// Read README.md for instructions on variadic strings

	result := ""
	lastIndex := len(strs) - 1 

	for i := 0; i < len(strs); i++ {
		if i == lastIndex {
			result += strs[i]
		} else {
			result +=  strs[i] + sep 
		}
	}

	return result
}
