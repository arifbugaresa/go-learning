package common

import "testing"

func TestIsEmptyField(t *testing.T) {
	type args struct {
		v interface{}
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test integer",
			args: args{
				v: 12,
			},
			want: false,
		},
		{
			name: "test empty integer",
			args: args{
				v: 0,
			},
			want: true,
		},
		{
			name: "test float",
			args: args{
				v: 12.00,
			},
			want: false,
		},
		{
			name: "test empty float",
			args: args{
				v: 0,
			},
			want: true,
		},
		{
			name: "test string",
			args: args{
				v: "test",
			},
			want: false,
		},
		{
			name: "test empty string",
			args: args{
				v: "",
			},
			want: true,
		},
		{
			name: "test empty interface",
			args: args{
				v: interface{}(nil),
			},
			want: true,
		},
		{
			name: "test nil",
			args: args{
				v: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmptyField(tt.args.v); got != tt.want {
				t.Errorf("IsEmptyField() = %v, want %v", got, tt.want)
			}
		})
	}
}
