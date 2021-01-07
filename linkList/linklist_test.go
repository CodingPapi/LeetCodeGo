package linkList

import (
	ut "github.com/CodingPapi/LeetCodeGo/util"
	"testing"
)

func TestListNode_toString(t *testing.T) {
	type fields struct {
		Val  int
		Next *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if got := l.toString(); got != tt.want {
				t.Errorf("ListNode.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_RemoveDuplicateFromSortedList(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := ut.BuildLinkList(tt.args.head)
			if got := RemoveDuplicateFromSortedList(l).toString(); got != tt.want {
				t.Errorf("ListNode.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}
