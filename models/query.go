/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"fmt"
	"net/url"
)

// Query ...
type Query struct {
	Type      string
	Format    string
	Filter    string
	FilterArg string
	Results   string
}

// QueryURL : generates the queries url
func (q *Query) QueryURL() url.Values {
	query := url.Values{}

	query.Add("type", q.Type)
	query.Add("format", q.Format)
	query.Add("filter", fmt.Sprintf("%s==%s", q.Filter, q.FilterArg))

	return query
}
