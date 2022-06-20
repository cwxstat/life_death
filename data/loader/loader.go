package loader

import (
	"github.com/cwxstat/life_death/data"
	"os"
	"strconv"
	"strings"
)

const (
	Row = 120
	Col = 7
	sex = 2
)

type build struct {
	mTotal  []float64
	wTotal  []float64
	skip    int
	data    [][]string
	numbers [Row][Col]float64
}

func NewBuild() *build {
	b := &build{}

	raw, err := myread()
	if err != nil {
		return b
	}
	b.data = mysplit(raw)
	return b
}

func getNum(skip *int, s string) float64 {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		*skip += 1
	}
	return num
}

func (b *build) Raw() [Row][Col]float64 {
	return b.numbers
}

func (b *build) N(i, j int) float64 {
	if i < 0 || i >= Row || j < 0 || j >= Col {
		return 0.0
	}
	return b.numbers[i][j]
}

func (b *build) Calc() {

	total := [sex][]float64{}
	sum := [sex]float64{}
	b.numbers = [120][7]float64{}
	for i, v := range b.data {
		if i <= 0 {
			continue
		}

		sum[0] += getNum(&b.skip, v[1])
		sum[1] += getNum(&b.skip, v[4])

		total[0] = append(total[0], sum[0])
		total[1] = append(total[1], sum[1])

		for j := 0; j < len(v); j += 1 {
			b.numbers[i-1][j] = getNum(&b.skip, v[j])
		}

	}

	b.mTotal = total[0]
	b.wTotal = total[1]

}

func myread() ([]byte, error) {
	b := make([]byte, 10000)

	f, err := os.Open(data.DataFile())
	if err != nil {
		return b, err
	}
	n, err := f.Read(b)

	return b[0:n], err
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
