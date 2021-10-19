package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
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

func multipliers(val map[string]float64, order ...string) {
	aux := make(map[string]string)

	for key, value := range val {

		temp := ""

		if math.Abs(value) >= k {
			val[key] /= 1000
			temp = "k"
		}
		if math.Abs(value) >= T {
			val[key] /= 1000
			temp = "T"
		}
		if math.Abs(value) >= M {
			val[key] /= 1000
			temp = "M"
		}
		if math.Abs(value) >= G {
			val[key] /= 1000
			temp = "G"
		}
		if math.Abs(value) <= m*1000 {
			val[key] *= 1000
			temp = "m"
		}
		if math.Abs(value) <= u*1000 {
			val[key] *= 1000
			temp = "u"
		}
		if math.Abs(value) <= n*1000 {
			val[key] *= 1000
			temp = "n"
		}
		if math.Abs(value) <= p*1000 {
			val[key] *= 1000
			temp = "p"
		}

		aux[key] = temp
	}

	for _, s := range order {
		if strings.Contains(s, "i") {
			fmt.Printf("%5s = %9.4f %sA\n", s, val[s], aux[s])
		} else if strings.Contains(s, "r") {
			fmt.Printf("%5s = %9.4f %sΩ\n", s, val[s], aux[s])
		} else {
			fmt.Printf("%5s = %9.4f %sV\n", s, val[s], aux[s])
		}
	}
}

func fixedPolCircuitR(vcc float64, rb float64, rc float64, beta float64) {
	fmt.Printf("\n- Fixed Polarization Circuit -\n%v", schematic[0])

	val := make(map[string]float64)
	val["ib"] = (vcc - 0.7) / rb
	val["ic"] = beta * val["ib"]
	val["vce"] = vcc - val["ic"]*rc
	val["vb"] = 0.7
	val["vc"] = val["vce"]
	val["vbc"] = val["vb"] - val["vc"]

	multipliers(map[string]float64{"vcc": vcc, "rb": rb, "rc": rc}, "vcc", "rb", "rc")
	println()

	multipliers(val, "ib", "ic", "vbc", "vce", "vb", "vc")
	println()
}

func stableEmitterPolR(vcc float64, rb float64, rc float64, re float64, beta float64) {
	fmt.Printf("\n- Stable Emitter Polarization -\n%v", schematic[1])

	val := make(map[string]float64)
	val["ib"] = (vcc - 0.7) / (rb + (beta+1)*re)
	val["ic"] = beta * val["ib"]
	val["vce"] = vcc - val["ic"]*(rc+re)
	val["vc"] = vcc - val["ic"]*rc
	val["ve"] = val["vc"] - val["vce"]
	val["vb"] = 0.7 + val["ve"]
	val["vbc"] = val["vb"] - val["vc"]

	multipliers(map[string]float64{"vcc": vcc, "rb": rb, "rc": rc, "re": re}, "vcc", "rb", "rc", "re")
	println()

	multipliers(val, "ib", "ic", "vce", "vc", "ve", "vb", "vbc")
	println()
}

func main() {
	args := make([]string, 0, 6)

	for _, arg := range os.Args {
		args = append(args, strings.ToUpper(arg))
	}

	if len(args) > 2 {
		if "SEPR" == args[1] && len(args[2:]) == 5 {
			aux := make([]float64, 0, 5)

			for _, arg := range args[2:] {
				if s, err := strconv.ParseFloat(arg, 64); err == nil {
					aux = append(aux, s)
				}
			}

			stableEmitterPolR(aux[0], aux[1], aux[2], aux[3], aux[4])
		} else if "SEPR" == args[1] && len(args[2:]) != 5 {
			fmt.Println("\n- ERRO -\nStable Emitter Polarization (res) has five arguments vcc, rb, rc, re, beta")
		}

		if "FPCR" == args[1] && len(args[2:]) == 4 {
			aux := make([]float64, 0, 4)

			for _, arg := range args[2:] {
				if s, err := strconv.ParseFloat(arg, 64); err == nil {
					aux = append(aux, s)
				}
			}

			fixedPolCircuitR(aux[0], aux[1], aux[2], aux[3])
		} else if "FPCR" == args[1] && len(args[2:]) != 4 {
			fmt.Println("\n- ERRO -\nFixed Polarization Circuit (res) has four arguments vcc, rb, rc, beta ")
		}
	} else if "HELP" == args[1] {
		fmt.Println("\nType => code values...\n\n" +
			"Stable Emitter Polarization (res) => SEPR vcc rb rc re beta\n" +
			"Fixed Polarization Circuit (res) => FPCR vcc rb rc beta\n\n" +
			"EXAMPLE: go run example.go FPCR 10 250000 20000 100\n ")
	} else {
		println("Without arguments.")
	}

	/*stableEmitterPolR(20, 430000, 2000, 1000, 50)
	fixedPolCircuitR(16, 470000, 2700, 90)*/
}
