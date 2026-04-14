package main

import "fmt"

func PrintRevComb() {
	for i := 9; i >= 0; i-- {
		for j := 9; j >= 0; j-- {
			for k := 9; k >= 0; k-- {
				if i != k && j != k && i > j && j > k {
					fmt.Print(i)
					fmt.Print(j)
					fmt.Print(k)
					if i != 2 && j != 1 && k != 0 {
						fmt.Print(", ")
					}
				}
			}
		} 
	}
}