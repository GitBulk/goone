package main

import "fmt"

func main() {
	fmt.Println("----- Function ------")

	fmt.Println(add(10, 5))
	fmt.Println(add2(10, 7))
	region, continent := location2("LA")
	fmt.Printf("I live in %s, %s", region, continent)

}

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func location(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "LA", "Los Angeles":
		region, continent = "California", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

func location2(city string) (region, continent string) {
	switch city {
	case "LA", "Los Angeles":
		region, continent = "California", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return
}
