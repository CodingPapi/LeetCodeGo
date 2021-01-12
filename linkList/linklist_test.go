package linkList

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
			l := BuildLinkList(tt.args.head)
			if got := RemoveDuplicateFromSortedList(l).ToString(); got != tt.want {
				t.Errorf("ListNode.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_RemoveDuplicateFromSortedListII(t *testing.T) {
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
			want: "2,3",
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
			want: "1,3",
		},
		{
			name: "basicTest2",
			args: args{
				head: []int{1, 2, 2, 2, 3},
			},
			want: "1,3",
		},
		{
			name: "basicTest2",
			args: args{
				head: []int{1, 1, 1, 2, 2, 2, 3},
			},
			want: "3",
		},
		{
			name: "basicTest3",
			args: args{
				head: []int{1, 1, 1, 1},
			},
			want: "",
		},
		{
			name: "basicTest4",
			args: args{
				head: []int{1, 1, 1},
			},
			want: "",
		},
		{
			name: "basicTest5",
			args: args{
				head: []int{0, 1, 2, 2, 3, 4},
			},
			want: "0,1,3,4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := BuildLinkList(tt.args.head)
			if got := RemoveDuplicateFromSortedListII(l).ToString(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicateFromSortedListII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseList(t *testing.T) {
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
				head: []int{1, 2, 3},
			},
			want: "3,2,1",
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
			want: "3,2,2,1",
		},
		{
			name: "basicTest2",
			args: args{
				head: []int{1},
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := BuildLinkList(tt.args.head)
			if got := ReverseList(l).ToString(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseListBetween(t *testing.T) {
	type args struct {
		head []int
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basicTest",
			args: args{
				head: []int{1, 2, 3},
				m:    1,
				n:    2,
			},
			want: "2,1,3",
		},
		{
			name: "basicTest0",
			args: args{
				head: []int{},
				m:    0,
				n:    0,
			},
			want: "",
		},
		{
			name: "basicTest1",
			args: args{
				head: []int{1, 2, 2, 3},
				m:    2,
				n:    3,
			},
			want: "1,2,2,3",
		},
		{
			name: "basicTest2",
			args: args{
				head: []int{1},
				m:    1,
				n:    1,
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := BuildLinkList(tt.args.head)
			if got := ReverseListBetween(l, tt.args.m, tt.args.n).ToString(); got != tt.want {
				t.Errorf("ReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeTwoSortedLists(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basicTest",
			args: args{
				l1: []int{1, 2, 3},
				l2: []int{1, 2},
			},
			want: "1,1,2,2,3",
		},
		{
			name: "basicTest0",
			args: args{
				l1: []int{},
				l2: []int{},
			},
			want: "",
		},
		{
			name: "basicTest1",
			args: args{
				l1: []int{1},
				l2: []int{},
			},
			want: "1",
		},
		{
			name: "basicTest2",
			args: args{
				l1: []int{},
				l2: []int{1},
			},
			want: "1",
		},
		{
			name: "basicTest3",
			args: args{
				l1: []int{1, 2},
				l2: []int{1, 2, 3},
			},
			want: "1,1,2,2,3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll1 := BuildLinkList(tt.args.l1)
			ll2 := BuildLinkList(tt.args.l2)
			if got := MergeTwoSortedLists(ll1, ll2).ToString(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoSortedLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeTwoSortedListsRecursion(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basicTest",
			args: args{
				l1: []int{1, 2, 3},
				l2: []int{1, 2},
			},
			want: "1,1,2,2,3",
		},
		{
			name: "basicTest0",
			args: args{
				l1: []int{},
				l2: []int{},
			},
			want: "",
		},
		{
			name: "basicTest1",
			args: args{
				l1: []int{1},
				l2: []int{},
			},
			want: "1",
		},
		{
			name: "basicTest2",
			args: args{
				l1: []int{},
				l2: []int{1},
			},
			want: "1",
		},
		{
			name: "basicTest3",
			args: args{
				l1: []int{1, 2},
				l2: []int{1, 2, 3},
			},
			want: "1,1,2,2,3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll1 := BuildLinkList(tt.args.l1)
			ll2 := BuildLinkList(tt.args.l2)
			if got := MergeTwoSortedListsRecursion(ll1, ll2).ToString(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoSortedLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartition(t *testing.T) {
	type args struct {
		l []int
		x int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basicTest",
			args: args{
				l: []int{1, 2, 3},
				x: 2,
			},
			want: "1,2,3",
		},
		{
			name: "basicTest0",
			args: args{
				l: []int{},
				x: 3,
			},
			want: "",
		},
		{
			name: "basicTest1",
			args: args{
				l: []int{1},
				x: 0,
			},
			want: "1",
		},
		{
			name: "basicTest2",
			args: args{
				l: []int{1},
				x: 2,
			},
			want: "1",
		},
		{
			name: "basicTest3",
			args: args{
				l: []int{1, 2, 4, 8, 6, 1},
				x: 4,
			},
			want: "1,2,1,4,8,6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll1 := BuildLinkList(tt.args.l)
			if got := Partition(ll1, tt.args.x).ToString(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoSortedLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMiddleAndCut(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basicTest",
			args: args{
				head: []int{1, 2, 3},
			},
			want: []string{"1", "2,3"},
		},
		{
			name: "basicTest0",
			args: args{
				head: []int{1, 2},
			},
			want: []string{"1", "2"},
		},
		{
			name: "basicTest1",
			args: args{
				head: []int{},
			},
			want: []string{},
		},
		{
			name: "basicTest2",
			args: args{
				head: []int{1},
			},
			want: []string{"1"},
		},
		{
			name: "basicTest3",
			args: args{
				head: []int{1, 2, 4, 8, 6, 1},
			},
			want: []string{"1,2,4", "8,6,1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := BuildLinkList(tt.args.head)
			right := FindMiddleAndCut(head)
			got := []string{}
			if head != nil {
				got = append(got, head.ToString())
			}
			if right != nil {
				got = append(got, right.ToString())
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMiddleAndCut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortList(t *testing.T) {
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
				head: []int{1, 2, 3},
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
				head: []int{3, 2, 2, 1},
			},
			want: "1,2,2,3",
		},
		{
			name: "basicTest2",
			args: args{
				head: []int{1},
			},
			want: "1",
		},
		{
			name: "basicTest1",
			args: args{
				head: []int{3, 2, 6, 1, 4},
			},
			want: "1,2,3,4,6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := BuildLinkList(tt.args.head)
			if got := SortList(head).ToString(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
