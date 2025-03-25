package config

import "os"

type ConfigStruct struct {
	InfluxDBURL  string
	InfluxOrg    string
	InfluxBucket string
	InfluxToken  string
}

var Config = LoadConfig()

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func LoadConfig() ConfigStruct {
	return ConfigStruct{
		//InfluxDBURL:  getEnv("INFLUXDB_URL", "http://influxdb2.influxdb.svc.cluster.local:80"),
		InfluxDBURL:  getEnv("INFLUXDB_URL", "http://192.168.3.76:31367"),
		InfluxOrg:    getEnv("INFLUXDB_ORG", "influxdata"),
		InfluxBucket: getEnv("INFLUXDB_BUCKET", "huhubucket"),
		InfluxToken:  getEnv("INFLUXDB_TOKEN", "h73A8DUGPWoaO2p0JdpuTDXUQx2nBiVQ0ba6k6X_5tFC9XMpiXAhZp1XI_ZMEuNi1wPJrAjY2sVkzUx4fFDZMA=="),
	}
}
