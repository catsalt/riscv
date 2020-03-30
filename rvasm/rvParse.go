//32位指令格式
//20200323:
// 增加jPmargin()对imm的,不同类型(I,U,J,B)的取值范围的检查;
// ld...sd...交换参数顺序,iPtranslate()...
// lr.d sc.d... 参数设置...
// fence, amo... 不知道怎么设置
package rvasm

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

//bigendian
// inst[1:0] 11
// inst[4:2] 000 	001 		010 		011 		100 	101 		110 				111
// inst[6:5] 																					(>32b)
// 00 		LOAD 	LOAD-FP 	custom-0 	MISC-MEM 	OP-IMM 	AUIPC 		OP-IMM-32 			48b
// 01 		STORE 	STORE-FP 	custom-1 	AMO 		OP 		LUI 		OP-32 				64b
// 10 		MADD 	MSUB 		NMSUB 		NMADD 		OP-FP 	reserved 	custom-2/rv128 		48b
// 11 		BRANCH 	JALR 		reserved 	JAL 		SYSTEM 	reserved 	custom-3/rv128 		≥80b

// 去掉整个src的注释, 空行, 行首尾空格
func aPtrimCommentS(src string) (cptcs []string) {
	for k, v := range strings.Split(src, "\r\n") {
		v, _ = bPformat(k, v)
		if len(v) > 0 {
			cptcs = append(cptcs, v)
		}
	}
	return cptcs
}

// 拆分(指令,注释), 格式化指令(用逗号分隔,去掉多余空格)
func bPformat(k int, s string) (ins, comment string) {
	ins = s
	for i, c := range s {
		if c == ';' || c == '#' || c == '/' {
			ins, comment = s[:i], s[i:]
			break
		}
	}
	ins = strings.TrimSpace(ins)
	i := strings.Index(ins, " ")
	if i >= 0 {
		ins = ins[:i] + "," + ins[i:]
		ins = strings.ReplaceAll(ins, " ", "")
	}
	if strings.Contains(comment, "//") {
		return ins, "\t\t" + comment
	}
	return ins, "\t\t//" + comment
}

var mapLable = make(map[string]int)

// 第一次生成lable表,供分解双行伪代码确认参数
func dPtable(srcLines []string) {
	for k, v := range srcLines {
		if strings.Contains(v, ":") {
			v = strings.TrimRight(v, ":")
			mapLable[v] = k
		}
	}
}

// 分解双行伪代码为两行!!!分解后,lable位置会变化
func ePtwoPseudo(srcLines []string) (eptp []string) {
	dPtable(srcLines)
	for k, v := range srcLines {
		ins, comment := bPformat(k, v)
		one, two, three, four := "", false, "", false
		d := strings.Split(ins, ",")
		switch d[0] {
		case "call":
			if len(d) == 2 {
				one, two, three = fmt.Sprintf("auipc x1,%s", d[1]), true, fmt.Sprintf("jalr x1,%s(x1)", d[1])
			}
		case "tail":
			if len(d) == 2 {
				one, two, three = fmt.Sprintf("auipc x6,%s", d[1]), true, fmt.Sprintf("jalr x0,%s(x6)", d[1])
			}
		case "li":
			if len(d) == 3 {
				one, two, three = fmt.Sprintf("lui %s,%s", d[1], d[2]), true, fmt.Sprintf("addi %s,%s,%s", d[1], d[1], d[2])
			}
		case "lb", "lh", "lw", "ld":
			if len(d) == 3 { //mapLable, 检查是否有lable参数, 有就是伪代码!
				if _, ok := mapLable[d[2]]; ok {
					one, two, three = fmt.Sprintf("auipc %s,%s", d[1], d[2]), true, fmt.Sprintf("%s(%s)", v, d[1])
				}
			}
			four = true
		case "sb", "sh", "sw", "sd", "fsw", "fsd", "flw", "fld":
			if len(d) == 4 {
				if _, ok := mapLable[d[2]]; ok {
					one, two, three = fmt.Sprintf("auipc %s,%s", d[3], d[2]), true, fmt.Sprintf("%s %s,%s(%s)", d[0], d[1], d[2], d[3])
				}
			}
			four = true
		case "lla", "la":
			if len(d) == 3 {
				if _, ok := mapLable[d[2]]; ok {
					one, two, three = fmt.Sprintf("auipc %s,%s", d[1], d[2]), true, fmt.Sprintf("addi %s,%s,%s", d[1], d[1], d[2])
				}
			}
			four = true
		default:
			four = true
		}
		if two {
			eptp = append(eptp, one+comment+v, three+comment+v)
		} else {
			if four {
				// log.Println("ePtwoPseudo......忽略", k, v)
				eptp = append(eptp, v)
			} else {
				log.Println("ePtwoPseudo......", k, v)
			}
		}
	}
	fPtalbe(eptp)
	return eptp
}

// 第二次生成lable表, 因为双行伪代码增加了行数, 导致标签位置变化
func fPtalbe(srcLines []string) {
	for k, v := range srcLines {
		if strings.Contains(v, ":") {
			srcLines[k] = "addi x0,x0,0\t\t//" + v
			v = strings.TrimRight(v, ":")
			mapLable[v] = k
		}
	}
}

// 分解单行伪代码为一行
func hPseudo(src []string) []string {
	switch src[0] {
	case "beqz", "bnez", "bgez", "bltz", "neg", "negw", "snez", "sgtz":
		return []string{mapPseudoRv[src[0]], src[1], "x0", src[2]}
	case "fmvs", "fabss", "fnegs", "fmvd", "fabsd", "fnegd":
		return []string{mapPseudoRv[src[0]], src[1], src[2], src[2]}
	case "bgtz", "blez", "csrw", "csrs", "csrc", "csrwi", "csrsi", "csrci":
		return []string{mapPseudoRv[src[0]], "x0", src[1], src[2]}
	case "bgt", "ble", "bgtu", "bleu":
		return []string{mapPseudoRv[src[0]], src[2], src[1], src[3]}
	case "csrr", "sltz":
		return []string{mapPseudoRv[src[0]], src[1], src[2], "x0"}
	case "sext.w", "mv":
		return []string{mapPseudoRv[src[0]], src[1], src[2], "0"}
	case "fence":
		log.Println("hPseudo......fence 用法参数的设定需要进一步阅读技术手册2.7_26页", src)
		return []string{"fence", "iorw", "iorw"}
	case "nop":
		return []string{"addi", "x0", "x0", "0"}
	case "not":
		return []string{mapPseudoRv[src[0]], src[1], src[2], "-1"}
	case "seqz":
		return []string{mapPseudoRv[src[0]], src[1], src[2], "1"}
	case "j":
		return []string{mapPseudoRv[src[0]], "x0", src[1]}
	case "jal":
		return []string{mapPseudoRv[src[0]], "x1", src[1]}
	case "jalr":
		return []string{mapPseudoRv[src[0]], "x1", "0(" + src[1] + ")"}
	case "jr":
		return []string{mapPseudoRv[src[0]], "x0", "0(" + src[1] + ")"}
	case "ret":
		return []string{mapPseudoRv[src[0]], "x0", "0(x1)"}
	default:
		log.Println("hPseudo......遗漏", src)
		return []string{}
	}
}

// 拆分(函数名,参数)交换括号内的值的顺序; 转换位小写;
func hPtrimBar(s string) (dptb []string) {
	s = strings.ToLower(s)
	for _, v := range strings.Split(s, ",") {
		i, j := strings.Index(v, "("), strings.Index(v, ")")
		if i != -1 && j != -1 {
			dptb = append(dptb, v[i+1:j])
			dptb = append(dptb, v[:i])
		} else {
			dptb = append(dptb, v)
		}
	}
	return dptb
}

// 将lable转换为数值;此处用于32位字(*4 byte),指令格式;pc寻址范围,jPmargin判定
func hPlable(k int, args []string) {
	for x, arg := range args {
		if l, ok := mapLable[arg]; ok {
			args[x] = strconv.Itoa((l - k) * 4)
		}
	}
}

// 单行转换为riscv.go 和 bin  格式.
func iPtranslate(k int, v string) (outgo string, outbin []byte) {
	ins, comment := bPformat(k, v)
	args := hPtrimBar(ins)
	args[0] = strings.ReplaceAll(args[0], ".", "")
	pars, ok := mapPseudo[args[0]] //如果是单行伪指令...
	if ok && len(pars) == len(args)-1 {
		args = hPseudo(args) //转换
		comment += ins
	}
	pars, ok = cErvMap[args[0]]
	if !ok || len(pars) != len(args)-1 {
		log.Println("iPtranslate......Wrong instruction: ", k, v)
		return outgo, outbin
	}
	switch args[0] { //sd rt,imm(rs1)-交换参数顺序-sd rs1,rt,imm
	case "sb", "sh", "sw", "sd", "fsw", "fsd":
		args[1], args[2] = args[2], args[1]
	}
	hPlable(k, args)
	outgo = args[0] + "("
	outArg := []uint32{}
	for i := 0; i < len(pars); i++ {
		j, err := jPmargin(args[0], pars[i], args[i+1])
		if err != nil {
			log.Println(k, v, err)
		}
		outgo += strconv.Itoa(j) + ","
		outArg = append(outArg, uint32(j))
	}
	outgo = strings.TrimRight(outgo, ",") + ")" + comment + "\n"
	outbin = mPbigEnd(bEncode(args[0], outArg))
	return outgo, outbin
}

// 根据形式参数parameter,检查实际参数argument是否超出定义区间, 输出argument的int值
// pars - {"rd", "rm", "csr", "rs1", "rs2", "rs3", "shamt", "uimm", "imm", "rl", "aq"}
func jPmargin(insName, par, arg string) (j int, err error) {
	over := func(max, min, j int, err error) error { //v[i] 对应 src[i+1]
		if err != nil {
			return err
		}
		if j > max || j < min {
			return fmt.Errorf("iPtranslate_jPmargin......应该! %d > %d > %d ", max, j, min)
		}
		return nil
	}
	switch par {
	case "rd", "rs1", "rs2", "rs3":
		j, err = jPregister(arg)
		err = over(31, 0, j, err)
	case "uimm":
		j, err = jPimm(arg)
		err = over(31, 0, j, err)
	case "imm":
		j, err = jPimm(arg)
		switch insName {
		case "lui", "auipc": //32位之高20位[12~31]
			err = over(1<<31-1, -1<<31, j, err)
		case "beq", "bne", "blt", "bltu", "bge", "bgeu": //13位地址高12位[1~12]
			err = over(1<<12-1, -1<<12, j, err)
		case "jal": //21位地址高20位[1~20]
			err = over(1<<20-1, -1<<20, j, err)
		default: //12位地址
			err = over(1<<11-1, -1<<11, j, err)
		}
	case "shamt":
		j, err = jPimm(arg)
		switch insName {
		case "slli", "srli", "srai":
			err = over(63, 0, j, err)
		default:
			err = over(31, 0, j, err)
		}
	case "rm": //还不清楚...用法
		j, err = jPimm(arg)
		err = over(7, 0, j, err)
	case "csr": //还不清楚...用法
		j, err = jPimm(arg)
		err = over(1<<12-1, 0, j, err)
	case "rl", "aq": //还不清楚...用法
		j, err = jPimm(arg)
		err = over(1, 0, j, err)
	default:
		err = fmt.Errorf("iPtranslate_jPmargin......遗漏形式参数")
	}
	return j, err
}

// 文本寄存器转数值
func jPregister(arg string) (int, error) {
	arg = strings.ToLower(arg)
	if p, ok := rvrgmap[arg]; ok {
		return p, nil
	}
	if p, ok := rvrgmapx[arg]; ok {
		return p, nil
	}
	return 0, fmt.Errorf("jPregister......%s", arg)
}

// 文本数值转数值
func jPimm(arg string) (int, error) {
	switch {
	case strings.Contains(arg, "0b"):
		return kPbin(arg)
	case strings.Contains(arg, "0x"):
		return kPhex(arg)
	case strings.Contains(arg, "0O"):
		return kPoct(arg)
	default:
		return kPatoi(arg)
	}
}

// 二进制转换
func kPbin(arg string) (p int, err error) {
	arg = strings.TrimLeft(arg, "0b")
	for i, m := len(arg)-1, 0; i >= 0; i-- {
		switch {
		case arg[i] == '0':
		case arg[i] == '1':
			p += 0b1 << m
		default:
			err = fmt.Errorf("kPbin......Bin should be 0-1: 0b%s", arg)
			return
		}
		m++
	}
	return p, err
}

// 八进制
func kPoct(arg string) (p int, err error) {
	arg = strings.TrimLeft(arg, "0o")
	for i, m := len(arg)-1, 0; i >= 0; i-- {
		switch {
		case arg[i] < 56 && arg[i] > 47:
			p += int(arg[i]-48) << m
		default:
			err = fmt.Errorf("kPoct......Oct should be 0-7: 0o%s", arg)
			return
		}
		m += 3
	}
	return p, err
}

// 十六进制
func kPhex(arg string) (p int, err error) {
	arg = strings.TrimLeft(arg, "0x")
	for i, m := len(arg)-1, 0; i >= 0; i-- {
		switch {
		case arg[i] < 58 && arg[i] > 47:
			p += int(arg[i]-48) << m
		case arg[i] < 103 && arg[i] > 96:
			p += int(arg[i]-87) << m
		default:
			err = fmt.Errorf("kPhex......Hex be 0-9 a-f: 0x%s", arg)
			return
		}
		m += 4
	}
	return p, err
}

// 十进制
func kPatoi(arg string) (p int, err error) {
	p, err = strconv.Atoi(arg)
	if err != nil {
		if strings.Contains(arg, "x") {
			err = fmt.Errorf("kPatoi......Register should be (x0~x31): %s", arg)
			fmt.Println(err, arg)
		}
		err = fmt.Errorf("kPatoi.....Digit should be 0~9: %s", arg)
	}
	return p, err
}
func mPbigEnd(v uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v)
	return b
}
func mPlittleEnd(v uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, v)
	return b
}

// 读取每一行asm源码, 转换为go源码,bin
func zParse(src string) (goout string, binout []byte) {
	srcLines := aPtrimCommentS(src)
	srcLines = ePtwoPseudo(srcLines)
	goout = "package\n"
	for k, v := range srcLines {
		outgo, outbin := iPtranslate(k, v)
		goout += outgo
		binout = append(binout, outbin...)
	}
	return goout, binout
}

func zPfile(in, out string) {
	f, err := ioutil.ReadFile(in)
	if err != nil {
		log.Println(err)
		return
	}
	outgo, outbin := zParse(string(f))
	ioutil.WriteFile(out+"go", []byte(outgo), 0777)
	ioutil.WriteFile(out+"bin", outbin, 0777)
}

func zPtestA() {
	f, _ := ioutil.ReadFile("test.asm")
	srcLines := aPtrimCommentS(string(f))
	dPtable(srcLines)
	for k, v := range mapLable {
		fmt.Println(k, v)
	}
	srcLines = ePtwoPseudo(srcLines)
	fmt.Println()
	fPtalbe(srcLines)
	for k, v := range mapLable {
		fmt.Println(k, v)
	}
	fmt.Println()
	for k, v := range srcLines {
		fmt.Println(k, v)
		fmt.Println(iPtranslate(k, v))

	}
}
func zPtestB() {
	binfile := []byte{}
	v := addi(0, 0, 0)
	bin := mPbigEnd(v)
	binfile = append(binfile, bin...)
	bin = mPbigEnd(v)
	binfile = append(binfile, bin...)
	for _, b := range binfile {
		fmt.Printf("%08b ", b)
	}
	ioutil.WriteFile("outBin", binfile, 0777)
	fmt.Println()
	f, _ := ioutil.ReadFile("outbin")
	for _, b := range f {
		fmt.Printf("%08b ", b)
	}
}
func zPtestC() {
	ins := "lr.d x5,(x20)"
	for k, v := range hPtrimBar(ins) {
		fmt.Println(k, v)
	}
	fmt.Println(hPtrimBar(ins), len(hPtrimBar(ins)))

}
