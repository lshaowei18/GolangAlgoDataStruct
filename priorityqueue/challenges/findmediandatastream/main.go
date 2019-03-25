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

func (this *MedianFinder) FindMedian() float64 {
	n := len(this.elements)

	i := (n - 1) / 2

	if n%2 != 0 {
		return float64(this.elements[i])
	}

	sum := float64(this.elements[i] + this.elements[i+1])

	return sum / 2
}

func (mf *MedianFinder) swap(i, j int) {
	mf.elements[i], mf.elements[j] = mf.elements[j], mf.elements[i]
}
