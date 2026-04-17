package main

func ConcatSlice(slice1, slice2 []int) []int {
	slice1 = append(slice1, slice2...)
	return slice1
}