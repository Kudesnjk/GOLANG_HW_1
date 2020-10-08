package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type options struct {
	PrintEntriesCount bool
	PrintRepeated     bool
	PrintUnRepeated   bool
	WithoutRegister   bool
	SkipNumFields     int
	SkipNumChars      int
	inputFileName     string
	outputFileName    string
}

func main() {
	cmdArgs := getCmdArgs()
	option, err := argsToOptions(cmdArgs)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := readFromInput(option)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := checkUniqString(data, option)
	err = writeIntoOutput(option, result)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func argsToOptions(cmdArgs []string) (*options, error) {
	o := &options{}
	isDigit := func(s string) bool {
		_, err := strconv.Atoi(s)
		return err == nil
	}

	isFile := func(s string) bool {
		return !strings.HasPrefix(s, "-")
	}

	for idx, val := range cmdArgs {
		notLast := idx+1 < len(cmdArgs)
		var err error
		switch val {
		case "-c":
			o.PrintEntriesCount = true
		case "-d":
			o.PrintRepeated = true
		case "-u":
			o.PrintUnRepeated = true
		case "-i":
			o.WithoutRegister = true
		case "-f":
			if notLast {
				o.SkipNumFields, err = strconv.Atoi(cmdArgs[idx+1])
				if err != nil || o.SkipNumFields <= 0 {
					return nil, errors.New("Use -f flag with num_fields > 0")
				}
			}
		case "-s":
			if notLast {
				o.SkipNumChars, err = strconv.Atoi(cmdArgs[idx+1])
				if err != nil || o.SkipNumChars <= 0 {
					return nil, errors.New("Use -s flag with num_chars > 0")
				}
			}
		}
		if isFile(val) && !isDigit(val) {
			if o.inputFileName == "" {
				o.inputFileName = val
			} else if o.inputFileName != "" {
				o.outputFileName = val
			}
		}
	}
	if (o.PrintEntriesCount && o.PrintRepeated) || (o.PrintRepeated && o.PrintUnRepeated) || (o.PrintEntriesCount && o.PrintUnRepeated) {
		return nil, errors.New(`Use -c | -d | -u flags apart due to it's meaning
		-c - count number of repeats of a string in the input
		-d - print only those lines that were repeated in the input
		-u - print only those lines that were not repeated in the input data
		`)
	}
	return o, nil
}

func readFromInput(option *options) ([]string, error) {
	if option.inputFileName != "" {
		inStream, err := os.OpenFile(option.inputFileName, os.O_RDONLY, 0755)
		if err != nil {
			return nil, err
		}
		defer inStream.Close()
		return readFromStream(inStream)
	}
	return readFromStream(os.Stdin)
}

func writeIntoOutput(option *options, data []string) error {
	if option.outputFileName != "" {
		outStream, err := os.OpenFile(option.outputFileName, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return err
		}
		defer outStream.Close()
		return writeIntoStream(outStream, data)
	}
	return writeIntoStream(os.Stdout, data)
}

func writeIntoStream(outStream io.Writer, data []string) error {
	for _, str := range data {
		_, err := outStream.Write([]byte(str + "\n"))
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

func getCmdArgs() []string {
	return os.Args[1:]
}

func checkUniqString(data []string, opt *options) []string {
	checkMap := map[string]int{}
	var returnData []string

	checkIFSkipNumCharss := func(str string, opt *options) (newStr string) {
		switch {
		case opt.SkipNumFields > 0:
			flagValue := opt.SkipNumFields
			for idx, char := range str {
				if char == ' ' && flagValue > 0 {
					newStr = str[idx+1:]
					flagValue--
				}
			}
		case opt.SkipNumChars > 0:
			flagValue := opt.SkipNumChars
			for idx := range str {
				if flagValue > 0 {
					newStr = str[idx+1:]
					flagValue--
				}
			}
		case opt.WithoutRegister:
			if newStr == "" {
				newStr = strings.ToLower(str)
			} else {
				newStr = strings.ToLower(newStr)
			}
		}
		if newStr == "" {
			newStr = str
		}
		return
	}

	for _, str := range data {
		str = checkIFSkipNumCharss(str, opt)
		if _, exist := checkMap[str]; exist {
			checkMap[str]++
		} else {
			checkMap[str] = 1
		}
	}

	for _, originStr := range data {
		str := checkIFSkipNumCharss(originStr, opt)
		if _, exist := checkMap[str]; exist {
			isRepeated := checkMap[str] > 1
			switch {
			case opt.PrintEntriesCount:
				returnData = append(returnData, strconv.Itoa(checkMap[str])+" "+originStr)
				delete(checkMap, str)
			case opt.PrintRepeated && isRepeated:
				returnData = append(returnData, originStr)
				delete(checkMap, str)
			case opt.PrintUnRepeated && !isRepeated:
				returnData = append(returnData, originStr)
				delete(checkMap, str)
			case !opt.PrintRepeated && !opt.PrintUnRepeated && !opt.PrintEntriesCount:
				returnData = append(returnData, originStr)
				delete(checkMap, str)
			}
		}
	}
	return returnData
}
