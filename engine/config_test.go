package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	assert := assert.New(t)

	config, err := LoadConfig("./test")
	assert.NoError(err)
	assert.Equal("http://localhost:3333", config.BaseURL)
	assert.Equal("test title", config.Title)
	assert.Equal("test description", config.Description)
	assert.Equal("3333", config.Port)
	assert.Equal("content", config.ContentDir)
	assert.Equal("dist", config.OutputDir)
}
