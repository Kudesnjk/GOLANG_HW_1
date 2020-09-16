package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type options struct {
	cFlag    bool
	dFlag    bool
	uFlag    bool
	iFlag    bool
	fFlag    int
	sFlag    int
	inputFN  string
	outputFN string
}

func main() {
	cmdArgs := getCmdArgs()
	option, err := argsToOptions(cmdArgs)
	checkError(err)
	data, err := readFromInput(option)
	checkError(err)
	result := checkUniqString(data, option)
	err = writeIntoOutput(option, result)
	checkError(err)
}

func argsToOptions(cmdArgs []string) (*options, error) {
	o := newOptions()
	for idx, val := range cmdArgs {
		notLast := idx+1 < len(cmdArgs)
		var err error
		switch val {
		case "-c":
			o.cFlag = true
		case "-d":
			o.dFlag = true
		case "-u":
			o.uFlag = true
		case "-i":
			o.iFlag = true
		case "-f":
			if notLast {
				o.fFlag, err = strconv.Atoi(cmdArgs[idx+1])
				if err != nil {
					return nil, err
				}
			}
		case "-s":
			if notLast {
				o.sFlag, err = strconv.Atoi(cmdArgs[idx+1])
				if err != nil {
					return nil, err
				}
			}
		}
		if isFile(val) && !isDigit(val) {
			if o.inputFN == "" {
				o.inputFN = val
			} else if o.inputFN != "" {
				o.outputFN = val
			}
		}
	}
	if (o.cFlag && o.dFlag) || (o.dFlag && o.uFlag) || (o.cFlag && o.uFlag) {
		return nil, errors.New("Inconsistent arguments")
	}
	return o, nil
}

func newOptions() *options {
	return &options{
		fFlag: -1,
		sFlag: -1,
	}
}

func readFromInput(option *options) ([]string, error) {
	if option.inputFN != "" {
		inStream, err := os.OpenFile(option.inputFN, os.O_RDONLY, 0755)
		if err != nil {
			return nil, err
		}
		defer inStream.Close()
		return readFromStream(inStream)
	}
	return readFromStream(os.Stdin)
}

func writeIntoOutput(option *options, data map[string]int) error {
	if option.outputFN != "" {
		outStream, err := os.OpenFile(option.outputFN, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return err
		}
		defer outStream.Close()
		return writeIntoStream(outStream, data)
	}
	return writeIntoStream(os.Stdin, data)
}

func writeIntoStream(outStream io.Writer, data map[string]int) error {
	for key := range data {
		_, err := outStream.Write([]byte(key + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}

func readFromStream(inStream io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(inStream)
	data := []string{}
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data, nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getCmdArgs() []string {
	return os.Args[1:]
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isFile(s string) bool {
	return !strings.HasPrefix(s, "-")
}

func checkUniqString(data []string, opt *options) map[string]int {
	checkMap := map[string]int{}

	for _, stringVal := range data {
		if _, exist := checkMap[stringVal]; exist {
			checkMap[stringVal]++
		} else {
			checkMap[stringVal] = 0
		}
	}
	return checkMap
}
