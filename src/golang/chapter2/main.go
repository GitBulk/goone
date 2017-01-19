package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world, I am Sieu Nhan Gao, time:", time.Now())
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println(len("Hello world"))
	fmt.Println("Hello world"[1])
	fmt.Println("Hello " + "world")
	fmt.Println("5 * 2 =", 5*2)
	var message string = "First Gopher"
	fmt.Println(message)
	message = "Second Gopher"
	fmt.Println(message)

	dogName := "Lucky"
	fmt.Println("My dog's name is", dogName)
	// fmt.Println("Enter a number: ")
	// var input float64
	// fmt.Scanf("%f", &input)
	// output := input * 2
	// fmt.Println(output)
	// fmt.Println(`1
	//     2
	//     3
	//     4
	//     5
	//     6`)
	for i := 0; i < 10; i++ {
		// fmt.Println("Index", i)
		if (i % 2) == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}

	// array
	var arrInt [5]int
	arrInt[4] = 8
	fmt.Println(arrInt)

	arrFloat := [5]float64{98, 93, 77, 82, 83}

	var total float64 = 0
	for _, value := range arrFloat {
		total += value
	}
	fmt.Println(total / float64(len(arrFloat)))

	fmt.Println("------ Slice ------")
	slice1 := []int{7, 5, 1}
	slice2 := append(slice1, 10, 8)
	fmt.Println(slice1, slice2)

	slice3 := make([]int, 2)
	copy(slice3, slice1)
	fmt.Println(slice1, slice3)

	fmt.Println("------ Map ------")
	dicUser := make(map[string]int)
	dicUser["toan"] = 30
	dicUser["lisa"] = 28
	fmt.Println(dicUser)
	delete(dicUser, "toan")
	fmt.Println(dicUser)
	fmt.Println(dicUser["lisa"])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	if name, ok := elements["Li"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("oop")
	}

	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5])

	complexElements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := complexElements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	smallest := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	min := smallest[0]
	for index := 1; index < len(smallest); index++ {
		if min > smallest[index] {
			min = smallest[index]
		}
	}
	fmt.Println(min)
}
