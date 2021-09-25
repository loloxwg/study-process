package main

import "testing"

func Test_maxMoney(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"11111111111111111111",
			args{
				12,
			},
			47,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxMoney(tt.args.n); got != tt.want {
				t.Errorf("maxMoney() = %v, want %v", got, tt.want)
			}
		})
	}
}
