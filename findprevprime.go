package main

func FindPrevPrime(n int) int{
	if n < 2 {
		return 0
	}
	for x := n; x >= 2; x--{
		isPrime := true

		for i := 2; i*i <= x; i++{
			if x%i == 0{
				isPrime = false
				break
			}
		}
		if isPrime{
			return x
		}
	}
	return 0
}


























func FindPrevPrime2(nb int) int {
	if nb < 2 {
		return 0
	}

	for i := nb; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}
	return 0
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}



