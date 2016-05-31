package goconfig

import (
	config "github.com/Unknwon/goconfig"
	"strings"
)

var _config *config.ConfigFile

func LoadConfigFile(fileName string, moreFiles ...string) (err error) {
	_config, err = config.LoadConfigFile(fileName, moreFiles...)
	return
}

func AppendConfigFiles(files ...string) (err error) {
	if _config == nil {
		if files == nil || len(files) == 0 {
			return nil
		} else if len(files) == 1 {
			return LoadConfigFile(files[0])
		}
		return LoadConfigFile(files[0], files[1:]...)
	}
	return _config.AppendFiles(files...)
}

func getSectionKey(str string) (string, string) {
	strs := strings.SplitN(str, "::", 2)
	if len(strs) == 0 {
		return "", ""
	} else if len(strs) == 1 {
		return "", strs[0]
	}
	return strs[0], strs[1]
}

func ConfigString(path string, defaultVal ...string) string {
	section, key := getSectionKey(path)
	return _config.MustValue(section, key, defaultVal...)
}

func ConfigBool(path string, defaultVal ...bool) bool {
	section, key := getSectionKey(path)
	return _config.MustBool(section, key, defaultVal...)
}

func ConfigInt(path string, defaultVal ...int) int {
	section, key := getSectionKey(path)
	return _config.MustInt(section, key, defaultVal...)
}

func ConfigInt64(path string, defaultVal ...int64) int64 {
	section, key := getSectionKey(path)
	return _config.MustInt64(section, key, defaultVal...)
}

func ConfigFloat64(path string, defaultVal ...float64) float64 {
	section, key := getSectionKey(path)
	return _config.MustFloat64(section, key, defaultVal...)
}
