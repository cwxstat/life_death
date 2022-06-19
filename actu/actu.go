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

func (a *Actu) Data() [][]string {
	return a.b.data[1:]
}
