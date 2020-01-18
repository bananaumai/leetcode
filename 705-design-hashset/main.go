package main

import "fmt"

type MyHashSet struct {
	arr []bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	arr := make([]bool, 1000001)
	return MyHashSet{arr}
}

func (this *MyHashSet) Add(key int) {
	this.arr[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.arr[key] = false
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.arr[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
func main() {
	obj := Constructor()
	obj.Add(1)
	obj.Add(2)
	fmt.Printf("obj.Contains(1) should be true: %v\n", obj.Contains(1))
	fmt.Printf("obj.Contains(2) should be true: %v\n", obj.Contains(1))
	fmt.Printf("obj.Contains(3) should be false: %v\n", obj.Contains(3))
	obj.Remove(2)
	fmt.Printf("obj.Contains(2) should be false: %v\n", obj.Contains(2))
	obj.Add(1000000)
	fmt.Printf("obj.Contains(1000000) should be true: %v\n", obj.Contains(1000000))
	// obj.Add(1000001)
}
