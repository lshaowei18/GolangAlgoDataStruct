package findmedian

type MedianFinder struct {
	elements []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	this.elements = append(this.elements, num)

	i := len(this.elements) - 1 // index of child
	for {
		p := (i - 1) / 2

		//index value more than parent value, break
		if this.elements[i] >= this.elements[p] {
			break
		}

		this.swap(i, p)

		i = p
	}
}

func (this *MedianFinder) Dequeue() int {
	dequeued := this.elements[0]

	this.elements[0] = this.elements[len(this.elements)-1]
	this.elements = this.elements[:len(this.elements)-1]

	i := 0

	for {
		n := len(this.elements)

		left := i*2 + 1

		if left >= n {
			break
		}

		swap := left

		if right := left + 1; right < n && this.less(right, left) {
			swap = right
		}

		if this.less(i, swap) {
			break
		}

		this.swap(i, swap)

		i = swap
	}
	return dequeued
}

func (this *MedianFinder) FindMedian() float64 {
	n := len(this.elements)
	mid := (n - 1) / 2

	median := 0

	for i := 0; i <= mid; i++ {
		median = this.Dequeue()
	}

	if n%2 != 0 {
		return float64(median)
	}

	median2 := this.Dequeue()
	sum := float64(median + median2)

	return sum / 2
}

func (mf *MedianFinder) swap(i, j int) {
	mf.elements[i], mf.elements[j] = mf.elements[j], mf.elements[i]
}

func (mf *MedianFinder) less(i, j int) bool {
	return mf.elements[i] < mf.elements[j]
}
