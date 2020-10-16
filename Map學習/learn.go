package main

import "fmt"

func main() {
	// 怎麼建立一個簡單的 Map 呢？
	// var x map[string]int
	x := make(map[int]int)
	x[1] = 10
	fmt.Println(x[0])

	fakeData := map[string]string{
		"one":   "one",
		"two":   "two",
		"three": "three",
		"four":  "four",
	}
	fmt.Println(fakeData)

	elements := map[string]map[string]string{
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
	fmt.Println(elements)
}
