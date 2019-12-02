package main

import (
	"reflect"
	"testing"
)

func Test_RunProgram(t *testing.T) {
	type args struct {
		program []int
		index   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 0, 0, 0, 99}, 0}, []int{2, 0, 0, 0, 99}},
		{"2", args{[]int{2, 3, 0, 3, 99}, 0}, []int{2, 3, 0, 6, 99}},
		{"3", args{[]int{2, 4, 4, 5, 99, 0}, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{"4", args{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, 0}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunProgram(tt.args.program, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runProgram() = %v, want %v", got, tt.want)
			}
		})
	}
}
