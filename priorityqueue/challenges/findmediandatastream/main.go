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
