package kafka

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
		Consumers   []Consumer `yaml:"consumers,flow"`
		Database
	}

	// Consumer kafka consumer
	Consumer struct {
		Name    string
		Brokers []string  `yaml:",flow"`
		GroupID string    `yaml:"groupId"`
		Topics  []string  `yaml:",flow"`
		SASL    SASL      `yaml:"SASL"`
		Remote  Remote    `yaml:"remote"`
		Log     LogConfig `yaml:"log"`
	}

	// SASL kafka SASL
	SASL struct {
		Enabled  bool
		User     string
		Password string
	}

	// Remote 远程调用
	Remote struct {
		// 回调地址
		URL       string
		QoS       QoS
		Heartbeat Heartbeat
	}

	// QoS 服务质量保证
	QoS struct {
		// 请求超时时间
		Timeout uint
		// 重试次数
		Retry uint
	}

	// Heartbeat 心跳机制，用于检测 kafka 服务与指定客户端是否保存联通
	Heartbeat struct {
		// 用于检查的远端地址
		HeathURL string
		// 心跳请求超时时间
		Timeout uint
		// 心跳频率（秒）
		Rate uint
	}

	// LogConfig 日志记录
	LogConfig struct {
		// 是否记录日志
		Enabled bool
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
