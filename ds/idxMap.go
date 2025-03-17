package ds

type IdxMap[T comparable] struct {
	data []*T
	set  map[*T]int
}

func NewIdxMap[T comparable]() IdxMap[T] {
	return IdxMap[T]{
		data: []*T{},
		set:  make(map[*T]int),
	}
}

func (this *IdxMap[T]) Add(el *T) {
	if _, exists := this.set[el]; exists {
		return
	}

	this.set[el] = len(this.data)
	this.data = append(this.data, el)
}

func (this IdxMap[T]) Has(el *T) bool {
	_, exists := this.set[el]
	return exists
}

func (this IdxMap[T]) Get(i int) *T {
	return this.data[i]
}

func (this *IdxMap[T]) Remove(el *T) {
	idx, exists := this.set[el]
	if !exists {
		return
	}

	lastIdx := len(this.data) - 1
	lastElm := this.data[lastIdx]

	this.data[idx] = lastElm
	this.set[lastElm] = idx
	this.data = this.data[:lastIdx]

	delete(this.set, el)
}

func (this *IdxMap[T]) Iter() []*T {
	return this.data
}

func (this IdxMap[T]) Len() int {
	return len(this.data)
}
