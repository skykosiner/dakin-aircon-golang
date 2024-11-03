package config

import "testing"

func TestConfig(t *testing.T) {
	config, err := GetConfig()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if config == (Config{}) {
		t.Log("Config is empty.")
		t.Fail()
	}
}
