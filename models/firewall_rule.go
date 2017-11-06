/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import "encoding/xml"

type FirewallRule struct {
	XMLName              xml.Name      `xml:"FirewallRule"`
	ID                   string        `xml:"Id"`
	Enabled              bool          `xml:"IsEnabled"`
	MatchOnTranslate     bool          `xml:"MatchOnTranslate"`
	Policy               string        `xml:"Policy"`
	Protocols            []interface{} `xml:"Protocols"`
	Port                 string        `xml:"Port"`
	SourcePort           string        `xml:"SourcePort"`
	SourceIp             string        `xml:"SourceIP"`
	DestinationIp        string        `xml:"DestinationIP"`
	SourcePortRange      string        `xml:"SourcePortRange"`
	DestinationPortRange string        `xml:"DestinationPortRange"`
	EnableLogging        bool          `xml:"EnableLogging"`
}

/*
<FirewallService>
                <IsEnabled>true</IsEnabled>
                <DefaultAction>drop</DefaultAction>
                <LogDefaultAction>false</LogDefaultAction>
                <FirewallRule>
                    <Id>1</Id>
                    <IsEnabled>true</IsEnabled>
                    <MatchOnTranslate>false</MatchOnTranslate>
                    <Policy>allow</Policy>
                    <Protocols>
                        <Any>true</Any>
                    </Protocols>
                    <Port>-1</Port>
                    <DestinationPortRange>Any</DestinationPortRange>
                    <DestinationIp>internal</DestinationIp>
                    <SourcePort>-1</SourcePort>
                    <SourcePortRange>Any</SourcePortRange>
                    <SourceIp>internal</SourceIp>
                    <EnableLogging>false</EnableLogging>
                </FirewallRule>
            </FirewallService>
*/
