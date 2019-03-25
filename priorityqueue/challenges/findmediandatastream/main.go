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
}

func (this *MedianFinder) FindMedian() float64 {
	n := len(this.elements)

	i := (n - 1) / 2

	if n%2 != 0 {
		return float64(this.elements[i])
	}

	sum := float64(this.elements[i] + this.elements[i+1])

	return sum / 2
}
