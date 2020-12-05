package main

import (
	"testing"
)

func Test_getRow(t *testing.T) {
	type args struct {
		seatString string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"BasicTest",
			args{"FBFBBFFRLR"},
			44,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRow(tt.args.seatString); got != tt.want {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getColumn(t *testing.T) {
	type args struct {
		seatString string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"BasicTest",
			args{"FBFBBFFRLR"},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getColumn(tt.args.seatString); got != tt.want {
				t.Errorf("getColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_seat_getID(t *testing.T) {
	type fields struct {
		row    int64
		column int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			"example 1",
			fields{
				70,
				7,
			},
			567,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := seat{
				row:    tt.fields.row,
				column: tt.fields.column,
			}
			if got := s.getID(); got != tt.want {
				t.Errorf("seat.getID() = %v, want %v", got, tt.want)
			}
		})
	}
}
