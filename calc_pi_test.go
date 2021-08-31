package main

import "testing"

func Test_estimate_pi(t *testing.T) {
	type args struct {
		num_iterations int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := estimate_pi(tt.args.num_iterations); got != tt.want {
				t.Errorf("estimate_pi() = %v, want %v", got, tt.want)
			}
		})
	}
}
