package main

import (
	"fmt"
	"github.com/cwxstat/life_death/actu"
)

func main() {
	a := actu.NewActu()
	fmt.Println(a.CSV())

}
