package main

import "testing"

func Test_isValidPassword(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example 1", args{"1-3 a: abcde"}, true},
		{"example 2", args{"1-3 b: cdefg"}, false},
		{"example 3", args{"2-9 c: ccccccccc"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidPassword(tt.args.input); got != tt.want {
				t.Errorf("isValidPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
