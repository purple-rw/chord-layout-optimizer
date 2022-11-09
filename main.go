package main

import "fmt"

func main() {
	fmt.Println("Test")
}

func optimize() {

	// 1. Create a pool containing `p` randomly-generated keyboard layouts.
	// 2. Score each keyboard according to a fitness function and sort the keyboards by score.
	// 3. Randomly delete half of the pool (giving preference to keyboards with lower fitness) and create a copy of each remaining keyboard.
	// 4. Mutate the copies by randomly swapping the positions of two keys `m` times.
	// 5. Repeat steps 2-4 until the best keyboard in the pool has not changed for `b` rounds.
	// 6. Place this best keyboard in pool O and sort each keyboard in O by score.
	// 7. Repeat steps 2-6 until the best keyboard in pool O has not changed for `o` rounds.
	// 8. Repeat steps 2-4 using pool O until the best keyboard in the pool has not changed for `q` rounds.
}