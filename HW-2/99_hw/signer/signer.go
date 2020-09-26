package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

type OrderedData struct {
	order int
	data  string
}

func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}
	in := make(chan interface{})

	for _, currentJob := range jobs {
		out := make(chan interface{})
		wg.Add(1)
		go func(currentJob job, wg *sync.WaitGroup, in, out chan interface{}) {
			defer wg.Done()
			defer close(out)
			currentJob(in, out)
		}(currentJob, wg, in, out)
		in = out
	}
	wg.Wait()
}

func SingleHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	for val := range in {
		wg.Add(1)
		data := fmt.Sprintf("%v", val)
		md5Data := DataSignerMd5(data)

		go func(data string, md5Data string) {
			defer wg.Done()

			goNumber := 2
			result := "~"
			localWg := &sync.WaitGroup{}
			localOut := make(chan OrderedData, goNumber)

			countCrc32 := func(od OrderedData) {
				defer localWg.Done()
				od.data = DataSignerCrc32(od.data)
				localOut <- od
			}

			od := OrderedData{0, data}
			odmd5 := OrderedData{1, md5Data}

			localWg.Add(goNumber)
			go countCrc32(od)
			go countCrc32(odmd5)

			localWg.Wait()
			close(localOut)
			for odVal := range localOut {
				if odVal.order == 0 {
					result = odVal.data + result
				} else {
					result = result + odVal.data
				}
			}
			fmt.Println(result)
			out <- result
		}(data, md5Data)
	}
	wg.Wait()
}

func MultiHash(in, out chan interface{}) {
	goNumber := 6
	wg := &sync.WaitGroup{}

	for inVal := range in {
		localOut := make(chan OrderedData, goNumber)
		localWg := &sync.WaitGroup{}

		localWg.Add(goNumber)
		for i := 0; i < goNumber; i++ {
			od := OrderedData{i, inVal.(string)}
			go func(od OrderedData) {
				defer localWg.Done()
				od.data = DataSignerCrc32(strconv.Itoa(od.order) + od.data)
				localOut <- od
			}(od)
		}

		wg.Add(1)
		go func() {
			localWg.Wait()
			close(localOut)
			defer wg.Done()

			odSlice := make([]OrderedData, 0, goNumber)
			for odVal := range localOut {
				odSlice = append(odSlice, odVal)
				if len(odSlice) == goNumber {
					break
				}
			}

			strSlice := make([]string, goNumber, goNumber)
			for _, val := range odSlice {
				strSlice[val.order] = val.data
			}
			result := ""
			for _, val := range strSlice {
				result += val
			}
			out <- result
		}()
	}
	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	var resultSlice []string
	for val := range in {
		resultSlice = append(resultSlice, val.(string))
	}
	sort.Strings(resultSlice)
	result := ""
	for idx, val := range resultSlice {
		result += val
		if idx != len(resultSlice)-1 {
			result += "_"
		}
	}
	out <- result
}
