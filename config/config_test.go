package config

import (
	"strings"
	"testing"
)

func TestConfig_GetServerAddr(t *testing.T) {
	ReadConfig()
	Config := &C

	want:="localhost:8080"
	serverAdds :=Config.GetServerAddr()
	if serverAdds !=""{
		if !strings.Contains(serverAdds, want){
			t.Errorf("Unexpected addrs %v  want %v", serverAdds,want)
		}
	}else {
		t.Log(serverAdds)
	}

}