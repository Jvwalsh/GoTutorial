package main

import "fmt"

// package main

// import "fmt"

// func main() {

// 	//MAPS!
// 	//Makes can have string or int keytypes
// 	//must be declared with make

// 	x := make(map[string]int)
// 	x["key"] = 10
// 	fmt.Println(x["key"])

// 	y := make(map[int]int)
// 	y[12323] = 10
// 	fmt.Println(y[12323])

// }

// package main

// import "fmt"

// func main() {
// 	elements := make(map[string]string)
// 	elements["H"] = "Hydrogen"
// 	elements["He"] = "Helium"
// 	elements["Li"] = "Lithium"
// 	elements["Be"] = "Beryllium"
// 	elements["B"] = "Boron"
// 	elements["C"] = "Carbon"
// 	elements["N"] = "Nitrogen"
// 	elements["O"] = "Oxygen"
// 	elements["F"] = "Fluorine"
// 	elements["Ne"] = "Neon"

// 	fmt.Println(elements["Li"])
// 	//map will return 0 and it will display as blank space
// 	fmt.Println(elements["Un"])
// 	if name, ok := elements["Un"]; ok {
// 		fmt.Println(name, ok)
// 	}
// 	if name, ok := elements["Ne"]; ok {
// 		fmt.Println(name, ok)
// 	}
// }

func main() {
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

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
