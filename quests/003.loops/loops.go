package loops

import "strings"

func SumEvenNumbers(n int) int {
	// TODO: implement
	// Read README.md for the instructions

	if n <= 0 {
		return 0  
	}

	sum := 0 

	for i := 1; i <= n; i++ {
		if i % 2 == 0 {
			sum += i
		}
	}
	return sum 
}

func KeepOnlyConsonants(strs []string) []string {
	res := []string{}
	// TODO: implement
	// Read README.md for the instructions

	for _, value := range strs{
		onlyConstants := ""
		onlyConstants = removeVowels(string(value))

		if(onlyConstants != ""){
			res = append(res, onlyConstants)
		}
		
	}

	return res
	
}

func removeVowels(s string) string {
	vowels := "aeiouAEIOU"
	result := ""
	for _, value := range s{
		if !strings.Contains(vowels,string(value)) {
			result += string(value)
		}
	}

	return result
	
}
