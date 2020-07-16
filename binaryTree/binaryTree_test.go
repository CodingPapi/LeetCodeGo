package binaryTree

import (
	"reflect"
	"testing"

	ut "github.com/CodingPapi/LeetCodeGo/util"
)

func Test_maxPathSum(t *testing.T) {
	type args struct {
		btData []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basicTest",
			args: args {
				btData: []int{1,2,3},
			},
			want: 6,
		},
		{
			name: "basicTest2",
			args: args {
				btData: []int{1,2,4,ut.NULL,-6},
			},
			want: 7,
		},
		{
			name: "negativeTest",
			args: args {
				btData: []int{-6},
			},
			want: -6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ut.BuildBinaryTree(tt.args.btData)
			if got := maxPathSum(root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		btData []int
		p []int
		q []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basicTest",
			args: args {
				btData: []int{1,2,3},
				p: []int{2},
				q: []int{3},
			},
			want: []int{1,2,3},
		},
		{
			name: "basicTest1",
			args: args {
				btData: []int{3,5,1,6,2,0,8,ut.NULL,ut.NULL,7,4},
				p: []int{5},
				q: []int{1},
			},
			want: []int{3,5,1,6,2,0,8,ut.NULL,ut.NULL,7,4},
		},
		{
			name: "basicTest2",
			args: args {
				btData: []int{3,5,1,6,2,0,8,ut.NULL,ut.NULL,7,4},
				p: []int{7},
				q: []int{4},
			},
			want: []int{2,7,4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ut.BuildBinaryTree(tt.args.btData)
			want := ut.BuildBinaryTree(tt.want)
			if got := lowestCommonAncestor(root, ut.BuildBinaryTree(tt.args.p), ut.BuildBinaryTree(tt.args.q)); !reflect.DeepEqual(got, want){
				t.Errorf("maxPathSum() = %v, want %v", got, want)
			}
		})
	}

}

func Test_levelOrder(t *testing.T) {
	type args struct {
		btData []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "nilTest",
			args: args {
				btData: []int{},
			},
			want: [][]int{},
		},
		{
			name: "basicTest",
			args: args {
				btData: []int{1,2,3},
			},
			want: [][]int{{1},{2,3}},
		},
		{
			name: "basicTest1",
			args: args {
				btData: []int{3,9,20,ut.NULL,ut.NULL,15,7},
			},
			want: [][]int{{3},{9,20},{15,7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ut.BuildBinaryTree(tt.args.btData)
			if got := levelOrder(root); !reflect.DeepEqual(got, tt.want){
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_levelOrderBottom(t *testing.T) {
	type args struct {
		btData []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "nilTest",
			args: args {
				btData: []int{},
			},
			want: [][]int{},
		},
		{
			name: "basicTest",
			args: args {
				btData: []int{1,2,3},
			},
			want: [][]int{{2,3},{1}},
		},
		{
			name: "basicTest1",
			args: args {
				btData: []int{3,9,20,ut.NULL,ut.NULL,15,7},
			},
			want: [][]int{{15,7},{9,20},{3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ut.BuildBinaryTree(tt.args.btData)
			if got := levelOrderBottom(root); !reflect.DeepEqual(got, tt.want){
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_zigzagLevelOrder(t *testing.T) {
	type args struct {
		btData []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "nilTest",
			args: args {
				btData: []int{},
			},
			want: [][]int{},
		},
		{
			name: "basicTest",
			args: args {
				btData: []int{1,2,3},
			},
			want: [][]int{{1},{3,2}},
		},
		{
			name: "basicTest1",
			args: args {
				btData: []int{3,9,20,ut.NULL,ut.NULL,15,7},
			},
			want: [][]int{{3},{20,9},{15,7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ut.BuildBinaryTree(tt.args.btData)
			if got := zigzagLevelOrder(root); !reflect.DeepEqual(got, tt.want){
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBST(t *testing.T) {
	type args struct {
		btData []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nilTest",
			args: args {
				btData: []int{},
			},
			want: true,
		},
		{
			name: "basicTest",
			args: args {
				btData: []int{1,2,3},
			},
			want: false,
		},
		{
			name: "basicTest1",
			args: args {
				btData: []int{2,1,3},
			},
			want: true,
		},
		{
			name: "basicTest2",
			args: args {
				btData: []int{5,1,4,ut.NULL,ut.NULL,3,6},
			},
			want: false,
		},
		{
			name: "basicTest3",
			args: args {
				btData: []int{5,1,7,ut.NULL,ut.NULL,3,8},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ut.BuildBinaryTree(tt.args.btData)
			if got := isValidBST(root); got != tt.want{
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
