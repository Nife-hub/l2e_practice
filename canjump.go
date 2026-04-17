package main

func CanJump(nums []uint) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}
	currentIndex := 0
	lastIndex := len(nums) - 1

	for currentIndex < len(nums) {
		if currentIndex == lastIndex {
			return true
		} 
		if nums[currentIndex] == 0 {
			return false
		}
		nextIndex := currentIndex + int(nums[currentIndex])

		if nextIndex > lastIndex {
			return false
		}
		currentIndex = nextIndex
	}
	return false
}