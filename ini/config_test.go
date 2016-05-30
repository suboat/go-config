package goconfig

import "testing"

func TestLoadConfigFile(t *testing.T) {
	if err := LoadConfigFile("../testdata/conf.ini"); err != nil {
		t.Fatal(err)
	}
	println(ConfigString("Demo", "key1", ""))
}
