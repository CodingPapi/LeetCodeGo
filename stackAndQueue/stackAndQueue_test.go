package stackAndQueue

import (
	"testing"
	// . "github.com/CodingPapi/LeetCodeGo/util"
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
