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
