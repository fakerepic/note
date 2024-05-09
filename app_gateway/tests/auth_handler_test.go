package tests

import (
	"pbapp/hooks/proxy/mw"
	"testing"
)

func TestGetIDFromPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				path: "/userdb-1234567890",
			},
			want: "1234567890",
		},
		{
			name: "test2",
			args: args{
				path: "/userdb-1234567890/asdf",
			},
			want: "1234567890",
		},
		{
			name: "test3",
			args: args{
				path: "/userdb_1234567890",
			},
			want: "",
		},
		{
			name: "test4",
			args: args{
				path: "/userdb_1234567890/asdf",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mw.GetIDFromPath(tt.args.path); got != tt.want {
				t.Errorf("GetIDFromPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
