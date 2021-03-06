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
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
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
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
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
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
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
				btData: []int{1,ut.NULL,1},
			},
			want: false,
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
		{
			name: "basicTest4",
			args: args {
				btData: []int{5,1,7,ut.NULL,ut.NULL,6,8},
			},
			want: true,
		},
		{
			name: "basicTest5",
			args: args {
				btData: []int{3,1,5,0,2,4,6,ut.NULL,ut.NULL,ut.NULL,3},
			},
			want: false,
		},
		{
			name: "basicTest6",
			args: args {
				btData: []int{3,ut.NULL,30,10,ut.NULL,ut.NULL,15,ut.NULL,45},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ut.BuildBinaryTree(tt.args.btData)
			if got := isValidBSTFast(root); got != tt.want{
				t.Errorf("isValidBST = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertIntoBst(t *testing.T) {
	type args struct {
		btData []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nilTest",
			args: args {
				btData: []int{},
				target: 1,
			},
			want: []int{1},
		},
		{
			name: "basicTest",
			args: args {
				btData: []int{1,2,3},
				target: 4,
			},
			want: []int{1,2,3,ut.NULL,ut.NULL,ut.NULL,4},
		},
		{
			name: "basicTest",
			args: args {
				btData: []int{4,2,7,1,3},
				target: 5,
			},
			want: []int{4,2,7,1,3,5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ut.BuildBinaryTree(tt.args.btData)
			want := ut.BuildBinaryTree(tt.want)
			if got := insertIntoBst(root, tt.args.target); !reflect.DeepEqual(got, want) {
				t.Errorf("isValidBST = %v, want %v", got, tt.want)
			}
		})
	}
}
