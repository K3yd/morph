package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n >= 12307 {
		fmt.Printf(" %d very big \n", n)
		return
	}

	for n < 12307 {
		if n < 0 {
			n *= -1
		} else if n%7 == 0 {
			n *= 39
		} else if n%9 == 0 {
			n = n*13 + 1
			continue	
		} else {
			n = (n + 2) * 3
		}

		
		if n%13 == 0 && n%9 == 0 {
			fmt.Println("service error")
			return
		} else {
			n++
		}
	}

	fmt.Printf("final result: %d\n", n)
}
