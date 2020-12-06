package main

import (
	"testing"
)

func Test_processPaxGroup(t *testing.T) {
	type args struct {
		chunk                 []string
		countGroupAnswersOnly bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Single answer",
			args{
				[]string{
					"abcd",
				},
				false,
			},
			4,
		},
		{
			"Empty chunk",
			args{
				[]string{},
				false,
			},
			0,
		},
		{
			"Multi chunk",
			args{
				[]string{
					"abcd",
					"abcd",
					"abcd",
				},
				false,
			},
			4,
		},
		{
			"Repeats",
			args{
				[]string{
					"aaaa",
					"abbcc",
					"def",
				},
				false,
			},
			6,
		},
		{
			"Single Person",
			args{
				[]string{
					"abc",
				},
				false,
			},
			3,
		},
		{
			"Group different",
			args{
				[]string{
					"a",
					"b",
					"c",
				},
				true,
			},
			0,
		},
		{
			"group similar",
			args{
				[]string{
					"ab",
					"ac",
				},
				true,
			},
			1,
		},
		{
			"group same",
			args{
				[]string{
					"a",
					"a",
					"a",
					"a",
				},
				true,
			},
			1,
		},
		{
			"single",
			args{
				[]string{
					"a",
				},
				true,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processPaxGroup(tt.args.chunk, tt.args.countGroupAnswersOnly); got != tt.want {
				t.Errorf("processChunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapValueIsTargetValue(t *testing.T) {
	type args struct {
		data        map[rune]int
		targetValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"empty map",
			args{
				map[rune]int{},
				0,
			},
			0,
		},
		{
			"single found",
			args{
				map[rune]int{
					'a': 1,
				},
				1,
			},
			1,
		},
		{
			"multiple found",
			args{
				map[rune]int{
					'a': 1,
					'b': 2,
					'c': 2,
					'd': 1,
					'e': 1,
				},
				1,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapValueIsTargetValue(tt.args.data, tt.args.targetValue); got != tt.want {
				t.Errorf("mapValueIsTargetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
