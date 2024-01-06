package game

import (
	"github.com/i582/cfmt/cmd/cfmt"
	"gopkg.in/ini.v1"
	"os"
)

type GameConfig struct {
	Config *ini.File
}

func ReadGameConfig(filename string) *GameConfig {
	config := GameConfig{}

	cfg, err := ini.Load("client.ini")
	if err != nil {
		cfmt.Printf("{{Fail to read file: %v.}}::red|bold\n", err)
		os.Exit(1)
	}
	config.Config = cfg

	return &config
}

func (self *GameConfig) GetValue(section string, key string) *ini.Key {
	sect, err := self.Config.GetSection(section)
	if err != nil || sect == nil {
		cfmt.Printf("{{Fail get section %s.}}::red|bold\n", section)
		return nil
	}

	if !sect.HasKey(key) {
		cfmt.Printf("{{Fail get key %s.%s}}::red|bold\n", section, key)
		return nil
	}

	return sect.Key(key)
}
