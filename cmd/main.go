// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

const (
	StatusPass    = "Pass"
	StatusSkipped = "Skipped"
	StatusFail    = "Fail"
	StatusError   = "Error"
)

type TestCaseResult struct {
	TestSuite string
	TestCase  string
	Status    string
	Msg       string
	Time      float64
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the path to the input file.")
		return
	}

	inputFile := flag.String("i", "", "Input file path")
	outputFile := flag.String("o", "", "Output file path")
	failed := flag.Bool("f", true, "Show failed tests")
	passed := flag.Bool("p", true, "Show passed tests")
	skipped := flag.Bool("s", true, "Show skipped tests")
	errored := flag.Bool("e", true, "Show errored tests")
	verbose := flag.Bool("v", false, "Show stderr/stdout of tests")
	flag.Parse()

	f, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	var testCaseResults []TestCaseResult

	testCaseResults, err = parseGolangFile(f, passed, skipped, failed, errored)
	if err != nil {
		_, err := f.Seek(0, io.SeekStart)
		if err != nil {
			panic(fmt.Errorf("unable to seek to start %w", err))
		}

		testCaseResults, err = parseGinkgoFile(f, passed, skipped, failed, errored)
		if err != nil {
			panic(err)
		}
	}

	if len(testCaseResults) == 0 {
		return
	}

	sort.Slice(testCaseResults, func(i, j int) bool {
		if testCaseResults[i].Status != testCaseResults[j].Status {
			switch testCaseResults[i].Status {
			case StatusError, StatusFail:
				switch testCaseResults[j].Status {
				case StatusError, StatusFail:
					return strings.Compare(testCaseResults[i].TestSuite, testCaseResults[j].TestSuite) < 0
				}
				return true
			}
		}
		return strings.Compare(testCaseResults[i].TestSuite, testCaseResults[j].TestSuite) < 0
	})

	fout := os.Stdout
	if outputFile != nil && len(*outputFile) != 0 {
		fout, err = os.Create(*outputFile)
		if err != nil {
			panic(fmt.Errorf("error creating output file: %s", err))
			return
		}
		defer fout.Close()
	}

	err = writeMarkdown(fout, testCaseResults, verbose)
	if err != nil {
		panic(fmt.Errorf("error creating output to %q: %s", fout.Name(), err))
	}

	if outputFile != nil && len(*outputFile) != 0 {
		fmt.Println("Markdown table saved to", *outputFile)
	}
}

func writeMarkdown(fout io.Writer, testCaseResults []TestCaseResult, verbose *bool) error {
	printOut := func(s string) error {
		_, err := io.WriteString(fout, s)
		if err != nil {
			return fmt.Errorf("unable to write string to output: %w", err)
		}
		return nil
	}

	err := printOut("| Status | Package | Time (seconds) |\n")
	if err != nil {
		return err
	}

	err = printOut("|--------|---------|----------------|\n")
	if err != nil {
		return err
	}

	for _, result := range testCaseResults {
		statusEmoji := emojiResult(result)
		row := fmt.Sprintf("| %-6s | %-10s | %-14.3f |\n", statusEmoji, result.TestSuite, result.Time)

		err = printOut(row)
		if err != nil {
			return err
		}
	}

	// Print more verbose text for the failing tests
	if !*verbose {
		return nil
	}
	for _, result := range testCaseResults {
		if result.Msg == "" {
			continue
		}

		row := fmt.Sprintf("\n### `%s`\n", result.TestSuite)
		err = printOut(row)
		if err != nil {
			return err
		}

		row = fmt.Sprintf("#### Test Case: %s\n", result.TestCase)
		err = printOut(row)
		if err != nil {
			return err
		}

		row = fmt.Sprintf("<details><summary>Click here for error message</summary>\n\n```\n%s\n```\n</details>\n", result.Msg)
		err = printOut(row)
		if err != nil {
			return err
		}
	}
	return printOut("\n")
}

func parseGolangFile(f *os.File, passed *bool, skipped *bool, failed *bool, errored *bool) ([]TestCaseResult, error) {
	var tsGo TestSuiteGolang
	if err := xml.NewDecoder(f).Decode(&tsGo); err != nil {
		return nil, err
	}
	var testCaseResults []TestCaseResult
	for _, suite := range tsGo.Suites {
		if suite.Name == "" {
			continue
		}
		testSuiteStatus := StatusPass
		if len(suite.Testcases) == 0 {
			testSuiteStatus = StatusSkipped
		}
		var failure, failureCaseName string
		for _, testcase := range suite.Testcases {
			testCaseStatus := status(testcase)
			switch testCaseStatus {
			case StatusPass, StatusSkipped:
			default:
				testSuiteStatus = testCaseStatus
			}
			if testcase.Failure != nil {
				failure = testcase.Failure.Text
				failureCaseName = testcase.Name
			}
			if testcase.Error != nil {
				failure = testcase.Error.Text
				failureCaseName = testcase.Name
			}
		}
		testCaseResults = addTestCase(testCaseResults, suite.Name, testSuiteStatus, failureCaseName, failure, suite.Time, passed, skipped, failed, errored)
	}
	return testCaseResults, nil
}

func parseGinkgoFile(f *os.File, passed *bool, skipped *bool, failed *bool, errored *bool) ([]TestCaseResult, error) {
	var tsGinkgo TestSuiteGinkgo

	err := xml.NewDecoder(f).Decode(&tsGinkgo)
	if err != nil {
		return nil, fmt.Errorf("unable to decode ginkgo junit XML: %w", err)
	}

	suite := tsGinkgo.Testsuite
	if suite.Name == "" {
		return nil, nil
	}
	var testCaseResults []TestCaseResult
	for _, testcase := range suite.Testcases {
		var failure string
		testCaseStatus := status(testcase)
		if testcase.Failure != nil {
			failure = testcase.Failure.Text
		}
		testCaseResults = addTestCase(testCaseResults, testcase.Name, testCaseStatus, testcase.Name, failure, testcase.Time, passed, skipped, failed, errored)
	}
	return testCaseResults, nil
}

func emojiResult(result TestCaseResult) string {
	var statusEmoji string
	switch result.Status {
	case StatusPass:
		statusEmoji = ":heavy_check_mark:"
	case StatusSkipped:
		statusEmoji = ":white_check_mark:"
	case StatusFail:
		statusEmoji = ":x:"
	case StatusError:
		statusEmoji = ":warning:"
	}
	return statusEmoji
}

func status(testcase Testcase) string {
	var status string
	switch {
	case testcase.Skipped != nil:
		status = StatusSkipped
	case testcase.Failure != nil:
		status = StatusFail
	case testcase.Error != nil:
		status = StatusError
	default:
		status = StatusPass
	}
	return status
}

func addTestCase(testCaseResults []TestCaseResult, name, status, testCase, failure string, timeElapsed float64, passed, skipped, failed, errored *bool) []TestCaseResult {
	testCaseResult := TestCaseResult{
		TestSuite: name,
		TestCase:  testCase,
		Status:    status,
		Time:      timeElapsed,
		Msg:       failure,
	}
	switch status {
	case StatusPass:
		if *passed {
			testCaseResults = append(testCaseResults, testCaseResult)
		}
	case StatusSkipped:
		if *skipped {
			testCaseResults = append(testCaseResults, testCaseResult)
		}
	case StatusFail:
		if *failed {
			testCaseResults = append(testCaseResults, testCaseResult)
		}
	case StatusError:
		if *errored {
			testCaseResults = append(testCaseResults, testCaseResult)
		}
	}
	return testCaseResults
}
