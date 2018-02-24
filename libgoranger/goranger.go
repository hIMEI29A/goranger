// Copyright 2018 hIMEI
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package libgoranger

import (
	"bufio"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// API endpoints
const (
	CITY    = "https://suip.biz/?act=iploc"
	COUNTRY = "https://suip.biz/?act=iploc"
	ISP     = "https://suip.biz/?act=ipintpr"
)

// PRE is a <pre> HTML tag
const PRE = "//pre"

// CountryCodes is a two-letter country codes
var CountryCodes = []string{
	"AD", "AE", "AF", "AG", "AI", "AL", "AM", "AO", "AQ", "AR", "AS",
	"AT", "AU", "AW", "AX", "AZ", "BA", "BB", "BD", "BE", "BF", "BG",
	"BH", "BI", "BJ", "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT",
	"BV", "BW", "BY", "BZ", "CA", "CC", "CD", "CF", "CG", "CH", "CI",
	"CK", "CL", "CM", "CN", "CO", "CR", "CU", "CV", "CW", "CX", "CY",
	"CZ", "DE", "DJ", "DK", "DM", "DO", "DZ", "EC", "EE", "EG", "EH",
	"ER", "ES", "ET", "FI", "FJ", "FK", "FM", "FO", "FR", "GA", "GB",
	"GD", "GE", "GF", "GG", "GH", "GI", "GL", "GM", "GN", "GP", "GQ",
	"GR", "GS", "GT", "GU", "GW", "GY", "HK", "HM", "HN", "HR", "HT",
	"HU", "ID", "IE", "IL", "IM", "IN", "IO", "IQ", "IR", "IS", "IT",
	"JE", "JM", "JO", "JP", "KE", "KG", "KH", "KI", "KM", "KN", "KP",
	"KR", "KW", "KY", "KZ", "LA", "LB", "LC", "LI", "LK", "LR", "LS",
	"LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MF", "MG", "MH",
	"MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU",
	"MV", "MW", "MX", "MY", "MZ", "NA", "NC", "NE", "NF", "NG", "NI",
	"NL", "NO", "NP", "NR", "NU", "NZ", "OM", "PA", "PE", "PF", "PG",
	"PH", "PK", "PL", "PM", "PN", "PR", "PS", "PT", "PW", "PY", "QA",
	"RE", "RO", "RS", "RU", "RW", "SA", "SB", "SC", "SD", "SE", "SG",
	"SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SR", "SS", "ST",
	"SV", "SX", "SY", "SZ", "TC", "TD", "TF", "TG", "TH", "TJ", "TK",
	"TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW", "TZ", "UA", "UG",
	"UM", "US", "UY", "UZ", "VA", "VC", "VE", "VG", "VI", "VN", "VU",
	"WF", "WS", "YE", "YT", "ZA", "ZM", "ZW",
}

// ReqType is a possible types of request
var ReqType = []string{
	"city",
	"country",
	"isp",
}

// ApiEndPoints is a slice of string constants representing API endpoints
var ApiEndPoints = []string{
	CITY,
	COUNTRY,
	ISP,
}

// GetTag gets inner value of html tag
func getTag(node *html.Node, tagexp string) string {
	return htmlquery.InnerText(htmlquery.FindOne(node, tagexp))
}

// ValidateCountry checks if given country code is correct
func ValidateCountry(ccode string) bool {
	valid := false

	for i := range CountryCodes {
		if ccode == CountryCodes[i] {
			valid = true
			break
		}
	}

	return valid
}

// Goranger is a main package's data type. EndPoint field may be one of the endpoint constants
type Goranger struct {
	EndPoint []string
}

// NewGoranger creates Goranger's instance
func NewGoranger() *Goranger {
	goranger := &Goranger{}
	goranger.EndPoint = ApiEndPoints

	return goranger
}

// SetEndPoint sets API endpoint by given request type
func (g *Goranger) SetEndPoint(reqType string) (string, error) {
	var point string

	switch {
	case reqType == ReqType[0]:
		point = g.EndPoint[0]

	case reqType == ReqType[1]:
		point = g.EndPoint[1]

	case reqType == ReqType[2]:
		point = g.EndPoint[2]

	default:
		errString := "Wrong request's type"
		err := errors.New(errString)

		return "", err
	}

	return point, nil
}

/*
// NewGoranger creates Goranger's instance
func NewGoranger(reqType string) (*Goranger, error) {
	var ttype string
	goranger := &Goranger{}

	switch {
	case reqType == ReqType[0]:
		ttype = CITY

	case reqType == ReqType[1]:
		ttype = COUNTRY

	case reqType == ReqType[2]:
		ttype = ISP

	default:
		errString := "Wrong request's type"
		err := errors.New(errString)

		return nil, err
	}

	goranger.EndPoint = ttype

	return goranger, nil
}
*/

// GetData makes POST request to site and returns response's body as *html.Node
func (g *Goranger) getData(reqType, req string) (*html.Node, error) {
	reqForm := url.Values{
		"url":    {req},
		"action": {"Submit"},
	}

	point, err := g.SetEndPoint(reqType)
	if err != nil {
		return nil, err
	}

	response, err := http.PostForm(point, reqForm)
	if err != nil {
		errString := "Network connection's error"
		newErr := errors.New(errString)
		return nil, newErr
	}

	defer response.Body.Close()

	body := bufio.NewReader(response.Body)
	node, err := htmlquery.Parse(body)
	if err != nil {
		errString := "HTML parsing error"
		newErr := errors.New(errString)
		return nil, newErr
	}

	return node, nil
}

// GetRange is a main package's method. It uses getData() to make request, parses
// <pre> tag's content and returns it as []string
func (g *Goranger) GetRange(reqType, req string) ([]string, error) {
	node, err := g.getData(reqType, req)
	if err != nil {
		return nil, err
	}

	ipRange := getTag(node, PRE)
	ipSplitted := strings.Split(ipRange, "\n")

	return ipSplitted, nil
}
