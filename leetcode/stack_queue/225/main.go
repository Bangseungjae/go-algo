package main

func main() {

}

type MyStack struct {
	data []int
}

func Constructor() MyStack {
	return MyStack{data: make([]int, 0, 100)}
}

func (this *MyStack) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyStack) Pop() int {
	if len(this.data) < 1 {
		return -1
	}
	data := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	return data
}

func (this *MyStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.data) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
