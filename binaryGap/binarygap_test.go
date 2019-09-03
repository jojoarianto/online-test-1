package binaryGap

import (
	"testing"
)

func Test_convertToBinary(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "skenario 1",
			args: args{
				input: 4,
			},
			want: "100",
		},
		{
			name: "skenario 2",
			args: args{
				input: 7,
			},
			want: "111",
		},
		{
			name: "skenario 3",
			args: args{
				input: 1041,
			},
			want: "10000010001",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBinary(tt.args.input); got != tt.want {
				t.Errorf("convertToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindGap(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "skenario 1",
			args: args{
				input: 1041,
			},
			want: 5,
		},
		{
			name: "skenario 2",
			args: args{
				input: 529,
			},
			want: 4,
		},
		{
			name: "skenario 2",
			args: args{
				input: 32,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindGap(tt.args.input); got != tt.want {
				t.Errorf("FindGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
