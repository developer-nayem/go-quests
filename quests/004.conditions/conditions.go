package conditions

func ClassifyRequest(age int, hasID bool, balance float64, isVIP bool) string {
	//TODO: implement
	// Read README.md for the instructions
	verdict := ""

	if age <= 0 || balance < 0 {
		verdict = "INVALID"
	} else if age < 18 || hasID == false {
		verdict = "REJECTED"
	} else if isVIP ==  true && balance >= 10000{
		verdict = "VIP_ACCESS"
	} else if balance >= 1000 {
		verdict = "STANDARD_ACCESS"
	} else {
		verdict = "LIMITED_ACCESS"
	}
	return verdict
}

func EvaluateGrade(score int) string {
	//TODO: implement
	// Read README.md for the instructions
	switch  {
	case score < 0 || score > 100:
		return "INVALID"
	case score >= 90 && score <= 100: 
		return "A"
	case score >= 80 && score <= 89: 
		return "B"
	case score >= 70 && score <= 79: 
		return "C"
	case score >= 60 && score <= 69: 
		return "D"
	default:
		/* code */
		return "F"
	}
}
