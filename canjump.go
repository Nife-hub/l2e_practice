package main

func CanJump(nums []uint) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}
	current := 0
	lastIndex := len(nums) - 1

	for current < len(nums) {
		if current == lastIndex {
			return true
		} 
		if nums[current] == 0 {
			return false
		}
		next := current + int(nums[current])

		if next > lastIndex {
			return false
		}
		current = next
	}
	return false
}