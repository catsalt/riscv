// txt 文本自动生成 go function.
package rvspec

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var rs = []string{"rd", "rm", "csr", "rs1", "rs2", "rs3", "shamt", "uimm", "imm", "rl", "aq"}

// 位置比较固定
func aTasmfunc(k int, v string) string {
	vs := strings.Split(v, " ")
	f, l, finner := strings.ReplaceAll(strings.ToLower(vs[len(vs)-1]), ".", ""), len(vs), ""
	fouter := fmt.Sprintf("//%s %s: %s\nfunc %s(%s uint32)(code uint32){\ncode=", strconv.Itoa(k), vs[len(vs)-1], v,
		f, strings.Join(cTpars(vs), ","))
	switch l {
	case 4:
		s := []string{"<<12|", "<<7|"}
		finner += dTinner(l, vs, s)
	case 6:
		switch f {
		case "ecall":
			finner = "0b1110011"
			fouter = fmt.Sprintf("//%s %s: %s\nfunc %s()(code uint32){\ncode=", strconv.Itoa(k), vs[len(vs)-1], v, f)
		case "ebreak":
			finner = "0b100000000000001110011"
			fouter = fmt.Sprintf("//%s %s: %s\nfunc %s()(code uint32){\ncode=", strconv.Itoa(k), vs[len(vs)-1], v, f)
		default:
			s := []string{"<<20|", "<<15|", "<<12|", "<<7|"}
			finner += dTinner(l, vs, s)
		}
	case 7:
		s := []string{"<<25|", "<<20|", "<<15|", "<<12|", "<<7|"}
		finner += dTinner(l, vs, s)
	case 8:
		if f == "fence" {
			finner = "fm<<28 | pred<<24 | succ<<20 | rs1<<15 | 0b000 | rd<<7 | 0b0001111"
		} else {
			s := []string{"<<27|", "<<25|", "<<20|", "<<15|", "<<12|", "<<7|"}
			finner += dTinner(l, vs, s)
		}
	case 9:
		s := []string{"<<27|", "<<26|", "<<25|", "<<20|", "<<15|", "<<12|", "<<7|"}
		finner += dTinner(l, vs, s)
	default:
		fmt.Println(k, v, "!!!!!!!")
	}
	return fmt.Sprintf("%s %s\nreturn code}\n", fouter, finner)
}

func cTpars(vs []string) (pars []string) {
	for _, w := range rs {
		for i := len(vs) - 1; i >= 0; i-- {
			if w == vs[i] {
				pars = append(pars, w)
				break
			}
			if w == "imm" && strings.Contains(vs[i], "imm[") {
				pars = append(pars, "imm")
				break
			}
		}
	}
	return pars
}

// 位置比较固定
func dTinner(l int, vs, s []string) (inner string) {
	for i := 0; i < l-2; i++ {
		switch vs[i] {
		case "rd", "rm", "csr", "rs1", "rs2", "rs3", "shamt", "uimm", "imm", "rl", "aq":
			inner += vs[i] + s[i]
		case "imm[4:1|11]", "imm[4:0]", "imm[31:12]", "imm[11:0]",
			"imm[12|10:5]", "imm[11:5]", "imm[20|10:1|11|19:12]":
			inner += gTshift(vs[i])
		default:
			if len(vs[i]) == 6 {
				inner += "0b" + vs[i] + "<<26|"
			} else {
				inner += "0b" + vs[i] + s[i]
			}
		}
	}
	inner += "0b" + vs[l-2]
	return inner
}
func gTshift(imm string) string {
	switch imm {
	case "imm[4:1|11]":
		return "(imm &0b11110|imm>>11 &0b1)<<7|"
	case "imm[4:0]":
		return "imm &0b11111 <<7|"
	case "imm[31:12]":
		return "imm>>12<<12|"
	case "imm[20|10:1|11|19:12]":
		return "(imm>>20 &0b1 <<19|imm>>1 &(1<<10-1)<<9|imm>>11 &0b1 <<8|imm>>12 &(1<<8-1)) <<12|"
	case "imm[11:0]":
		return "imm &(1<<12-1) <<20|"
	case "imm[12|10:5]":
		return "(imm>>12 &0b1 <<6 | imm>>5 &(1<<6-1)) <<25|"
	case "imm[11:5]":
		return "imm>>5 &(1<<7-1) <<25|"
	}
	return imm
}

func zTrfile(in, out string) {
	f, _ := ioutil.ReadFile(in)
	srcasm := "package rvasm\n"
	for k, v := range strings.Split(string(f), "\r\n") {
		if v != "" {
			srcasm += aTasmfunc(k, v)
		}
	}
	ioutil.WriteFile(out, []byte(srcasm), 0777)
}

var xT = `
//185 EBREAK _ 000000000001 00000 000 00000 1110011 EBREAK
func ebreak() uint32 {
	return 0b100000000000001110011
}

//186 ECALL _ 000000000000 00000 000 00000 1110011 ECALL
func ecall() uint32 {
	return 0b1110011
}

//187 FENCE _ fm pred succ rs1 000 rd 0001111 FENCE
func fence(rd, fm, rs1, pred, succ uint32) (code uint32) {
	code = fm<<28 | pred<<24 | succ<<20 | rs1<<15 | 0b000 | rd<<7 | 0b0001111
	return code
}`
var yT = `
	"ebreak":   []string{},
	"ecall":    []string{},
	"fence":    []string{"rd", "fm", "rs1", "pred", "succ"},
`
