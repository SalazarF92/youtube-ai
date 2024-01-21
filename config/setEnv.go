package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func SetEnv(envVarName string, value string) error {
	os.Setenv(envVarName, value)
	return nil
}

func SetEnvFromStruct(envVarName string, data interface{}) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	print(string(jsonBytes))

	os.Setenv(envVarName, string(jsonBytes))
	return nil
}
