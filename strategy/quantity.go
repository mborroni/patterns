package rotation

type Quantity struct {}

func newQuantityStrategy() *Quantity {
	return &Quantity{}
}

func (q *Quantity) Rotate(items []string, times int) []string {
	// strategy logic
	return items
}