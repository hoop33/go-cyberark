package cyberark

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewClientHostIsRequired(t *testing.T) {
	_, err := NewClient()
	assert.NotNil(t, err)
	assert.Equal(t, "host is required", err.Error())
}

func TestSetHostFailsOnBlank(t *testing.T) {
	_, err := NewClient(SetHost(""))
	assert.NotNil(t, err)
	assert.Equal(t, "host cannot be empty", err.Error())
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

func TestSetSkipCertVerificationSetsToTrue(t *testing.T) {
	c, err := NewClient(
		SetHost("foo"),
		SetSkipCertVerification(true),
	)
	assert.Nil(t, err)
	assert.True(t, c.skipCertVerification)
}

func TestSetSkipCertVerificationSetsToFalse(t *testing.T) {
	c, err := NewClient(
		SetHost("foo"),
		SetSkipCertVerification(false),
	)
	assert.Nil(t, err)
	assert.False(t, c.skipCertVerification)
}

func TestDurationDefaultsTo30(t *testing.T) {
	c, err := NewClient(
		SetHost("foo"),
	)
	assert.Nil(t, err)
	assert.Equal(t, time.Duration(30), c.timeout)
}

func TestSetTimeoutSetsTimeout(t *testing.T) {
	c, err := NewClient(
		SetHost("foo"),
		SetTimeout(15),
	)
	assert.Nil(t, err)
	assert.Equal(t, time.Duration(15), c.timeout)
}

func TestGetPasswordReturnsNonNil(t *testing.T) {
	c, err := NewClient(
		SetHost("foo"),
	)
	assert.Nil(t, err)
	assert.NotNil(t, c.GetPassword())
}
