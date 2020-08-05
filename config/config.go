package config

type Config struct {
	Id      int
	Sinks   []DeviceConfig
	Sources []DeviceConfig
}

type DeviceConfig struct {
	Id  int
	URI string
}
