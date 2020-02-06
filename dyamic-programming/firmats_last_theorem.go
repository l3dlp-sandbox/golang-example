package dyamic_programming

func FermatsLastTheorem() {
	//var a,b,c int
	//for{
	//
	//}
}
func powerCalculation(val, pow int) int {
	var result int
	if pow == 0 {
		return 1
	}
	result = val
	for i := 1; i < pow; i++ {
		result = result * val
	}
	return result
}
