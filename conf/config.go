package conf

import (
	"github.com/BurntSushi/toml"
)

type (
	// Config 配置文件信息
	Config struct {
		// Consumers kafka 消费者集合
		Consumers []ConsumerConfig
	}

	// ConsumerConfig kafka 消费者配置
	ConsumerConfig struct {
		// Name 配置的消费者名称
		Name string

		// Brokers 要侦听的 kafka broker 集合
		Brokers []string

		// GroupID kafka 分组 Id
		GroupID string

		// Topics 要侦听的 topic 集合，针对每个 topic 会采用单独的 goroutine 进行侦听
		Topics []string

		// SASL 对于有验证要的 kafka，可设置 SASL
		SASL SASL
	}

	// SASL kafka SASL
	SASL struct {
		// Enabled 是否启用 SASL
		Enabled  bool
		User     string
		Password string
	}

	// RPC 远程调用
	// 对于 RESTFul 请求，目前发送 POST 请求，参数会以 application/json 格式传递
	RPC struct {
		// Url 回调地址
		URL string
		QoS QoS
	}

	// QoS 服务质量保证
	QoS struct {
		// 请求超时时间
		Timeout uint

		// 重试次数
		Retry uint
	}

	// LogConfig 日志配置
	LogConfig struct {
	}
)

// DecodeViaFile 将 toml 配置文件内容反序列化成一个对象
func DecodeViaFile(filename string) (*Config, error) {
	conf := Config{}
	if _, err := toml.DecodeFile(filename, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
