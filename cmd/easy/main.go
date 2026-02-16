package main

func main() {
	numbers := [5]int8{1, 2, 3, 4, 5}
	var sum int8
	for _, n := range numbers {
		sum += n
	}
}
