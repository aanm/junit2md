// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/xml"
)

type Testsuite struct {
	XMLName   xml.Name   `xml:"testsuite"`
	Name      string     `xml:"name,attr"`
	Tests     int        `xml:"tests,attr"`
	Failures  int        `xml:"failures,attr"`
	Errors    int        `xml:"errors,attr"`
	ID        int        `xml:"id,attr"`
	Hostname  string     `xml:"hostname,attr"`
	Time      float64    `xml:"time,attr"`
	Timestamp string     `xml:"timestamp,attr"`
	Testcases []Testcase `xml:"testcase"`
	Filename  string     `xml:"-"`
}

type TestSuiteGinkgo struct {
	Testsuite
}

type TestSuiteGolang struct {
	XMLName  xml.Name    `xml:"testsuites"`
	Tests    int         `xml:"tests,attr"`
	Failures int         `xml:"failures,attr"`
	Skipped  int         `xml:"skipped,attr"`
	Suites   []Testsuite `xml:"testsuite"`
	Filename string      `xml:"-"`
}

type Testcase struct {
	XMLName   xml.Name `xml:"testcase"`
	Name      string   `xml:"name,attr"`
	Classname string   `xml:"classname,attr"`
	Time      float64  `xml:"time,attr"`
	Failure   *Failure `xml:"failure,omitempty"`
	Skipped   *Skipped `xml:"skipped,omitempty"`
	Error     *Error   `xml:"error,omitempty"`
	Filename  string   `xml:"-"`
}

type Skipped struct{}

type Failure struct {
	Message string `xml:"message,attr"`
	Text    string `xml:",chardata"`
}

type Error struct {
	Message string `xml:"message,attr"`
	Text    string `xml:",chardata"`
}
