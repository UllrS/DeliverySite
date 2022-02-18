package config

import (
	"encoding/json"
	"fmt"
	"knocker/pkg/tools"
	"os"
)

var Config Configs

type Configs struct {
	Server struct {
		Port int    `json:"port"`
		Ip   string `json:"ip"`
	} `json:"server"`
	Database struct {
		System  string `json:"system"`
		Login   string `json:"login"`
		Pwd     string `json:"pwd"`
		Connect string `json:"connect"`
		Addr    string `json:"addr"`
		Port    int    `json:"port"`
		Name    string `json:"name"`
	} `json:"database"`
}

func (c *Configs) GetSqlSettings() (string, string) {
	return c.Database.System, fmt.Sprintf("%s:%s@%s(%s:%d)/%s", c.Database.Login, c.Database.Pwd, c.Database.Connect, c.Database.Addr, c.Database.Port, c.Database.Name)
}
func (c *Configs) GetAddress() string {
	return fmt.Sprintf("%s:%d", c.Server.Ip, c.Server.Port)
}
func (c *Configs) GetServerSettings() string {
	return fmt.Sprintf("")
}
func InitConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		tools.Logger.Fatal(err.Error())
	}
	defer file.Close()
	Config = Configs{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		tools.Logger.Fatal(err.Error())
	}
}
