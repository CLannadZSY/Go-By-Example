package main

import "fmt"

func main() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n < 4; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	for i, j := 0, 1; i < 10 && j < 10; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
}
