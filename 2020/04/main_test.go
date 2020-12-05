package main

import (
	"reflect"
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		pp idCard
	}
	tests := []struct {
		name string
		args args
		want idCardCheck
	}{
		{
			"Valid Passport",
			args{
				idCard{
					1920,
					2010,
					2020,
					height{
						150,
						"cm",
					},
					"amb",
					"#ffffff",
					"123456789",
					101,
				},
			},
			idCardCheck{
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
			},
		},
		{
			"Valid North Pole ID",
			args{
				idCard{
					1920,
					2010,
					2020,
					height{
						150,
						"cm",
					},
					"amb",
					"#ffffff",
					"123456789",
					0,
				},
			},
			idCardCheck{
				false,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				false,
			},
		},
		{
			"invalid",
			args{
				idCard{
					1919,
					2009,
					2031,
					height{
						77,
						"in",
					},
					"blk",
					"#GHIJKL",
					"0123",
					0,
				},
			},
			idCardCheck{
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.pp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseIDCard(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want idCard
	}{
		{
			"basic NPID imperial",
			args{
				"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f",
			},
			idCard{
				1980,
				2012,
				2030,
				height{
					74,
					"in",
				},
				"grn",
				"#623a2f",
				"087499704",
				0,
			},
		},
		{
			"basic passport metric",
			args{
				"eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
			},
			idCard{
				1989,
				2014,
				2029,
				height{
					165,
					"cm",
				},
				"blu",
				"#a97842",
				"896056539",
				129,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseIDCard(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseIDCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testChunk(t *testing.T) {
	type args struct {
		chunk       []string
		allowIDCard bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"valid Passport",
			args{
				[]string{
					"hcl:#888785",
					"hgt:164cm byr:2001 iyr:2015 cid:88",
					"pid:545766238 ecl:hzl",
					"eyr:2022",
				},
				false,
			},
			true,
		},
		{
			"valid IDCard acceptable",
			args{
				[]string{
					"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
					"hcl:#623a2f",
				},
				true,
			},
			true,
		},
		{
			"valid IDCard unacceptable",
			args{
				[]string{
					"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
					"hcl:#623a2f",
				},
				false,
			},
			false,
		},
		{
			"invalid Passport",
			args{
				[]string{
					"hcl:#888785",
					"hgt:164cm byr:2001 iyr:2015 cid:88",
					"pid:5457662380 ecl:hzl",
					"eyr:2022",
				},
				false,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testChunk(tt.args.chunk, tt.args.allowIDCard); got != tt.want {
				t.Errorf("testChunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
