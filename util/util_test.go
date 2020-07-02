package util

import (
	"reflect"
	"testing"

	bt "github.com/CodingPapi/LeetCodeGo/binaryTree"
)

func Test_buildBinaryTree(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want *bt.TreeNode
	}{
		{
			name: "basic test",
			args: args {
				input: []int {1, 2},
			},
			want: &bt.TreeNode{Val: 1, Left: &bt.TreeNode{Val: 2, Left: nil, Right: nil}, Right: nil},
		},
		{
			name: "basic test1",
			args: args {
				input: []int {1, 2, 3, bt.NULL, 4},
			},
			want: &bt.TreeNode{Val: 1, Left: &bt.TreeNode{Val: 2, Left: nil, Right: &bt.TreeNode{Val: 4}}, Right: &bt.TreeNode{Val: 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildBinaryTree(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
