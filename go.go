package main

import "fmt"

func main() {
	limit := 1000000
	for counter := 1; counter <= limit; counter++ {
		// if counter%100000 == 0 {
		fmt.Println(counter)
		// }

		// if counter == limit {
		// 	fmt.Println("finished", counter)
		// }
	}
}
