package lruCache

type LRUCache struct {
	capacity int
	cache map[int]int
	head *node
	tail *node
}


type node struct {
	value int
	next *node
}

func Constructor(capacity int) LRUCache {
	newNode := &node{-1, nil}
	return LRUCache{
		capacity,
		make(map[int]int),
		newNode,
		newNode,
	}
}

func (l *LRUCache) Get(key int) int {
	value, ok := l.cache[key]
	if ok {
		l.sort(key)
		return value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if _, ok := l.cache[key]; ok {
		l.sort(key)
		l.cache[key] = value
	} else {
		if l.capacity == len(l.cache) {
			index := l.head.next.value
			delete(l.cache, index)
			l.removeFirst()
		}
		l.append(key)
		l.cache[key] = value
	}
}

func (l *LRUCache) sort(key int) {
	prev := l.head
	for prev.next.value != key {
		prev = prev.next
	}
	if prev.next.next != nil {
		l.tail.next = prev.next
		l.tail = prev.next
		prev.next = l.tail.next
		l.tail.next = nil
	}
}

func (l *LRUCache) removeFirst() {
	l.head.next = l.head.next.next
	if l.head.next == nil {
		l.tail = l.head
	}
}

func (l *LRUCache) append(key int) {
	newNode := &node{key, nil}
	l.tail.next = newNode
	l.tail = newNode
}
