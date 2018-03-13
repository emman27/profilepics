package imaging

import "testing"

func Test_initials(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{name: "Rakesh"}, want: "R"},
		{args: args{name: "Emmanuel Goh"}, want: "EG"},
		{args: args{name: "Xu yangbo"}, want: "XY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initials(tt.args.name); got != tt.want {
				t.Errorf("initials() = %v, want %v", got, tt.want)
			}
		})
	}
}
