package configuration

import (
	"bytes"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"reflect"
)

func ViperSetValue() {
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.Set("Verbose", true)
	viper.Set("LogFile", "LogFile")
	viper.Set("host.port", 5899) // set subset
	viper.RegisterAlias("loud", "Verbose")

}

func ViperSetDefaultValue() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")
}

func ViperGetValues() {
	// REQUIRED if the config file does not have the extension in the name
	result := viper.Get("Taxonomies")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
	s := result.(map[string]string)["tag"]
	fmt.Printf("tag: %s", s)
	fmt.Println(viper.GetBool("Verbose"))
	fmt.Println(viper.Get("loud"))
	fmt.Println(viper.Get("host"))
	fmt.Println(viper.Get("host.port"))
	fmt.Println(viper.Get("Taxonomies.category")) //nil
}

func ViperEnvConfig() {
	viper.SetEnvPrefix("spf") // will be uppercased automatically
	viper.BindEnv("id")

	os.Setenv("SPF_ID", "13") // typically done outside of the app

	id := viper.Get("id") // 13
	fmt.Println(id)
}

func ReadConfigFile(inputType string) {
	viper.SetConfigType(inputType)
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	fmt.Println(viper.Get("name"))
}

func WriteConfigFile(outputType string) {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")

	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	viper.Get("name") // this would be "steve"
	fmt.Println(viper.Get("clothing.jacket"))
	viper.SetConfigType(outputType) // or viper.SetConfigType("YAML")

	viper.WriteConfig() // writes current config to predefined path set by 'viper.AddConfigPath()' and 'viper.SetConfigName'
	viper.SafeWriteConfig()
	viper.WriteConfigAs("config." + outputType)
	viper.SafeWriteConfigAs("config-safe." + outputType) // will error since it has already been written
}

func WatchConfigChanges() {
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
}

func ViperConfigurationFileToStruct() {

}

func ViperRemoteConfig() {
	_ = viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.json")
	viper.SetConfigType("json") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	_ = viper.ReadRemoteConfig()
	_ = viper.AddRemoteProvider("etcd3", "http://127.0.0.1:4001", "/config/hugo.json")
	_ = viper.ReadRemoteConfig()
}
