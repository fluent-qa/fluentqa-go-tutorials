package configuration

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestViperDefaultValue(t *testing.T) {
	ViperSetValue()
	ViperSetDefaultValue()
	ViperGetValues()
}

func TestViperEnvConfig(t *testing.T) {
	ViperEnvConfig()
}

func TestWriteConfigFile(t *testing.T) {
	WriteConfigFile("json")

}

func TestReadConfigFile(t *testing.T) {
	ReadConfigFile("json")
}

type testConfig struct {
	BaseConfig `mapstructure:"core"`
	Name       string
}

var testConfigInstance = testConfig{
	BaseConfig: BaseConfig{
		DB: DBConfig{
			Driver: "sqlite",
			DSN:    "./test.db",
		},
		HTTP: HTTPConfig{
			Address: ":8080",
		},
		LogLevel: "info",
	},
	Name: "Test Config",
}

func TestViperConfigToStruct(t *testing.T) {
	viper.Set("config", testConfigInstance)
	fmt.Println(viper.Get("config").(testConfig).DB.Driver)
	//ReadConfigFile, then unmarshal to Struct
}
