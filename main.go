package main

import "fmt"

type Node struct {
	key, value string
	prev, next *Node
}

type lruCache struct {
	capacity int
	data     map[string]*Node
	head     *Node
	tail     *Node
}

type LRUCache interface {
	Add(key, value string) bool

	Get(key string) (value string, ok bool)

	Remove(key string) (ok bool)
}

func NewLRUCache(n int) *lruCache {

	return &lruCache{
		capacity: n,
		data:     make(map[string]*Node),
	}
}

func (l *lruCache) Add(key, value string) bool {
	if _, ok := l.data[key]; ok {
		return false
	}

	node := &Node{key: key, value: value}

	if len(l.data) >= l.capacity {
		tail := l.tail
		l.RemoveNode(l.tail)
		fmt.Println("tail==", tail.key)
		delete(l.data, tail.key)
	}

	l.data[key] = node
	l.AddNode(node)
	fmt.Println("head", l.head.key, "tail", l.tail.key)
	return true
}

func (l *lruCache) Get(key string) (string, bool) {

	val, ok := l.data[key]
	if !ok {
		return "", false
	}
	l.moveToFront(val)

	fmt.Println("head", l.head.key, "tail", l.tail.key)
	return val.value, true
}

func (l *lruCache) moveToFront(node *Node) {
	l.RemoveNode(node)
	l.AddNode(node)
}

func (l *lruCache) Remove(key string) bool {
	val, ok := l.data[key]
	if !ok {
		fmt.Println("head", l.head.key, "tail", l.tail.key)
		return false
	}
	fmt.Println("head", l.head.key, "tail", l.tail.key)

	delete(l.data, key)
	l.RemoveNode(val)

	return true
}

func (l *lruCache) RemoveNode(node *Node) {

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.tail = node.prev
	}
}

func (l *lruCache) AddNode(node *Node) {

	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l *lruCache) PrintListNode() {
	current := l.head
	for current != nil {
		fmt.Print(current.key, " <-> ")
		current = current.next
	}
	fmt.Println("nil")
	fmt.Printf("len(l.data) %d , l.capacity %d \n", len(l.data), l.capacity)
	fmt.Println(l.data)
	fmt.Println("===============")
}

func main() {

	lru := NewLRUCache(4)

	lru.Add("A", "A_VAL")
	// lru.PrintListNode()

	lru.Add("B", "B_VAL")
	// lru.PrintListNode()

	lru.Add("C", "C_VAL")
	// lru.PrintListNode()

	lru.Add("D", "D_VAL")
	lru.PrintListNode()

	lru.Add("G", "G_VAL")
	lru.PrintListNode()

	lru.Add("H", "H_VAL")
	lru.PrintListNode()

	fmt.Println("========GET NODE============")

	val, ok := lru.Get("C")
	fmt.Println(val, ok)
	lru.PrintListNode()

	val2, ok2 := lru.Get("D")
	fmt.Println(val2, ok2)
	lru.PrintListNode()

	lru.Add("M", "M_VAL")
	lru.PrintListNode()

	val3, ok3 := lru.Get("C")
	fmt.Println(val3, ok3)
	lru.PrintListNode()

	fmt.Println("========REMOVE NODE============")
	fmt.Println(lru.Remove("D"))
	lru.PrintListNode()
	fmt.Println(lru.Remove("C"))
	lru.PrintListNode()

}

func PrintListNode(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.key, " <-> ")
		current = current.next
	}
	fmt.Println("nil")
}
