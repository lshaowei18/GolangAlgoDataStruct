package topkfrequent

type IntFrequency struct {
	Val  int
	Freq int
}

type FreqSlice []IntFrequency

func topKFrequent(nums []int, k int) (ints []int) {
	m := make(map[int]int)

	for _, n := range nums {
		m[n]++
	}

	s := FreqSlice{}

	for key, freq := range m {
		i := IntFrequency{key, freq}
		s.Enqueue(i)
	}

	for i := 0; i < k; i++ {
		n := s.Dequeue()
		ints = append(ints, n)
	}

	return
}

func (s *FreqSlice) Enqueue(intf IntFrequency) {
	(*s) = append((*s), intf)
	i := len((*s)) - 1 // index of child

	for {
		p := (i - 1) / 2 //index of parent

		//Index frequency must be lesser for swapping to stop
		if (*s)[i].Freq <= (*s)[p].Freq {
			break
		}
		//Swap
		s.swap(i, p)
		i = p
	}
}

func (s *FreqSlice) Dequeue() int {
	dequeued := (*s)[0]

	(*s)[0] = (*s)[len((*s))-1] //Replace last node with first node
	(*s) = (*s)[:len((*s))-1]

	i := 0

	for {
		n := len((*s))

		left := i*2 + 1

		//Once left child is out of range, both children will be
		if left >= n {
			break
		}

		swap := left

		// right is not out of range and more than left
		if right := left + 1; right < n && !(*s).less(right, left) {
			swap = right
		}

		//if i is higher than swap child, break
		if !(*s).less(i, swap) {
			break
		}

		(*s).swap(i, swap)

		i = swap
	}
	return dequeued.Val

}

//Higher number has lower priority
func (s *FreqSlice) less(i, j int) bool {
	return (*s)[i].Freq < (*s)[j].Freq
}

func (s *FreqSlice) swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
