package config

// Config 根结构体
type Config struct {
	Host Host `yaml:"host"`
}

// Host host节点
type Host struct {
	Port string `yaml:"port"`
}
