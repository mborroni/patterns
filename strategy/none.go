package rotation

type None struct {}

func newNoneStrategy() *None {
	return &None{}
}

func (t *None) Rotate(items []string, _ int) []string {
	// strategy logic
	return items
}