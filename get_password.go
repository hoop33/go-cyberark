package cyberark

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

const getPasswordPath = "Accounts"

// GetPasswordService gets passwords
type GetPasswordService struct {
	client      *Client
	appID       string
	address     string
	database    string
	folder      string
	object      string
	policyID    string
	query       string
	queryFormat string
	reason      string
	safe        string
	timeout     int
	userName    string
}

// GetPasswordResult returns the result from getting passwords
type GetPasswordResult struct {
	StatusCode int
	ErrorCode  string
	ErrorMsg   string
	Content    string
	UserName   string
	Address    string
	Database   string
	PolicyID   string
	Properties map[string]string
}

func newGetPasswordService(client *Client) *GetPasswordService {
	return &GetPasswordService{
		client:  client,
		timeout: 30,
	}
}

// AppID sets the app ID
func (s *GetPasswordService) AppID(appID string) *GetPasswordService {
	s.appID = appID
	return s
}

// Address sets the address
func (s *GetPasswordService) Address(address string) *GetPasswordService {
	s.address = address
	return s
}

// Database sets the database
func (s *GetPasswordService) Database(database string) *GetPasswordService {
	s.database = database
	return s
}

// Folder sets the folder
func (s *GetPasswordService) Folder(folder string) *GetPasswordService {
	s.folder = folder
	return s
}

// Object sets the object
func (s *GetPasswordService) Object(object string) *GetPasswordService {
	s.object = object
	return s
}

// PolicyID sets the policy ID
func (s *GetPasswordService) PolicyID(policyID string) *GetPasswordService {
	s.policyID = policyID
	return s
}

// Query sets the query
func (s *GetPasswordService) Query(query string) *GetPasswordService {
	s.query = query
	return s
}

// QueryFormat sets the query format
func (s *GetPasswordService) QueryFormat(queryFormat string) *GetPasswordService {
	s.queryFormat = queryFormat
	return s
}

// Reason sets the reason
func (s *GetPasswordService) Reason(reason string) *GetPasswordService {
	s.reason = reason
	return s
}

// Safe sets the safe
func (s *GetPasswordService) Safe(safe string) *GetPasswordService {
	s.safe = safe
	return s
}

// Timeout sets the connection timeout
func (s *GetPasswordService) Timeout(timeout int) *GetPasswordService {
	s.timeout = timeout
	return s
}

// UserName sets the user name
func (s *GetPasswordService) UserName(userName string) *GetPasswordService {
	s.userName = userName
	return s
}

// Do runs the service
func (s *GetPasswordService) Do() (*GetPasswordResult, error) {
	if s.client == nil {
		return nil, errors.New("Client is required")
	}
	if s.appID == "" {
		return nil, errors.New("AppID is required")
	}

	path, params := s.buildURL()

	resp, err := s.client.PerformRequest("GET", path, params, nil)
	if err != nil {
		return nil, err
	}
	ret := new(GetPasswordResult)
	if err := json.Unmarshal(resp.Body, ret); err != nil {
		return nil, err
	}
	ret.StatusCode = resp.StatusCode
	return ret, nil
}

func (s *GetPasswordService) buildURL() (string, url.Values) {
	params := url.Values{}

	setParam(&params, "appId", s.appID)
	setParam(&params, "address", s.address)
	setParam(&params, "database", s.database)
	setParam(&params, "folder", s.folder)
	setParam(&params, "object", s.object)
	setParam(&params, "policyID", s.policyID)
	setParam(&params, "query", s.query)
	setParam(&params, "queryFormat", s.queryFormat)
	setParam(&params, "reason", s.reason)
	setParam(&params, "safe", s.safe)
	setParam(&params, "timeout", strconv.Itoa(s.timeout))
	setParam(&params, "userName", s.userName)

	return getPasswordPath, params
}

func setParam(params *url.Values, key string, value string) {
	if value != "" {
		params.Set(key, value)
	}
}
