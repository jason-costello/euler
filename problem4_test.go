package euler

import "testing"

func Test_isPalindromic(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{n: 1001},
			want: true,

		},
		{
			name: "one",
			args: args{n: 10001},
			want: true,

		},
		{
			name: "one",
			args: args{n: 1011},
			want: false,

		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromic(tt.args.n); got != tt.want {
				t.Errorf("isPalindromic() = %v, want %v", got, tt.want)
			}
		})
	}
}



func Test_prob41(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "one",
			args: args{
				min: 10,
				max: 99,
			},
			want: 9009,
		},
		{
			name: "one",
			args: args{
				min: 100,
				max: 999,
			},
			want: 9009,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prob4(tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("prob4() = %v, want %v", got, tt.want)
			}
		})
	}
}
