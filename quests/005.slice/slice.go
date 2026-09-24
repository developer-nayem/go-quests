package slice

// naive approach
// func ProcessScores(scores []int) []int {
// 	//TODO: implement
// 	// Read README.md for the instructions

// 	filterdResult := []int{}

// 	for _, value := range scores {
//       if value >= 0 && value  <= 100 {
//           filterdResult = append(filterdResult,value)
//       }
//   	}

//   	normalizedResult := []int{}

//   	for _, value := range filterdResult{
//   		if value < 40{
//   			value = 40
//   			normalizedResult = append(normalizedResult, value)
//   		} else {
//   			normalizedResult = append(normalizedResult, value)
//   		}
//   	}

//   	bonusResult := []int{}

//   	if(len(normalizedResult) > 5){
//   		for _, value := range normalizedResult{
//   			if(value + 5 <= 100){
//   				bonusResult = append(bonusResult, value+5)
//   			}else {
//   				diff := 100 - value
//   				value += diff
//   				bonusResult = append(bonusResult, value)
//   			}
//   		}
//   	}

//   	if(len(bonusResult) == 0){
//   		return normalizedResult
//   	}else {
//   		return bonusResult
//   	}
// }

// optimized way 

func ProcessScores(scores []int) []int {

	result := make([]int, 0, len(scores))

	for _, value := range scores{
		if( value >= 0 && value <= 100){
			if(value < 40){
				value = 40
			}
			result = append(result, value)

		}
	}

	if(len(result) > 5){
		for i := range result {
			result[i] += 5
			if(result[i] > 100){
				result[i] = 100
			}

		}
	}
	return result 
}
