package actu

import (
	"github.com/cwxstat/life_death/data/loader"
)

type Loader interface {
	Calc()
	N(int, int) float64
	Raw() [loader.Row][loader.Col]float64
}

type Actu struct {
	b Loader
}

func NewActu() *Actu {

	b := loader.NewLoad()
	b.Calc()
	a := &Actu{b: b}
	return a
}

func (a *Actu) Data() [loader.Row][loader.Col]float64 {
	return a.b.Raw()
}

func (a *Actu) Bucket(years int) [2][]float64 {
	out := [2][]float64{}
	if years <= 0 {
		return out
	}

	for i := 0; i < loader.Row; i += years {
		var msum, wsum float64
		for j := i; j < i+1 && j+i < loader.Row; j += 1 {

			msum += a.b.N(i+j, 1)
			wsum += a.b.N(i+j, 4)

		}
		out[0] = append(out[0], msum)
		out[1] = append(out[1], wsum)
	}

	return out

}
