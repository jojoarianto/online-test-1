package findMissingInteger

import "testing"

func TestSolutionFindMissingInteger(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "scenario 1",
			args: args{ A: []int{1,3,6,10,11,15} },
			want: 2,
		},
		{
			name: "scenario 2",
			args: args{ A: []int{1,1,1,1} },
			want: 5,
		},
		{
			name: "scenario 3",
			args: args{ A: []int{1,1,3,4} },
			want: 10,
		},
		{
			name: "scenario 4",
			args: args{ A: []int{2} },
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolutionFindMissingInteger(tt.args.A); got != tt.want {
				t.Errorf("SolutionFindMissingInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
