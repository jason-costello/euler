package euler

import (
	"reflect"
	"testing"
)

func Test_problem3(t *testing.T) {
	type args struct {
		num     float64
		divisor float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "one",
			args: args{
				num:     13195.0,
				divisor: 2.0,
			},
			want: 29.0,
		},
		{
			name: "one",
			args: args{
				num:     600851475143.0,
				divisor: 2.0,
			},
			want: 6857.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem3(tt.args.num, tt.args.divisor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("problem3() = %v, want %v", got, tt.want)
			} else {
				t.Log(got)
			}
		})
	}
}
