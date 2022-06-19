package actu

type BuildInter interface {
	calc()
	n(int, int) float64
	raw() [Row][Col]float64
}

type Actu struct {
	b BuildInter
}

func NewActu() *Actu {
	b := newBuild()
	b.calc()
	a := &Actu{b: b}
	return a
}

func (a *Actu) Data() [Row][Col]float64 {
	return a.b.raw()
}

func (a *Actu) Range(years int) [2][]float64 {
	out := [2][]float64{}
	if years <= 0 {
		return out
	}

	for i := 0; i < Row; i += years {
		var msum, wsum float64
		for j := i; j < i+1 && j+i < Row; j += 1 {

			msum += a.b.n(i+j, 1)
			wsum += a.b.n(i+j, 4)

		}
		out[0] = append(out[0], msum)
		out[1] = append(out[1], wsum)
	}

	return out

}
