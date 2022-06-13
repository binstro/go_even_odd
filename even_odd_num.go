package goevenodd

func CheckEvenOdd(num int) string {
	if num%2 == 0 {
		return "genap"
	} else {
		return "ganjil"
	}
}
