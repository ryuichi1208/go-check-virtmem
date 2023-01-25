package main

import "testing"

func Test_parseMemInfo(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "nomal01",
			args: args{
				path: "./t/test001.txt",
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "nomal02",
			args: args{
				path: "./t/test002.txt",
			},
			want:    50,
			wantErr: false,
		},
		{
			name: "err01",
			args: args{
				path: "./t/test00x.txt",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseMemInfo(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMemInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseMemInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseArgs(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := parseArgs(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("parseArgs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
