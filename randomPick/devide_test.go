package randomPick

import "testing"

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "divideTest0",
			args: args{
				dividend: 2147483647,
				divisor:  2,
			},
			want: 1073741823,
		},
		{
			name: "divideTest1",
			args: args{
				dividend: 10,
				divisor:  3,
			},
			want: 3,
		},
		{
			name: "divideTest2",
			args: args{
				dividend: 2,
				divisor:  2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
