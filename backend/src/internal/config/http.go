package config

import "strconv"

type HTTPConfig struct {
	Host string
	Port int
}

func (http *HTTPConfig) fill(getEnv func(key, defaultValue string) (string, error)) error {
	host, err := getEnv("HTTP_HOST", "127.0.0.1")
	if err != nil {
		return err
	}
	http.Host = host

	portStr, err := getEnv("HTTP_PORT", "8000")
	if err != nil {
		return err
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return err
	}
	http.Port = port

	return nil
}

func (http *HTTPConfig) GetAddr() string {
	return http.Host + ":" + strconv.Itoa(http.Port)
}
