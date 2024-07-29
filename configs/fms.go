package configs

import (
	"encoding/json"
)

type FMSConfig struct {
	Product Product `json:"product"`
	MQTT    MQTT    `json:"mqtt"`
}

type Product struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Display     string `json:"display"`
	Description string `json:"description"`
}

type MQTT struct {
	Address []string `json:"address"`
}

func (c *FMSConfig) String() string {
	s, _ := json.Marshal(c)
	return string(s)
}
