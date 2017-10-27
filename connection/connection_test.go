/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

/*
import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernestio/ernest-go-sdk/config"
	"github.com/stretchr/testify/suite"
)

// ConnectionTestSuite : Test suite for connection
type ConnectionTestSuite struct {
	url string
	suite.Suite
	Connection *Conn
}

// SetupTest : sets up test suite
func (suite *ConnectionTestSuite) SetupTest() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/test", testhandler)
	mux.HandleFunc("/auth", handleauth)
	server := httptest.NewServer(mux)

	suite.Connection = New(config.New(server.URL))
}

func (suite *ConnectionTestSuite) TestGet() {
	resp, err := suite.Connection.Get("/api/test")
	body, rerr := ioutil.ReadAll(resp.Body)

	suite.Nil(err)
	suite.Nil(rerr)
	suite.Equal(resp.StatusCode, 200)
	suite.Equal(body, []byte(`{"status":"ok"}`))
}

func (suite *ConnectionTestSuite) TestPost() {
	data := []byte(`{"id":"test"}`)

	resp, err := suite.Connection.Post("/api/test", "application/json", data)
	body, rerr := ioutil.ReadAll(resp.Body)

	suite.Nil(err)
	suite.Nil(rerr)
	suite.Equal(resp.StatusCode, 201)
	suite.Equal(body, data)
}

func (suite *ConnectionTestSuite) TestPut() {
	data := []byte(`{"id":"test"}`)

	resp, err := suite.Connection.Put("/api/test", "application/json", data)
	body, rerr := ioutil.ReadAll(resp.Body)

	suite.Nil(err)
	suite.Nil(rerr)
	suite.Equal(resp.StatusCode, 202)
	suite.Equal(body, data)
}

func (suite *ConnectionTestSuite) TestDelete() {
	resp, err := suite.Connection.Delete("/api/test")
	body, rerr := ioutil.ReadAll(resp.Body)

	suite.Nil(err)
	suite.Nil(rerr)
	suite.Equal(resp.StatusCode, 200)
	suite.Equal(body, []byte(`{"status":"ok"}`))
}

func (suite *ConnectionTestSuite) TestAuthentication() {
	suite.Connection.config.Username = "user"
	suite.Connection.config.Password = "pass"

	err := suite.Connection.Authenticate()
	suite.Nil(err)
	suite.Equal(suite.Connection.config.Token, "test-token")

	suite.Connection.config.Token = ""

	suite.Connection.config.Username = "baduser"
	suite.Connection.config.Password = "badpass"

	err = suite.Connection.Authenticate()
	suite.NotNil(err)
	suite.NotEqual(suite.Connection.config.Token, "test-token")
}

// TestConnectionTestSuite : Test suite for connection
func TestConnectionTestSuite(t *testing.T) {
	suite.Run(t, new(ConnectionTestSuite))
}
*/
