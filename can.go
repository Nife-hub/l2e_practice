package main

func CanJump2(nums []uint) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}
	currentIndex := 0          // start at index 0
	lastIndex := len(nums) - 1 // last number of the slice

	for currentIndex < len(nums) {
		if currentIndex == lastIndex {
			return true // when currentIndex gets to lastIndex, it becomes true
		}
		if nums[currentIndex] == 0 {
			return false // if the number in the in the slice 0, return false
		}
		nextIndex := currentIndex + int(nums[currentIndex])

		if nextIndex > lastIndex {
			return false
		}
		currentIndex = nextIndex
	}
	return false
}
