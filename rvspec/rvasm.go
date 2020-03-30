// rvasm.go
package rvspec

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var ps = []string{"rd", "rm", "csr", "rs1", "rs2", "rs3", "shamt", "uimm", "imm", "rl", "aq"}

//指令名:形式参数
var psmap = make(map[string][]string)

//相同形式参数:指令集合
var mps = make(map[string][]string)

func rvPars(v string) {
	vs, m := strings.Split(v, " "), []string{}
	for _, w := range ps {
		for i := len(vs) - 1; i >= 0; i-- {
			if w == vs[i] {
				m = append(m, w)
				break
			}
			if w == "imm" && strings.Contains(vs[i], "imm[") {
				m = append(m, w)
				break
			}
		}
	}
	f := strings.ReplaceAll(strings.ToLower(vs[len(vs)-1]), ".", "")
	n := strings.Join(m, ",")
	if v, ok := mps[n]; ok {
		mps[n] = append(v, f)
	} else {
		mps[n] = []string{f}
	}
	psmap[f] = m
}
func rvOrder() []string {
	t := []string{}
	for k, _ := range psmap {
		t = append(t, k)
	}
	sort.Strings(t)
	return t
}
func rvArgs(l int) string {
	switch l {
	case 0:
		return "()\n"
	case 1:
		return "(args[0])\n"
	case 2:
		return "(args[0],args[1])\n"
	case 3:
		return "(args[0],args[1],args[2])\n"
	case 4:
		return "(args[0],args[1],args[2],args[3])\n"
	case 5:
		return "(args[0],args[1],args[2],args[3],args[4])\n"
	default:
		return ""
	}
	return ""
}

func aEpars() string {
	out := "func aEpars(fn string)(pars []string,err error){\nswitch fn{"
	for k, v := range mps {
		sort.Strings(v)
		out += fmt.Sprintf("case \"%s\":\npars=[]string{\"%s\"}\n",
			strings.Join(v, "\",\""), strings.Join(strings.Split(k, ","), "\",\""))

	}
	out += "default: err=fmt.Errorf(\"aEpars......no %s\",fn)}\nreturn pars,err}\n"
	return out
}
func bEncode() string {
	out := "func bEncode(fn string,args []uint32)(code uint32){\nswitch fn{"
	for _, v := range rvOrder() {
		out += fmt.Sprintf("case \"%s\":\ncode=%s%s", v, v, rvArgs(len(psmap[v])))
	}
	out += "default: fmt.Printf(\"bEncode......no %s\",fn)}\nreturn code}\n"
	return out
}

//生成 函数名 : 形式参数切片 的map
func cErvMap() string {
	out := "var cErvMap = map[string][]string{\n"
	for _, v := range rvOrder() {
		switch v {
		case "ecall", "ebreak":
			out += fmt.Sprintf("\"%s\":[]string{},\n", v)
		default:
			out += fmt.Sprintf("\"%s\":[]string{\"%s\"},\n", v, strings.Join(psmap[v], "\",\""))
		}

	}
	return out + "}\n"
}

func rvFile(in, out string) {
	f, _ := ioutil.ReadFile(in)
	// srcasm := "package rvasm\n"
	rv := strings.Split(string(f), "\r\n")
	for _, v := range rv {
		if v != "" {
			rvPars(v)
		}
	}
	srcmap := "package rvasm\nimport(\"fmt\")\n" + bEncode() + cErvMap()
	ioutil.WriteFile(out, []byte(srcmap), 0777)
}
