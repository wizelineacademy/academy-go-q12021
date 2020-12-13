/**
Test Config
*/
package config

import (
	"fmt"
	"strings"
	"testing"
)

// Test read config
func TestConfig(t *testing.T) {
	err := ReadConfig("config")
	if err != nil {
		t.Errorf("config not valid: %v", err)
	}
}

// Test read config fail
func TestConfigFail(t *testing.T) {
	err := ReadConfig("configFailx")
	if err != nil {
		fmt.Println("config invalid")
	} else {
		t.Log("config is valid")
		t.Fail()
	}
}

// Test get server address from config
func TestConfig_GetServerAddr(t *testing.T) {
	err := ReadConfig("config")
	if err != nil {
		t.Log("server addrs fail")
		t.Fail()
	}
	Config := &C

	want := "localhost:8080"
	serverAdds := Config.GetServerAddr()
	if serverAdds != "" {
		if !strings.Contains(serverAdds, want) {
			t.Errorf("Unexpected addrs %v  want %v", serverAdds, want)
		}
	} else {
		t.Log("server address null")
		t.Fail()
	}
}
