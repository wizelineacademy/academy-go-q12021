// Test config imports
package config

import (
	"fmt"
	"strings"
	"testing"
)

// Test read config
func TestReadConfig(t *testing.T) {
	err := ReadConfig("config")
	if err != nil {
		t.Errorf("config not valid: %v", err)
	}
}

// Test read config failing
func TestReadConfigFail(t *testing.T) {
	err := ReadConfig("configFailx")
	if err != nil {
		fmt.Println("config invalid")
	} else {
		t.Log("config is valid")
		t.Fail()
	}
}

// Test get server address from config
func TestGetServerAddr(t *testing.T) {
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
