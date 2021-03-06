package chainlib

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/go-ini/ini"
)

//GetConfig ...
func GetConfig() (*ini.File, error) {
	var cfg *ini.File
	var err error
	if runtime.GOOS == "linux" {
		fmt.Println("Unix/Linux type OS detected")
		cfg, err = ini.Load(os.Getenv("HOME") + "/.local/share/datxio/noddatx/config/config.ini")
	} else if runtime.GOOS == "darwin" {
		fmt.Println("Mac OS detected")
		cfg, err = ini.Load(os.Getenv("HOME") + "/Library/Application Support/datxio/noddatx/config/config.ini")
	} else {
		return nil, fmt.Errorf("%s detected,not support", runtime.GOOS)
	}

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

//GetCfgProducerName ...
func GetCfgProducerName() string {
	var result string
	cfg, err := GetConfig()
	if err != nil {
		fmt.Printf("Get config err:%v\n", err)
		return result
	}

	result = cfg.Section("").Key("producer-name").String()
	return result
}

//GetCfgProducerKey Get node key
func GetCfgProducerKey() []string {
	cfg, err := GetConfig()
	if err != nil {
		fmt.Printf("Get config err:%v\n", err)
		return nil
	}

	sig := cfg.Section("").Key("signature-provider").String()
	result := strings.Split(sig, "=KEY:")
	return result
}
