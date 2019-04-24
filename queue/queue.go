package queue

type Item interface {
}

type IitemQueue interface {
	Init() ItemQueue
	Enqueue(t Item)
	Dequeue() *Item
	IsEmpty() bool
	Size() int
}

// Item the type of the queue
type ItemQueue struct {
	items []Item
}


// New creates a new ItemQueue
func (s *ItemQueue) Init() *ItemQueue {
	s.items = []Item{}
	return s
}

// Enqueue adds an Item to the end of the queue
func (s *ItemQueue) Enqueue(t Item) {
	s.items = append(s.items, t)
}

// dequeue
func (s *ItemQueue) Dequeue() Item {
	item := s.items[0] // 先进先出
	s.items = s.items[1:len(s.items)]
	return item
}

func (s *ItemQueue) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of Items in the queue
func (s *ItemQueue) Size() int {
	return len(s.items)
}

func NewItemQueue() *ItemQueue {
	que := &ItemQueue{}
	que.Init()
	return que
}