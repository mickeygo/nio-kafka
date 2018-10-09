package goka

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type (
	// Config goka config
	Config struct {
		Version     string
		Author      string
		Description string
		Consumers   []*Consumer `yaml:",flow"`
		Producers   []*Producer `yaml:",flow"`
		HTTPServer  `yaml:"http"`
		Database
	}

	// Consumer kafka consumer
	Consumer struct {
		Name    string
		Brokers []string `yaml:",flow"`
		Group   string
		Topics  []string `yaml:",flow"`
		SASL
	}

	// Producer kafka producer
	Producer struct {
		Name    string
		Brokers []string `yaml:",flow"`
		Topics  []string `yaml:",flow"`
		SASL
	}

	// SASL kafka SASL
	SASL struct {
		Enabled  bool
		User     string
		Password string
	}

	// HTTPServer http
	HTTPServer struct {
		Port   int16
		Prefix string
	}

	// Database database config
	Database struct {
		Host     string
		Catalog  string
		UserID   string
		Password string
	}
)

// Unmarshal Unmarshal to object by yaml format data
func Unmarshal(in []byte) (*Config, error) {
	out := Config{}
	err := yaml.Unmarshal(in, &out)
	return &out, err
}

// UnmarshalViaFile Unmarshal to object by a yml file
func UnmarshalViaFile(filename string) (*Config, error) {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return Unmarshal(in)
}
