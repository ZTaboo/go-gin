package config

// Config 根结构体
type Config struct {
	Host Host `yaml:"host"`
	Jwt  Jwt  `yaml:"jwt"`
}

// Host host节点
type Host struct {
	Port string `yaml:"port"`
}

type Jwt struct {
	SigningKey string `yaml:"signingKey"`
	Issuer     string `yaml:"issuer"`
}
