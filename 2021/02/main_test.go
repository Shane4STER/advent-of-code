package main

import (
	"reflect"
	"testing"
)

func Test_stringToCommand(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    Command
		wantErr bool
	}{
		{
			"Forward",
			args{
				"forward 4",
			},
			Command{
				FORWARD,
				4,
			},
			false,
		},
		{
			"Up",
			args{
				"up 5",
			},
			Command{
				UP,
				5,
			},
			false,
		},
		{
			"Down",
			args{
				"down 6",
			},
			Command{
				DOWN,
				6,
			},
			false,
		},
		{
			"Error",
			args{
				"backward 4",
			},
			Command{
				0,
				0,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stringToCommand(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("stringToCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringToCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_move(t *testing.T) {
	type fields struct {
		x   int
		d   int
		aim int
	}
	type args struct {
		cmd Command
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Position
	}{
		{
			"Forward",
			fields{0, 0, 0},
			args{Command{FORWARD, 5}},
			Position{5, 0, 0},
		},
		{
			"Down",
			fields{5, 0, 0},
			args{Command{DOWN, 5}},
			Position{5, 0, 5},
		},
		{
			"Forward and Down",
			fields{5, 0, 5},
			args{Command{FORWARD, 8}},
			Position{13, 40, 5},
		},
		{
			"Up",
			fields{13, 40, 5},
			args{Command{UP, 3}},
			Position{13, 40, 2},
		},
		{
			"Down Again",
			fields{13, 40, 2},
			args{Command{DOWN, 8}},
			Position{13, 40, 10},
		},
		{
			"Forward",
			fields{13, 40, 10},
			args{Command{FORWARD, 2}},
			Position{15, 60, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos := &Position{
				x:   tt.fields.x,
				d:   tt.fields.d,
				aim: tt.fields.aim,
			}
			pos.move(tt.args.cmd)
			if !reflect.DeepEqual(*pos, tt.want) {
				t.Errorf("pos.move() = %v, want %v", pos, tt.want)
			}
		})
	}
}

func TestPosition_hPos(t *testing.T) {
	type fields struct {
		x   int
		d   int
		aim int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"Basic Test",
			fields{15, 60, 10},
			900,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos := &Position{
				x:   tt.fields.x,
				d:   tt.fields.d,
				aim: tt.fields.aim,
			}
			if got := pos.hPos(); got != tt.want {
				t.Errorf("Position.hPos() = %v, want %v", got, tt.want)
			}
		})
	}
}
