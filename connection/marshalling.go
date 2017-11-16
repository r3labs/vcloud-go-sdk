/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
)

// ReadXML : reads a json response into an interface
func ReadXML(body io.ReadCloser, x interface{}) error {
	defer body.Close()
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if DEBUG {
		fmt.Println(string(data))
	}

	return xml.Unmarshal(data, x)
}
