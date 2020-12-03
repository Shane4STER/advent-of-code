package main

import (
	"reflect"
	"testing"
)

func Test_countTrees(t *testing.T) {
	type args struct {
		mapTemplate []string
		chosenSlope slope
	}
	testMap := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{testMap, slope{1, 1}}, 2},
		{"example 2", args{testMap, slope{3, 1}}, 7},
		{"example 3", args{testMap, slope{5, 1}}, 3},
		{"example 4", args{testMap, slope{7, 1}}, 4},
		{"example 5", args{testMap, slope{1, 2}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTrees(tt.args.mapTemplate, tt.args.chosenSlope); got != tt.want {
				t.Errorf("countTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextPos(t *testing.T) {
	type args struct {
		currentPos coord
		delta      slope
	}
	tests := []struct {
		name string
		args args
		want coord
	}{
		{"moveRightOnly", args{coord{0, 0}, slope{1, 0}}, coord{1, 0}},
		{"moveDownOnly", args{coord{0, 0}, slope{0, 1}}, coord{0, 1}},
		{"moveDownAndRight", args{coord{5, 5}, slope{5, 5}}, coord{10, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextPos(tt.args.currentPos, tt.args.delta); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextPos() = %v, want %v", got, tt.want)
			}
		})
	}
}
