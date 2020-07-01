package main

import "fmt"

func main() {
	// Using the code from the previous example, add a record to your map.
	// Now print the map out using the “range” loop

	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	m["job_odd"] = []string{"Hats", "Golf", "Buttling"}

	for lastFirst, slice := range m {
		fmt.Println(lastFirst)
		for idx, value := range slice {
			fmt.Printf("\t Index %v \t %v\n", idx, value)
		}
	}

	// output:
	// job_odd
	// 		Index 0         Hats
	// 		Index 1         Golf
	// 		Index 2         Buttling
	// bond_james
	// 		Index 0         Shaken, not stirred
	// 		Index 1         Martinis
	// 		Index 2         Women
	// moneypenny_miss
	// 		Index 0         James Bond
	// 		Index 1         Literature
	// 		Index 2         Computer Science
	// no_dr
	// 		Index 0         Being evil
	// 		Index 1         Ice cream
	// 		Index 2         Sunsets
}
