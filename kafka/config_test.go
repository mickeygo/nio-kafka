package kafka

import (
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Unmarshal(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshalViaFile(t *testing.T) {
	conf, err := UnmarshalViaFile("../config.yml")
	if err != nil {
		t.Errorf("Unmarsha error: %v ", err)
	}

	if conf.Author != "gang.yang" {
		t.Errorf("conf.Author[%s] == gang.yang", conf.Author)
	}
}
