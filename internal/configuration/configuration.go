package config

type Configuration struct {
	Server  ServerConfiguration `yaml:"server"`
	Lorawan LorwanConfiguration `yaml:"lorawan"`
}

type ServerConfiguration struct {
	Port string
}

type LorwanConfiguration struct {
	Baseurl  string `yaml:"baseurl"`
	Endpoint string `yaml:"endpoint"`
}
