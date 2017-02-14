package algorithms

//var toBe [10]int = [10]int{1,3,4,7,9,10,11,16,20,40}

func benchmarkBSortModified(input []int) {
  size := len(input)
  swapped := true
  for swapped {
    swapped = false
    for index := 1; index < size-1; index++ {
      if input[index-1] > input[index] {
        input[index], input[index-1] = input[index-1], input[index]
        swapped = true
      }
    }
  }
}

func QSort(values []int) {
	qsort(values, 0, len(values)-1)
}

func qsort(values []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	qsort(values, l, i-2)
	qsort(values, i, r)
}
