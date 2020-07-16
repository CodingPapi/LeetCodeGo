package util

import (
	"reflect"
	"testing"
)

func Test_buildBinaryTree(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "basic test",
			args: args {
				input: []int {1, 2},
			},
			want: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: nil},
		},
		{
			name: "basic test1",
			args: args {
				input: []int {1, 2, 3, NULL, 4},
			},
			want: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildBinaryTree(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
