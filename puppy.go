package puppy

func AddNumbers(num int) int {
	var sum = 0
	for num > 0 {
		var temp = num % 10
		sum += temp
		num = num / 10
	}
	return sum
}

func Substract(x int, y int) int {
	return x - y
}
