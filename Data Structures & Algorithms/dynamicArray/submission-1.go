type DynamicArray struct {
	capacity int
	length   int
	arr      []int
}

func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		capacity: capacity,
		length:   0,
		arr:      make([]int, capacity),
	}

}

func (da *DynamicArray) Get(i int) int {
	if i > da.length {
		return -1
	}

	return da.arr[i]
}

func (da *DynamicArray) Set(i int, n int) {
	if i > da.length {
		return
	}

	da.arr[i] = n
}

func (da *DynamicArray) Pushback(n int) {
	if da.length >= da.capacity {
		da.resize()
	}

	da.arr[da.length] = n
	da.length++
}

func (da *DynamicArray) Popback() int {
	if da.length == 0 {
		return -1
	}

	da.length--
	return da.Get(da.length)
}

func (da *DynamicArray) resize() {
	da.capacity = da.capacity * 2
	newArr := make([]int, da.capacity)

	copy(newArr, da.arr)
	da.arr = newArr
}

func (da *DynamicArray) GetSize() int {
	return da.length
}

func (da *DynamicArray) GetCapacity() int {
	return da.capacity
}
