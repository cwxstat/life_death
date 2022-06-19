package main

import (
	"os"
	"strings"
	"strconv"
)

type build struct {

	mTotal []float64
	wTotal []float64
	skip int
	data [][]string
}

func newBuild() *build {
	b := &build{}

	raw, err := myread()
	if err != nil {
		return b
	}
	b.data = mysplit(raw)
	return b
}

func getNum(skip *int,s string) float64 {
	num,err := strconv.ParseFloat(s, 64)
	if err != nil {
		*skip += 1
	}
	return num
}


func (b *build)Calc() *build {

	total := [2][]float64{}
	sum := [2]float64{}
	for i,v := range b.data {
		if i <= 0 {
			continue
		}
		
		sum[0] += getNum(&b.skip,v[1])
		sum[1] += getNum(&b.skip,v[4])

		total[0] = append(total[0],sum[0])
		total[1] = append(total[1],sum[1])

	}

	b.mTotal = total[0]
	b.wTotal = total[1]
	return b
}


func myread() ([]byte, error) {
	b := make([]byte, 10000)

	f, err := os.Open("data")
	if err != nil {
		return b, err
	}
	n, err := f.Read(b)

	return b[0:n], nil
}

func mysplit(b []byte) [][]string {
	data := [][]string{}
	s := string(b)
	lines := strings.Split(s, "\n")
	for _, v := range lines {
		data = append(data, strings.Split(v, "\t"))

	}
	return data
}
