package main

func FindPrevPrime(n int) int{
	if n < 2 {
		return 0
	}
	for x := n; x >= 2; x-- {
		isPrime := true

		for i := 2; i*i <= x; i++{
			if x%i == 0 {
				isPrime = false
				break
			}
		} 
		if isPrime {
			return x
		}
	}
	return 0
	
}