package main


// IntIntMapKV returns slices over the keys and values, separately, of a given map[int]int.
func IntIntMapKV(d map[int] int) ([]int, []int) {
	keys := make([]int, len(d))
	vals := make([]int, len(d))

	i := 0
	for k, v := range d {
		keys[i] = k
		vals[i] = v
		i++
	}

	return keys, vals
}