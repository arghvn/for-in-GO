package main

// for is an example of methods in GO

import "fmt"

func main() {

	i := 1

	for i <= 3 {

		fmt.Println(i)

		i = i + 1

	}

	for j := 7; j <= 9; j++ {

		fmt.Println(j)

	}

	for {

		fmt.Println("loop")

		break

	}

	for n := 0; n <= 5; n++ {

		if n%2 == 0 {

			continue

		}

		fmt.Println(n)

	}

}

// output :
// 1
// 2
// 3
// 7
// 8
// 9
// loop
// 1
// 3
// 5

// more details :
//The third example is an infinite loop that can be stopped with the break keyword
