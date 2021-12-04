package main

import (
	"reflect"
	"testing"
)

func Test_parseCard(t *testing.T) {
	type args struct {
		cardStr []string
	}
	tests := []struct {
		name string
		args args
		want BingoCard
	}{
		{
			"basic card",
			args{
				[]string{
					"22 13 17 11  0",
					" 8  2 23  4 24",
					"21  9 14 16  7",
					" 6 10  3 18  5",
					" 1 12 20 15 19",
				},
			},
			BingoCard{
				false,
				[][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				[][]bool{
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseCard(tt.args.cardStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBingoCard_callNumber(t *testing.T) {
	type fields struct {
		hasWon  bool
		values  [][]int
		matches [][]bool
	}
	type args struct {
		number int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   BingoCard
	}{
		{
			"Call matching",
			fields{
				false,
				[][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				[][]bool{
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
				},
			},
			args{
				18,
			},
			BingoCard{
				false,
				[][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				[][]bool{
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, true, false},
					{false, false, false, false, false},
				},
			},
		},
		{
			"Call existing",
			fields{
				false,
				[][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				[][]bool{
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, true, false},
					{false, false, false, false, false},
				},
			},
			args{
				18,
			},
			BingoCard{
				false,
				[][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				[][]bool{
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, true, false},
					{false, false, false, false, false},
				},
			},
		},
		{
			"Call non-matching",
			fields{
				false,
				[][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				[][]bool{
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
				},
			},
			args{
				99,
			},
			BingoCard{
				false,
				[][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				[][]bool{
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := &BingoCard{
				values:  tt.fields.values,
				matches: tt.fields.matches,
			}
			card.callNumber(tt.args.number)
			if !reflect.DeepEqual(*card, tt.want) {
				t.Errorf("card.callNumber() = %v, want %v", card, tt.want)
			}
		})
	}
}

func TestBingoCard_isWinner(t *testing.T) {
	type fields struct {
		hasWon  bool
		values  [][]int
		matches [][]bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"Row 0 Win",
			fields{
				false,
				[][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				[][]bool{
					{true, true, true, true, true},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
				},
			},
			true,
		},
		{
			"Col 2 Win",
			fields{
				false,
				[][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				[][]bool{
					{false, false, true, false, false},
					{false, false, true, false, false},
					{false, false, true, false, false},
					{false, false, true, false, false},
					{false, false, true, false, false},
				},
			},
			true,
		},
		{
			"No win",
			fields{
				false,
				[][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				[][]bool{
					{false, true, true, true, true},
					{true, false, true, true, true},
					{true, true, false, true, true},
					{true, true, true, false, true},
					{true, true, true, true, false},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := &BingoCard{
				values:  tt.fields.values,
				matches: tt.fields.matches,
			}
			if got := card.isWinner(); got != tt.want {
				t.Errorf("BingoCard.isWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBingoCard_score(t *testing.T) {
	type fields struct {
		hasWon  bool
		values  [][]int
		matches [][]bool
	}
	type args struct {
		lastCall int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			"Example",
			fields{
				false,
				[][]int{
					{14, 21, 17, 24, 4},
					{10, 16, 15, 9, 19},
					{18, 8, 23, 26, 20},
					{22, 11, 13, 6, 5},
					{2, 0, 12, 3, 7},
				},
				[][]bool{
					{true, true, true, true, true},
					{false, false, false, true, false},
					{false, false, true, false, false},
					{false, true, false, false, true},
					{true, true, false, false, true},
				},
			},
			args{
				24,
			},
			4512,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := &BingoCard{
				values:  tt.fields.values,
				matches: tt.fields.matches,
			}
			if got := card.score(tt.args.lastCall); got != tt.want {
				t.Errorf("BingoCard.score() = %v, want %v", got, tt.want)
			}
		})
	}
}
