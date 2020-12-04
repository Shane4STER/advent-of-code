package main

import (
	"reflect"
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		pp       passport
		allowNPC bool
	}
	validPassport := passport{
		1990,
		2019,
		2025,
		height{
			180,
			"cm",
		},
		"brn",
		"#ffffff",
		860033327,
		101,
	}
	validNPC := passport{
		1985,
		2018,
		2022,
		height{
			75,
			"in",
		},
		"amb",
		"#fffffd",
		333278600,
		0,
	}
	var invalidPassport passport
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid passport", args{validPassport, false}, true},
		{"valid passport", args{validPassport, true}, true},
		{"valid NPC Accept", args{validNPC, true}, true},
		{"valid NPC Reject", args{validNPC, false}, false},
		{"invalid passport", args{invalidPassport, false}, false},
		{"invalid passport", args{invalidPassport, true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.pp, tt.args.allowNPC); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parsePassport(t *testing.T) {
	type args struct {
		input string
	}
	examplePassport := passport{
		1990,
		2019,
		2025,
		height{
			180,
			"cm",
		},
		"brn",
		"#fffffd",
		860033327,
		101,
	}
	tests := []struct {
		name string
		args args
		want passport
	}{
		{"example 1", args{"ecl:brn pid:860033327 eyr:2025 hcl:#fffffd byr:1990 iyr:2019 cid:101 hgt:180cm"}, examplePassport},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsePassport(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePassport() = %v, want %v", got, tt.want)
			}
		})
	}
}
