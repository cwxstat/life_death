package actu

type Actu struct {
	b *build
}

func NewActu() *Actu {
	b := newBuild()
	b.calc()
	a := &Actu{b: b}
	return a
}

func (a *Actu) Data() [Row][Col]float64 {
	return a.b.numbers
}

func (a *Actu) Range(years int) [2][]float64 {
	out := [2][]float64{}

	for i := 0; i < len(a.b.numbers); i += years {
		var msum, wsum float64
		for j := i; j < i+1 && j+i < len(a.b.numbers); j += 1 {

			msum += a.b.numbers[i+j][1]
			wsum += a.b.numbers[i+j][4]
		}
		out[0] = append(out[0], msum)
		out[1] = append(out[1], wsum)
	}

	return out

}
