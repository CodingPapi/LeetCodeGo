package stackAndQueue

import (
	"testing"
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
