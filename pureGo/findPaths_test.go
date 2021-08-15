package pureGo

import "testing"

func Test_findPaths(t *testing.T) {
	type args struct {
		m           int
		n           int
		maxMove     int
		startRow    int
		startColumn int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "findPathsT0",
			args: args{
				m:           36,
				n:           5,
				maxMove:     50,
				startRow:    15,
				startColumn: 3,
			},
			want: 390153306,
		},
		{
			name: "findPathsT1",
			args: args{
				m:           8,
				n:           7,
				maxMove:     16,
				startRow:    1,
				startColumn: 5,
			},
			want: 102984580,
		},
		{
			name: "findPathsT2",
			args: args{
				m:           45,
				n:           35,
				maxMove:     47,
				startRow:    20,
				startColumn: 31,
			},
			want: 951853874,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPaths(tt.args.m, tt.args.n, tt.args.maxMove, tt.args.startRow, tt.args.startColumn); got != tt.want {
				t.Errorf("findPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doRe(t *testing.T) {
	type args struct {
		m           int
		n           int
		maxMove     int
		startRow    int
		startColumn int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doRe(tt.args.m, tt.args.n, tt.args.maxMove, tt.args.startRow, tt.args.startColumn); got != tt.want {
				t.Errorf("doRe() = %v, want %v", got, tt.want)
			}
		})
	}
}
