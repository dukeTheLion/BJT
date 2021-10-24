package main

import (
	"fmt"
	"io/ioutil"
	"math"
	rand2 "math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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
var example = make([]string, 4)
var descriptive = make([]string, 4)
var title string

func init() {
	list := []string{"schematic.txt", "example.txt", "title.txt", "descriptive.txt"}

	for i, txt := range list {
		dat, err := ioutil.ReadFile(txt)

		if err != nil {
			fmt.Print("Schematic error.")
		}

		if i == 0 {
			schematic = strings.Split(string(dat), "CUT")
		}

		if i == 1 {
			example = strings.Split(string(dat), "\nCUT")
		}

		if i == 2 {
			title = string(dat)
		}
		if i == 3 {
			descriptive = strings.Split(string(dat), "\nCUT")
		}

	}
}

func begin(osArgs []string) {
	args := make([]string, 0, 6)

	for _, arg := range osArgs {
		args = append(args, strings.ToUpper(arg))
	}

	if len(args) >= 2 {
		if "SR" == args[1] || "FR" == args[1] || "VDR" == args[1] || "FC" == args[1] {
			if "SR" == args[1] && len(args[2:]) == 5 {
				aux := converter(args, 5)

				stableEmitterPolarizationR(aux[0], aux[1], aux[2], aux[3], aux[4])
			} else if "SR" == args[1] && len(args[2:]) != 5 {
				fmt.Println("\n- ERRO -\nStable Emitter Polarization (res) has five arguments vcc, rb, rc, re, beta\n ")
			}

			if "FR" == args[1] && len(args[2:]) == 4 {
				aux := converter(args, 4)

				fixedPolarizationCircuitR(aux[0], aux[1], aux[2], aux[3])
			} else if "FR" == args[1] && len(args[2:]) != 4 {
				fmt.Println("\n- ERRO -\nFixed Polarization Circuit (res) has four arguments vcc, rb, rc, beta\n ")
			}

			if "FC" == args[1] && len(args[2:]) == 4 {
				aux := converter(args, 4)

				fixedPolarizationCircuitC(aux[0], aux[1], aux[2], aux[3])
			} else if "FC" == args[1] && len(args[2:]) != 4 {
				fmt.Println("\n- ERRO -\nFixed Polarization Circuit (current) has four arguments vcc, ib, ic, vce\n ")
			}

			if "VDR" == args[1] && len(args[2:]) == 6 {
				aux := converter(args, 6)

				voltageDividerBiasCircuit(aux[0], aux[1], aux[2], aux[3], aux[4], aux[5])
			} else if "VDR" == args[1] && len(args[2:]) != 5 {
				fmt.Println("\n- ERRO -\nVoltage Divider Bias Circuit (res) has six arguments vcc, r1, r2, rc, re, beta\n ")
			}

			if "SR" != args[1] && "FR" != args[1] && "VDR" != args[1] && "FC" != args[1] {
				fmt.Printf("The argment \"%s\" does not exist", args[1])
			}
		} else if "HELP" == args[1] {
			help(args)
		} else {
			fmt.Println("\n Without valid arguments. HELP form more information\n ")
		}
	} else {
		println("\n Without arguments. HELP form more information\n ")
	}
}

func help(args []string) {
	rand2.Seed(time.Now().Unix())
	rand := rand2.Intn(3)

	if len(args) == 2 {
		fmt.Print("\nS - Simplified help\nF - Full instruction\nSP - Simplified help in portuguese\n ")
	} else {
		if args[2] == "S" {
			fmt.Printf("\n%s%s\n%s%s%s%s\n ",
				title, example[rand], descriptive[0], descriptive[1], descriptive[2], descriptive[3])
		}
		if args[2] == "F" {
			fmt.Print(title)
			fmt.Printf("%s%s%s\n%s%s%s%s%s%s\n",
				example[rand], descriptive[0], schematic[0], descriptive[1], schematic[1], descriptive[2], schematic[2],
				descriptive[3], schematic[0])
		}
		if args[2] == "SP" {
			fmt.Println("\nNão implementado\n ")
		}
	}
}

func converter(args []string, n int) []float64 {
	aux := make([]float64, 0, n)
	ero := make([]string, 0, n)

	for _, arg := range args[2:] {
		if s, err := strconv.ParseFloat(arg, 64); err == nil {
			aux = append(aux, s)
		} else {
			ero = append(ero, arg)
		}
	}

	if len(aux) != n {
		fmt.Printf("Some value is not a number: %v\n", ero)
		os.Exit(2)
	}

	return aux
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
		} else if strings.Contains(s, "v") {
			fmt.Printf("%5s = %9.4f %sV\n", s, val[s], aux[s])
		} else {
			fmt.Printf("%5s = %9.4f %s\n", s, val[s], aux[s])
		}
	}
}

func fixedPolarizationCircuitR(vcc float64, rb float64, rc float64, beta float64) {
	fmt.Printf("\n- Fixed Polarization Circuit -\n%v", schematic[0])

	val := make(map[string]float64)
	val["ib"] = (vcc - 0.7) / rb
	val["ic"] = beta * val["ib"]
	val["vce"] = vcc - val["ic"]*rc
	val["vb"] = 0.7
	val["vc"] = val["vce"]
	val["vbc"] = val["vb"] - val["vc"]

	multipliers(map[string]float64{"vcc": vcc, "rb": rb, "rc": rc, "β": beta}, "vcc", "rb", "rc", "β")
	println()

	multipliers(val, "ib", "ic", "vbc", "vce", "vb", "vc")
	println()
}

func stableEmitterPolarizationR(vcc float64, rb float64, rc float64, re float64, beta float64) {
	fmt.Printf("\n- Stable Emitter Polarization -\n%v", schematic[1])

	val := make(map[string]float64)
	val["ib"] = (vcc - 0.7) / (rb + (beta+1)*re)
	val["ic"] = beta * val["ib"]
	val["vce"] = vcc - val["ic"]*(rc+re)
	val["vc"] = vcc - val["ic"]*rc
	val["ve"] = val["vc"] - val["vce"]
	val["vb"] = 0.7 + val["ve"]
	val["vbc"] = val["vb"] - val["vc"]

	multipliers(map[string]float64{"vcc": vcc, "rb": rb, "rc": rc, "re": re, "β": beta}, "vcc", "rb", "rc", "re", "β")
	println()

	multipliers(val, "ib", "ic", "vce", "vc", "ve", "vb", "vbc")
	println()
}

func voltageDividerBiasCircuit(vcc float64, r1 float64, r2 float64, rc float64, re float64, beta float64) {
	fmt.Printf("\n- Stable Emitter Polarization -\n%v", schematic[2])

	val := make(map[string]float64)
	val["rth"] = (r1 * r2) / (r1 + r2)
	val["eth"] = (r2 * vcc) / (r1 + r2)
	val["ib"] = (val["eth"] - 0.7) / (val["rth"] + ((beta + 1) * re))
	val["ic"] = beta * val["ib"]
	val["vce"] = vcc - val["ic"]*(rc+re)

	multipliers(map[string]float64{"vcc": vcc, "r1": r1, "r2": r2, "rc": rc, "re": re, "β": beta}, "vcc", "r1", "r2", "rc", "re", "β")
	println()

	multipliers(val, "rth", "eth", "ib", "ic", "vce")
	println()
}

func fixedPolarizationCircuitC(vcc float64, ib float64, ic float64, vce float64) {
	fmt.Printf("\n- Fixed Polarization Circuit -\n%v", schematic[0])

	val := make(map[string]float64)
	val["rb"] = (vcc - 0.7) / ib
	val["rc"] = (vcc - vce) / ic
	val["β"] = ic / ib

	multipliers(map[string]float64{"vcc": vcc, "ib": ib, "ic": ic, "vce": vce}, "vcc", "ib", "ic", "vce")
	println()

	multipliers(val, "rb", "rc", "β")
	println()
}

func main() {
	begin(os.Args)
}
