package rotation

const (
	none = "none"
	quantity = "quantity"
)

type strategy interface {
	Rotate([]string, int) []string
}

type Rotation struct {
	strategy strategy
}

func NewStrategy() *Rotation {
	return &Rotation{}
}

func (c *Rotation) Use(strategy string) {
	switch strategy {
	case none:
		c.strategy = newNoneStrategy()
	case quantity:
		c.strategy = newQuantityStrategy()
	default:
		c.strategy = nil
	}
}

func (c *Rotation) Rotate(items []string, t int) []string {
	return c.strategy.Rotate(items, t)
}