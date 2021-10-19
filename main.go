package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

const (
	k = 1000.0
	M = 1000000.0
	G = 1000000000.0
	T = 1000000000000.0

	m = 0.001
	u = 0.000001
	n = 0.000000001
	p = 0.000000000001
)

var schematic = make([]string, 3)

func init() {
	dat, err := ioutil.ReadFile("schematic.txt")

	if err != nil {
		fmt.Print("Schematic error.")
	}

	schematic = strings.Split(string(dat), "CUT")
}

func multipliers(val map[string]float64) {
	for key, value := range val {
		aux := ""

		if math.Abs(value) > k {
			val[key] /= 1000
			aux = "k"
		}
		if math.Abs(value) > T {
			val[key] /= 1000
			aux = "T"
		}
		if math.Abs(value) > M {
			val[key] /= 1000
			aux = "M"
		}
		if math.Abs(value) > G {
			val[key] /= 1000
			aux = "G"
		}
		if math.Abs(value) < m*1000 {
			val[key] *= 1000
			aux = "m"
		}
		if math.Abs(value) < u*1000 {
			val[key] *= 1000
			aux = "u"
		}
		if math.Abs(value) < n*1000 {
			val[key] *= 1000
			aux = "n"
		}
		if math.Abs(value) < p*1000 {
			val[key] *= 1000
			aux = "p"
		}

		if strings.Contains(key, "i") {
			fmt.Printf("%3s = %9.4f %sA\n", key, val[key], aux)
		} else {
			fmt.Printf("%3s = %9.4f %sV\n", key, val[key], aux)
		}
	}
}

func fixedPolCircuitR(vcc float64, rb float64, rc float64, beta float64) {
	fmt.Printf("- Fixed Polarization Circuit -\n%v", schematic[0])

	val := make(map[string]float64)
	val["ib"] = (vcc - 0.7) / rb
	val["ic"] = beta * val["ib"]
	val["vce"] = vcc - val["ic"]*rc
	val["vb"] = 0.7
	val["vc"] = val["vce"]
	val["vbc"] = val["vb"] - val["vc"]

	multipliers(val)

}

func main() {
	fixedPolCircuitR(700, 68*k, 0.82*k, 125)
}
