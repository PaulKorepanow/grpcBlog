package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"simpleCase",
			args{
				a: []int{1, 2, 3, 4, 5, 6},
				b: []int{3, 123, 1},
			},
			[]int{3, 1},
		},
		{
			"nilCase",
			args{
				a: []int{1, 2, 3, 4, 5, 6},
				b: nil,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := intersection(tt.args.a, tt.args.b)
			assert.ElementsMatch(t, res, tt.want)
		})
	}
}
