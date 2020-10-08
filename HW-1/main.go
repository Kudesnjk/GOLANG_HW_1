package main

import (
	"bufio"
	"errors"
	"flag"
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
	option, err := InitFlags()
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := readFromInput(&option)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := checkUniqString(data, &option)
	err = writeIntoOutput(&option, result)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func InitFlags() (options, error) {
	flagCPtr := flag.Bool("c", false, "for number of occurrences of lines in the input")
	flagDPtr := flag.Bool("d", false, "print only those lines that were repeated in the input data")
	flagUPtr := flag.Bool("u", false, "print only those lines that have not been repeated in the input data")

	flagIPtr := flag.Bool("i", false, "case-insensitive")
	flagFPtr := flag.Int("f", 0, "ignore the first num_fields fields in the line")
	flagSPtr := flag.Int("s", 0, "ignore the first num_chars characters in the string")

	flag.Parse()

	opt := options{
		PrintEntriesCount: *flagCPtr,
		PrintRepeated:     *flagDPtr,
		PrintUnRepeated:   *flagUPtr,
		SkipNumFields:     *flagFPtr,
		SkipNumChars:      *flagSPtr,
		WithoutRegister:   *flagIPtr,
	}
	if !opt.PrintEntriesCount && opt.PrintRepeated && opt.PrintUnRepeated ||
		opt.PrintEntriesCount && !opt.PrintRepeated && opt.PrintUnRepeated ||
		opt.PrintEntriesCount && opt.PrintRepeated && !opt.PrintUnRepeated {
		return opt, errors.New("invalid arguments passed")
	}

	if !opt.PrintEntriesCount && !opt.PrintUnRepeated && !opt.PrintRepeated {
		opt.PrintUnRepeated = true
		opt.PrintRepeated = true
	}

	return opt, nil
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
