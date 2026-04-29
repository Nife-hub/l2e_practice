package main

func ConcatSlice(slice1, slice2 []int) []int {
	if len(slice1) == 0 && len(slice2) == 0 {
		return nil
	}
	slice1 = append(slice1, slice2...)
	return slice1
}