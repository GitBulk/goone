package main

import "fmt"

func main() {
	slice := make([]string, 3)
	slice[0] = "toan"
	slice[1] = "tran"
	slice[2] = "kaka"
	fmt.Println(slice)

	fmt.Println("-- append slice --")
	slice = append(slice, "manchester", "united")
	fmt.Println(slice)

	fmt.Println("-- copy slice --")
	copySlice := make([]string, len(slice))
	copy(copySlice, slice)
	fmt.Println("CopySlice: ", copySlice)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
