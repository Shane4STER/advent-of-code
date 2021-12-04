package main

import "testing"

func Test_filterPrefix(t *testing.T) {
	type args struct {
		useFrequent bool
		haystack    []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"CO2 Scrubber Rating",
			args{
				false,
				[]string{
					"00100",
					"11110",
					"10110",
					"10111",
					"10101",
					"01111",
					"00111",
					"11100",
					"10000",
					"11001",
					"00010",
					"01010",
				},
			},
			"01010",
		},
		{
			"Oxygen Generator Rating",
			args{
				true,
				[]string{
					"00100",
					"11110",
					"10110",
					"10111",
					"10101",
					"01111",
					"00111",
					"11100",
					"10000",
					"11001",
					"00010",
					"01010",
				},
			},
			"10111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterPrefix(tt.args.useFrequent, tt.args.haystack); got != tt.want {
				t.Errorf("filterPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
