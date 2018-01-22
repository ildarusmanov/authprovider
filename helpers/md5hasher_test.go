package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMD5Hash(t *testing.T) {
	hashMap := map[string]string{
		"test1": "5a105e8b9d40e1329780d62ea2265d8a",
		"test2": "ad0234829205b9033196ba818f7a872b",
		"test3": "8ad8757baa8564dc136c1e07507f4a98",
	}

	for src, hash := range hashMap {
		assert.Equal(t, GetMD5Hash(src), hash)
	}
}
