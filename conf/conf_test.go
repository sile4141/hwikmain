package conf

import (
	"github.com/hwikpass/hwik-go/core/logger"
	"github.com/magiconair/properties"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"strings"
	"testing"
)

func TestConfigDev(t *testing.T) {

	configPath, _ := filepath.Abs("./conf.properties")
	p, err := properties.LoadFile(configPath, properties.UTF8)
	if err != nil {
		logger.Error(err)
	}

	for _, v := range p.Keys() {
		if startWith(v, "local") {
			key := strings.Replace(v, "local.", "dev.", -1)
			val, _ := p.Get(key)
			assert.NotEmpty(t, val, key)

			key = strings.Replace(v, "local.", "staging.", -1)
			val, _ = p.Get(key)
			assert.NotEmpty(t, val, key)

			key = strings.Replace(v, "local.", "prod.", -1)
			val, _ = p.Get(key)
			assert.NotEmpty(t, val, key)
		}
	}

}

func startWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}
