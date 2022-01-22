package sets

type Sets[T comparable] map[T]bool

// T is the generic type it can be any comparable data type
// Constructor to create new set
// Example :-  New(int)() to create a int set
// New(string)() to create a string set

func New[T comparable]() Sets[T] {

	return make(Sets[T])
}

func (s Sets[T]) Add(values ...T) {

	for _, value := range values {
		s[value] = true
	}
}

func (s Sets[T]) Delete(values ...T) {
	for _, value := range values {
		delete(s, value)
	}
}

func (s Sets[T]) Len() int {

	return len(s)
}

func (s Sets[T]) Has(value T) bool {

	_, ok := s[value]

	return ok
}

func (s Sets[T]) Iterate(it func(T)) {

	for v := range s {
		it(v)
	}

}
