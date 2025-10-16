package fizz

import "fmt"

// Fizz gibt die Zahlen von 1 bis 30 aus und ersetzt dabei
// jede durch 3 teilbare Zahl durch fizz und jede durch 5 teilbare
// Zahl durch buzz.
// Bei Zahlen dei sowohl durch 3 als auch durch 5 teilbar sind,
// Wird fizzbuzz ausgegeben
func FizzImproved(x, y, n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)

		if i%y == 0 && i%x == 0 {
			fmt.Println("fizzbuzz")
		} else if i%x == 0 {
			fmt.Println("fizz")

		} else if i%y == 0 {
			fmt.Println("buzz")

		} else {
			fmt.Println(i)
		}

	}
}
