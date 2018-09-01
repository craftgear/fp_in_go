package fibonacci

type memoized func(int) int

func MyFibMemoized(n int) int {
	return getMemoizedFib()(n)
}

func memoize(mf memoized) memoized {
	cache := make(map[int]int)
	return func(key int) int {
		if val, found := cache[key]; found {
			return val
		}
		temp := mf(key)
		cache[key] = temp
		return temp
	}
}

func getMemoizedFib() func(int) int {
	var innerFib func(x int) int

	innerFib = memoize(func(x int) int {
		if x == 0 {
			return 0
		} else if x <= 2 {
			return 1
		} else {
			return innerFib(x-2) + innerFib(x-1)
		}
	})
	return innerFib
}
