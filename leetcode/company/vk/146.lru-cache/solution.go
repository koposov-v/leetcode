package main

type Node struct {
	Key   int
	Value int
	Next  *Node
	Prev  *Node
}

type LRUCache struct {
	cap        int
	values     map[int]*Node
	Head, Tail *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}

	return LRUCache{
		cap:    capacity,
		values: make(map[int]*Node, capacity),
		Head:   head,
		Tail:   tail,
	}
}

func (l *LRUCache) newNode(key int, val int) *Node {
	return &Node{
		Key:   key,
		Value: val,
	}
}

func (l *LRUCache) Get(key int) int {
	v, ok := l.values[key]
	if !ok {
		return -1
	}
	l.updateHead(v)
	return v.Value
}

func (l *LRUCache) updateHead(node *Node) {

}

func (l *LRUCache) Put(key int, value int) {
	l.values[key] = &Node{
		Key:   key,
		Value: value,
		Next:  l.Head,
	}

	l.Head.Prev = l.values[key]
	l.Head = l.values[key]

	if len(l.values) > l.cap {
		l.remove(l.Tail)
	}
}

func (l *LRUCache) remove(node *Node) {
	node.Prev.Next = nil
	delete(l.values, node.Key)
}
