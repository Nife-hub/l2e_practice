package main

func CanJump(nums []uint) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1{
		return true
	}

	current := 0
	last := len(nums)-1

	for current < len(nums) {
		if current == last {
			return true
		}
		if nums[current] == 0{
			return false
		}
		next := current + int(nums[current])
		if next > last {
			return false
		}
		current = next
	}
	return false
}



func CanJump(arr []uint) bool{
	if len(arr) == 0{
		return false
	}
	if len(arr) == 1{
		return true
	}

	position := 0
	lastIndex := len(arr)-1

	for {
		if arr[position] == 0 {
			return false
		}
		position = position + int(arr[position])

		if position > lastIndex {
			return false
		}
		if position == lastIndex{
			return true
		}
	}
}