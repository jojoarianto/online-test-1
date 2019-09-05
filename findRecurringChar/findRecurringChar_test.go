package findRecurringChar

import "testing"

func TestFind(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "skenario 1",
			args: args{
				in: "DBCABA",
			},
			want: "B",
		},
		{
			name: "skenario 2",
			args: args{
				in: "ABBA",
			},
			want: "B",
		},
		{
			name: "skenario 3",
			args: args{
				in: "ABCDE",
			},
			want: "",
		},
		{
			name: "skenario 4",
			args: args{
				in: "ABCDEA",
			},
			want: "A",
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.in); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
