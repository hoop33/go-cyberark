package cyberark

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientHostIsRequired(t *testing.T) {
	_, err := NewClient()
	assert.NotNil(t, err)
	assert.Equal(t, "host is required", err.Error())
}

func TestNewClientAddsProtocolWhenNotSpecified(t *testing.T) {
	c, err := NewClient(SetHost("foo"))
	assert.Nil(t, err)
	assert.Equal(t, "https://foo/", c.host)
}

func TestNewClientAllowsProtocol(t *testing.T) {
	c, err := NewClient(SetHost("http://foo"))
	assert.Nil(t, err)
	assert.Equal(t, "http://foo/", c.host)
}

func TestBuildURLBuildsCorrectly(t *testing.T) {
	c, err := NewClient(SetHost("foo"))
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("https://foo/%smy/path", basePath), c.buildURL("my/path"))
}
