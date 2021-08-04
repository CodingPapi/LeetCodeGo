package randomPick

import "testing"
import "reflect"

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nextP1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{1, 3, 2},
		},
		{
			name: "nextP2",
			args: args{
				nums: []int{1, 1, 5},
			},
			want: []int{1, 5, 1},
		},
		{
			name: "nextP3",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "nextP4",
			args: args{
				nums: []int{4, 2, 0, 2, 3, 2, 0},
			},
			want: []int{4, 2, 0, 3, 0, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("nextPermutation() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
