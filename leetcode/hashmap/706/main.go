package main

func main() {

}

type MyHashMap struct {
	data []int
}

const NIL = -1

func Constructor() MyHashMap {
	init := make([]int, 1_000_001)
	for i := range init {
		init[i] = NIL
	}
	return MyHashMap{data: init}
}

func (this *MyHashMap) Put(key int, value int) {
	this.data[key] = value
}

func (this *MyHashMap) Get(key int) int {
	return this.data[key]
}

func (this *MyHashMap) Remove(key int) {
	this.data[key] = NIL
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
