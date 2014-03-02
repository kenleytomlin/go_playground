package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("s:", s)
	fmt.Println("s len:", len(s))
	// Append returns a new slice
	d := append(s, "d", "e")
	// If you don't use the variable then program will not compile
	fmt.Println("d:", d)
	e := make([]string, len(d))
	// Copy doesn't return a value
	copy(e, d)
	fmt.Println("copied d to e, e is:", e)
	f := d[2:5]
	fmt.Println("took index 2 to index 5 from d", f)
	fmt.Println("take from 0 to 5 of d", d[:5])
	fmt.Println("take from 2 to 3:", d[2:3])
	fmt.Println("take from 3 to the end:", d[3:])

	// Single line initialization of slice
	g := []string{"f", "g", "h", "i", "j"}
	fmt.Println("g:", g)
	// Can also add a cap to inner lengths
	twoD := make([][]int, 3, 3)
	for i := 0; i < 3; i++ {
		// the inner length of slices can be calculated unlike arrays
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
