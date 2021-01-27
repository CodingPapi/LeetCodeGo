package stackAndQueue

import (
	"reflect"
	"testing"

	. "github.com/CodingPapi/LeetCodeGo/util"
)

func Test_RemoveDuplicateFromSortedList(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basicTest",
			args: args{
				head: []int{1, 1, 2, 3},
			},
			want: "1,2,3",
		},
		{
			name: "basicTest0",
			args: args{
				head: []int{},
			},
			want: "",
		},
		{
			name: "basicTest1",
			args: args{
				head: []int{1, 2, 2, 3},
			},
			want: "1,2,3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testS := Constructor()
			testS.Push(-2)
			testS.Push(0)
			testS.Push(-3)
			testS.GetMin()
			testS.Pop()
			testS.Top()
			testS.GetMin()
			// l := BuildLinkList(tt.args.head)
			// if got := RemoveDuplicateFromSortedList(l).ToString(); got != tt.want {
			// 	t.Errorf("ListNode.toString() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestEvalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				tokens: []string{"1", "2", "+"},
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				tokens: []string{"1", "2", "3", "*", "+"},
			},
			want: 7,
		},
		{
			name: "test3",
			args: args{
				tokens: []string{"1", "4", "3", "/", "+"},
			},
			want: 2,
		},
		{
			name: "test3",
			args: args{
				tokens: []string{"4", "13", "5", "/", "+"},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvalRPN(tt.args.tokens); got != tt.want {
				t.Errorf("EvalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				s: "3[a]2[bc]",
			},
			want: "aaabcbc",
		},
		{
			name: "test2",
			args: args{
				s: "3[a2[c]]",
			},
			want: "accaccacc",
		},
		{
			name: "test3",
			args: args{
				s: "2[abc]3[cd]ef",
			},
			want: "abcabccdcdcdef",
		},
		{
			name: "test3",
			args: args{
				s: "abc3[cd]xyz",
			},
			want: "abccdcdcdxyz",
		},
		{
			name: "test4",
			args: args{
				s: "ab2[c3[m]x2[d]y]z",
			},
			want: "abcmmmxddycmmmxddyz",
		},
		{
			name: "test5",
			args: args{
				s: "10[a]2[bc]",
			},
			want: "aaaaaaaaaabcbc",
		},
		{
			name: "test6",
			args: args{
				s: "2[a]2[2[b]c]",
			},
			want: "aabbcbbc",
		},
		{
			name: "test7",
			args: args{
				s: "3[z]2[2[y]pq4[2[jk]e1[f]]]ef",
			},
			want: "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeString(tt.args.s); got != tt.want {
				t.Errorf("DecodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInroderTraversal(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basicTest",
			args: args{
				data: []int{1, 2, 3},
			},
			want: []int{2, 1, 3},
		},
		{
			name: "basicTest2",
			args: args{
				data: []int{1, 2, 4, NULL, -6},
			},
			want: []int{2, -6, 1, 4},
		},
		{
			name: "negativeTest",
			args: args{
				data: []int{-6},
			},
			want: []int{-6},
		},
	}
	for _, tt := range tests {
		root := BuildBinaryTree(tt.args.data)
		t.Run(tt.name, func(t *testing.T) {
			if got := InroderTraversal(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InroderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumIsLands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basicTest",
			args: args{
				grid: [][]byte{{}},
			},
			want: 0,
		},
		{
			name: "basicTest1",
			args: args{
				grid: [][]byte{
					{'1', '1', '1'},
					{'0', '1', '0'},
					{'1', '1', '1'},
				},
			},
			want: 1,
		},
		{
			name: "basicTest1",
			args: args{
				grid: [][]byte{
					{'1', '1', '1', '1', '0'},
					{'1', '1', '0', '1', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumIsLands(tt.args.grid); got != tt.want {
				t.Errorf("NumIsLands() = %v, want %v", got, tt.want)
			}
		})
	}
}
