package randomPick

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{s: "2147483646"},
			want: 2147483646,
		},
		{
			name: "test2",
			args: args{s: "-91283472332"},
			want: -2147483648,
		},
		{
			name: "test3",
			args: args{s: "  +021"},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
