package loader

import (
	"os"
	"strconv"
	"strings"

	"github.com/cwxstat/life_death/data"
)

const (
	Row = 120
	Col = 7
	sex = 2
)

type load struct {
	mTotal  []float64
	wTotal  []float64
	skip    int
	data    [][]string
	numbers [Row][Col]float64
}

func NewLoad() *load {
	b := &load{}

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

func (l *load) Raw() [Row][Col]float64 {
	return l.numbers
}

func (l *load) N(i, j int) float64 {
	if i < 0 || i >= Row || j < 0 || j >= Col {
		return 0.0
	}
	return l.numbers[i][j]
}

func (l *load) Calc() {

	total := [sex][]float64{}
	sum := [sex]float64{}
	l.numbers = [120][7]float64{}
	for i, v := range l.data {
		if i <= 0 {
			continue
		}

		sum[0] += getNum(&l.skip, v[1])
		sum[1] += getNum(&l.skip, v[4])

		total[0] = append(total[0], sum[0])
		total[1] = append(total[1], sum[1])

		for j := 0; j < len(v); j += 1 {
			l.numbers[i-1][j] = getNum(&l.skip, v[j])
		}

	}

	l.mTotal = total[0]
	l.wTotal = total[1]

}

func size(file string) (int64,error) {
	info,err := os.Stat(data.DataFile(file))
	if err != nil {
		return 0, err
	}
	return info.Size(),nil
}


func myread() ([]byte, error) {

	size, err := size("lifeexp.txt")
	if err != nil {
		return nil,nil
	}
	b := make([]byte, size)


	f, err := os.Open(data.DataFile("lifeexp.txt"))
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
