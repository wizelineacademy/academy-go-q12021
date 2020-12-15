// Test config imports
package config

import (
	"fmt"
	"testing"
)

// Test read config
func TestReadConfig(t *testing.T) {
	err := ReadConfig("config.yml")
	if err != nil {
		t.Errorf("config not valid: %v", err)
	}
}

// Test read config failing
func TestReadConfigFail(t *testing.T) {
	err := ReadConfig("config/config.yml")
	if err != nil {
		fmt.Println("config invalid")
	} else {
		t.Log("config is valid")
		t.Fail()
	}
}
