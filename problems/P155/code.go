package P155

type MinStack struct {
	i   int
	arr []int
}

func Constructor() MinStack {
	return *new(MinStack)
}

func (this *MinStack) Push(x int) {
	if this.i == 0 {
		this.i = 1
		this.arr = []int{x}

		return
	}

	this.i++
	this.arr = append(this.arr, x)
}

func (this *MinStack) Pop() {
	if this.i > 0 {
		this.i--
		this.arr = this.arr[:this.i]
	}
}

func (this *MinStack) Top() int {
	if this.i == 0 {
		return 0
	}

	return this.arr[this.i-1]
}

func (this *MinStack) GetMin() int {
	if this.i == 0 {
		return 0
	}

	min := this.arr[0]
	for i:=1;i<this.i;i++{
		if min > this.arr[i] {
			min = this.arr[i]
		}
	}

	return min
}
