package conf

import (
	"reflect"
	"testing"
)

func TestDecodeViaFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "config_test01",
			args: args{
				filename: "config_test.toml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeViaFile(tt.args.filename)
			t.Log(err)
			t.Logf("%v", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeViaFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeViaFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
