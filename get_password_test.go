package cyberark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func DoShouldFailWhenNoClient(t *testing.T) {
	s := &GetPasswordService{}
	_, err := s.Do()
	assert.NotNil(t, err)
	assert.Equal(t, "Client is required", err.Error())
}

func DoShouldFailWhenNoAppID(t *testing.T) {
	c, err := NewClient(SetHost("foo"))
	assert.Nil(t, err)

	_, err = c.GetPassword().Do()
	assert.NotNil(t, err)
	assert.Equal(t, "AppID is required", err.Error())
}

func TestAppIDShouldGetSet(t *testing.T) {
	s := newGetPasswordService(nil).AppID("app")
	assert.Equal(t, "app", s.appID)
}

func TestAddressShouldGetSet(t *testing.T) {
	s := newGetPasswordService(nil).Address("address")
	assert.Equal(t, "address", s.address)
}

func TestDatabaseShouldGetSet(t *testing.T) {
	s := newGetPasswordService(nil).Database("database")
	assert.Equal(t, "database", s.database)
}

func TestFolderShouldGetSet(t *testing.T) {
	s := newGetPasswordService(nil).Folder("folder")
	assert.Equal(t, "folder", s.folder)
}

func TestObjectShouldGetSet(t *testing.T) {
	s := newGetPasswordService(nil).Object("object")
	assert.Equal(t, "object", s.object)
}

func TestPolicyIDShouldGetSet(t *testing.T) {
	s := newGetPasswordService(nil).PolicyID("policy")
	assert.Equal(t, "policy", s.policyID)
}

func TestQueryShouldGetSet(t *testing.T) {
	s := newGetPasswordService(nil).Query("query")
	assert.Equal(t, "query", s.query)
}

func TestQueryFormatShouldGetSet(t *testing.T) {
	s := newGetPasswordService(nil).QueryFormat("query format")
	assert.Equal(t, "query format", s.queryFormat)
}

func TestReasonShouldGetSet(t *testing.T) {
	s := newGetPasswordService(nil).Reason("reason")
	assert.Equal(t, "reason", s.reason)
}

func TestSafeShouldGetSet(t *testing.T) {
	s := newGetPasswordService(nil).Safe("safe")
	assert.Equal(t, "safe", s.safe)
}

func TestTimeoutShouldGetSet(t *testing.T) {
	s := newGetPasswordService(nil).Timeout(1000)
	assert.Equal(t, 1000, s.timeout)
}

func TestUserNameShouldGetSet(t *testing.T) {
	s := newGetPasswordService(nil).UserName("user")
	assert.Equal(t, "user", s.userName)
}
