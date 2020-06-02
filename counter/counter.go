package counter

type Counter interface {
	Increment()
	Decrement()
	Count() int
}
