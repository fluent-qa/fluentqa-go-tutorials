package orm

import "github.com/qdriven/qfluent-go/config"

type Oracle struct {
	config.GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *Oracle) Dsn() string {
	return "oracle://" + m.Username + ":" + m.Password + "@" + m.Path + ":" + m.Port + "/" + m.Dbname + "?" + m.Config

}

func (m *Oracle) GetLogMode() string {
	return m.LogMode
}
