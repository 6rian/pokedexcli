package main

import (
	"reflect"
	"testing"
)

func Test_parseInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Input should be lowercased",
			args: args{"TESTING"},
			want: []string{"testing"},
		},
		{
			name: "Input should be trimmed",
			args: args{"  trim   "},
			want: []string{"trim"},
		},
		{
			name: "Input should be split",
			args: args{"catch pikachu"},
			want: []string{"catch", "pikachu"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
