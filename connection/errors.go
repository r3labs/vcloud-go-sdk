/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import (
	"encoding/xml"
	"errors"
	"net/http"
)

// Error ...
type Error struct {
	XMLName        xml.Name `xml:"Error"`
	MinorErrorCode string   `xml:"minorErrorCode,attr"`
	MajorErrorCode string   `xml:"majorErrorCode,attr"`
	Message        string   `xml:"message,attr"`
}

func status(s int) error {
	switch s {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted:
		return nil
	case http.StatusBadRequest:
		return errors.New("ernest responded with code 400 : 'bad request'")
	case http.StatusUnauthorized:
		return errors.New("you are not authorized to perform this action, please log in")
	case http.StatusForbidden:
		return errors.New("you are not autorized to perform this action with your level of permissions")
	case http.StatusNotFound:
		return errors.New("the resource does not exist")
	case http.StatusInternalServerError:
		return errors.New("ernest responded with code 500 : 'internal server error'")
	default:
		return errors.New("ernest unknown error")
	}
}

func errored(resp *http.Response, fallback error) (*http.Response, error) {
	var e Error

	err := ReadXML(resp.Body, &e)
	if err != nil {
		return resp, err
	}

	if e.Message == "" {
		return resp, fallback
	}

	return resp, errors.New(e.Message)
}
