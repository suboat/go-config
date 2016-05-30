package goconfig

import config "github.com/Unknwon/goconfig"

var _config *config.ConfigFile

func LoadConfigFile(fileName string, moreFiles ...string) (err error) {
	_config, err = config.LoadConfigFile(fileName, moreFiles...)
	return
}

func AppendConfigFiles(files ...string) (err error) {
	if _config == nil {
		return LoadConfigFile(files[0])
	}
	return _config.AppendFiles(files...)
}

func ConfigString(section, key string, df string) string {
	s, err := _config.GetValue(section, key)
	if err != nil {
		return df
	}
	return s
}

func ConfigBool(section, key string, df bool) bool {
	b, err := _config.Bool(section, key)
	if err != nil {
		return df
	}
	return b
}

func ConfigInt(section, key string, df int) int {
	i, err := _config.Int(section, key)
	if err != nil {
		return df
	}
	return i
}
