package exercise

import "container/list"

const base = 769

type Node struct {
	key, value int
}

type MyHashMap struct {
	Node []list.List
}

func Constructor() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}

func (this *MyHashMap) hash(key int) int {
	return key % base
}

func (this *MyHashMap) Put(key int, value int) {
	h := this.hash(key)
	for e := this.Node[h].Front(); e != nil; e = e.Next() {
		if el := e.Value.(Node); el.key == key {
			e.Value = Node{key, value}
			return
		}
	}
	this.Node[h].PushBack(Node{key, value})
}

func (this *MyHashMap) Get(key int) int {
	h := this.hash(key)
	for e := this.Node[h].Front(); e != nil; e = e.Next() {
		if el := e.Value.(Node); el.key == key {
			return el.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	h := this.hash(key)
	for e := this.Node[h].Front(); e != nil; e = e.Next() {
		if e.Value.(Node).key == key {
			this.Node[h].Remove(e)
		}
	}
}
