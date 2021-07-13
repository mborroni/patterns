package rotation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	times = 0
)

var (
	dummyItems = []string{"item1", "item2", "item3"}
)

func TestStrategyPattern(t *testing.T) {
	ass := assert.New(t)
	rotation := NewStrategy()

	type fields struct {
		strategy string
		items    []string
		times    int
	}

	type want struct {
		items []string
	}

	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{name: "none",
			fields: fields{
				strategy: none,
				items:    dummyItems,
				times: times,
			},
			want: want{
				items: dummyItems,
			},
		},
		{name: "none",
			fields: fields{
				strategy: quantity,
				items:    dummyItems,
			},
			want: want{
				items: dummyItems,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotation.Use(tt.fields.strategy)
			items := rotation.Rotate(tt.fields.items, tt.fields.times)
			ass.Equal(tt.want.items, items)
		})
	}
}
