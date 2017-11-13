/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package queries

import (
	"fmt"

	"github.com/r3labs/vcloud-go-sdk/connection"
	"github.com/r3labs/vcloud-go-sdk/models"
)

// RecordsFilter : query records with filter
func (q *Queries) RecordsFilter(qt, filter, page string) (*models.QueryResultRecords, error) {
	var m models.QueryResultRecords

	path := fmt.Sprintf(apiroute+"?type=%s&page=%s&format=records&filter=%s", qt, page, filter)

	resp, err := q.Conn.Get(path)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &m, connection.ReadXML(resp.Body, &m)
}
