# Viper

> Viper is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats. It supports:
    setting defaults
    reading from JSON, TOML, YAML, HCL, envfile and Java properties config files
    live watching and re-reading of config files (optional)
    reading from environment variables
    reading from remote config systems (etcd or Consul), and watching changes
    reading from command line flags
    reading from buffer
    setting explicit values


## Viper: Set Configurations

```go
func ViperSetValue() {
    viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
    viper.SetDefault("ContentDir", "content")
    viper.SetDefault("LayoutDir", "layouts")
    viper.Set("Verbose", true)
    viper.Set("LogFile", "LogFile")
    viper.Set("host.port", 5899) // set subset
}
```

## Viper: Get Configurations

```go
func ViperDefaultValue() {
// REQUIRED if the config file does not have the extension in the name
    result := viper.Get("Taxonomies")
    fmt.Println(result)
    fmt.Println(reflect.TypeOf(result))
    s := result.(map[string]string)["tag"]
    fmt.Printf("tag: %s", s)
    fmt.Println(viper.GetBool("Verbose"))
    fmt.Println(viper.Get("host"))
    fmt.Println(viper.Get("host.port"))
    fmt.Println(viper.Get("Taxonomies.category"))
}
```
