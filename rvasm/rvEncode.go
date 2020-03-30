package rvasm

import (
	"fmt"
)

func bEncode(fn string, args []uint32) (code uint32) {
	switch fn {
	case "add":
		code = add(args[0], args[1], args[2])
	case "addi":
		code = addi(args[0], args[1], args[2])
	case "addiw":
		code = addiw(args[0], args[1], args[2])
	case "addw":
		code = addw(args[0], args[1], args[2])
	case "amoaddd":
		code = amoaddd(args[0], args[1], args[2], args[3], args[4])
	case "amoaddw":
		code = amoaddw(args[0], args[1], args[2], args[3], args[4])
	case "amoandd":
		code = amoandd(args[0], args[1], args[2], args[3], args[4])
	case "amoandw":
		code = amoandw(args[0], args[1], args[2], args[3], args[4])
	case "amomaxd":
		code = amomaxd(args[0], args[1], args[2], args[3], args[4])
	case "amomaxud":
		code = amomaxud(args[0], args[1], args[2], args[3], args[4])
	case "amomaxuw":
		code = amomaxuw(args[0], args[1], args[2], args[3], args[4])
	case "amomaxw":
		code = amomaxw(args[0], args[1], args[2], args[3], args[4])
	case "amomind":
		code = amomind(args[0], args[1], args[2], args[3], args[4])
	case "amominud":
		code = amominud(args[0], args[1], args[2], args[3], args[4])
	case "amominuw":
		code = amominuw(args[0], args[1], args[2], args[3], args[4])
	case "amominw":
		code = amominw(args[0], args[1], args[2], args[3], args[4])
	case "amoord":
		code = amoord(args[0], args[1], args[2], args[3], args[4])
	case "amoorw":
		code = amoorw(args[0], args[1], args[2], args[3], args[4])
	case "amoswapd":
		code = amoswapd(args[0], args[1], args[2], args[3], args[4])
	case "amoswapw":
		code = amoswapw(args[0], args[1], args[2], args[3], args[4])
	case "amoxord":
		code = amoxord(args[0], args[1], args[2], args[3], args[4])
	case "amoxorw":
		code = amoxorw(args[0], args[1], args[2], args[3], args[4])
	case "and":
		code = and(args[0], args[1], args[2])
	case "andi":
		code = andi(args[0], args[1], args[2])
	case "auipc":
		code = auipc(args[0], args[1])
	case "beq":
		code = beq(args[0], args[1], args[2])
	case "bge":
		code = bge(args[0], args[1], args[2])
	case "bgeu":
		code = bgeu(args[0], args[1], args[2])
	case "blt":
		code = blt(args[0], args[1], args[2])
	case "bltu":
		code = bltu(args[0], args[1], args[2])
	case "bne":
		code = bne(args[0], args[1], args[2])
	case "csrrc":
		code = csrrc(args[0], args[1], args[2])
	case "csrrci":
		code = csrrci(args[0], args[1], args[2])
	case "csrrs":
		code = csrrs(args[0], args[1], args[2])
	case "csrrsi":
		code = csrrsi(args[0], args[1], args[2])
	case "csrrw":
		code = csrrw(args[0], args[1], args[2])
	case "csrrwi":
		code = csrrwi(args[0], args[1], args[2])
	case "div":
		code = div(args[0], args[1], args[2])
	case "divu":
		code = divu(args[0], args[1], args[2])
	case "divuw":
		code = divuw(args[0], args[1], args[2])
	case "divw":
		code = divw(args[0], args[1], args[2])
	case "ebreak":
		code = ebreak()
	case "ecall":
		code = ecall()
	case "faddd":
		code = faddd(args[0], args[1], args[2], args[3])
	case "faddq":
		code = faddq(args[0], args[1], args[2], args[3])
	case "fadds":
		code = fadds(args[0], args[1], args[2], args[3])
	case "fclassd":
		code = fclassd(args[0], args[1])
	case "fclassq":
		code = fclassq(args[0], args[1])
	case "fclasss":
		code = fclasss(args[0], args[1])
	case "fcvtdl":
		code = fcvtdl(args[0], args[1], args[2])
	case "fcvtdlu":
		code = fcvtdlu(args[0], args[1], args[2])
	case "fcvtdq":
		code = fcvtdq(args[0], args[1], args[2])
	case "fcvtds":
		code = fcvtds(args[0], args[1], args[2])
	case "fcvtdw":
		code = fcvtdw(args[0], args[1], args[2])
	case "fcvtdwu":
		code = fcvtdwu(args[0], args[1], args[2])
	case "fcvtld":
		code = fcvtld(args[0], args[1], args[2])
	case "fcvtlq":
		code = fcvtlq(args[0], args[1], args[2])
	case "fcvtls":
		code = fcvtls(args[0], args[1], args[2])
	case "fcvtlud":
		code = fcvtlud(args[0], args[1], args[2])
	case "fcvtluq":
		code = fcvtluq(args[0], args[1], args[2])
	case "fcvtlus":
		code = fcvtlus(args[0], args[1], args[2])
	case "fcvtqd":
		code = fcvtqd(args[0], args[1], args[2])
	case "fcvtql":
		code = fcvtql(args[0], args[1], args[2])
	case "fcvtqlu":
		code = fcvtqlu(args[0], args[1], args[2])
	case "fcvtqs":
		code = fcvtqs(args[0], args[1], args[2])
	case "fcvtqw":
		code = fcvtqw(args[0], args[1], args[2])
	case "fcvtqwu":
		code = fcvtqwu(args[0], args[1], args[2])
	case "fcvtsd":
		code = fcvtsd(args[0], args[1], args[2])
	case "fcvtsl":
		code = fcvtsl(args[0], args[1], args[2])
	case "fcvtslu":
		code = fcvtslu(args[0], args[1], args[2])
	case "fcvtsq":
		code = fcvtsq(args[0], args[1], args[2])
	case "fcvtsw":
		code = fcvtsw(args[0], args[1], args[2])
	case "fcvtswu":
		code = fcvtswu(args[0], args[1], args[2])
	case "fcvtwd":
		code = fcvtwd(args[0], args[1], args[2])
	case "fcvtwq":
		code = fcvtwq(args[0], args[1], args[2])
	case "fcvtws":
		code = fcvtws(args[0], args[1], args[2])
	case "fcvtwud":
		code = fcvtwud(args[0], args[1], args[2])
	case "fcvtwuq":
		code = fcvtwuq(args[0], args[1], args[2])
	case "fcvtwus":
		code = fcvtwus(args[0], args[1], args[2])
	case "fdivd":
		code = fdivd(args[0], args[1], args[2], args[3])
	case "fdivq":
		code = fdivq(args[0], args[1], args[2], args[3])
	case "fdivs":
		code = fdivs(args[0], args[1], args[2], args[3])
	case "fence":
		code = fence(args[0], args[1])
	case "fencei":
		code = fencei(args[0], args[1], args[2])
	case "feqd":
		code = feqd(args[0], args[1], args[2])
	case "feqq":
		code = feqq(args[0], args[1], args[2])
	case "feqs":
		code = feqs(args[0], args[1], args[2])
	case "fld":
		code = fld(args[0], args[1], args[2])
	case "fled":
		code = fled(args[0], args[1], args[2])
	case "fleq":
		code = fleq(args[0], args[1], args[2])
	case "fles":
		code = fles(args[0], args[1], args[2])
	case "flq":
		code = flq(args[0], args[1], args[2])
	case "fltd":
		code = fltd(args[0], args[1], args[2])
	case "fltq":
		code = fltq(args[0], args[1], args[2])
	case "flts":
		code = flts(args[0], args[1], args[2])
	case "flw":
		code = flw(args[0], args[1], args[2])
	case "fmaddd":
		code = fmaddd(args[0], args[1], args[2], args[3], args[4])
	case "fmaddq":
		code = fmaddq(args[0], args[1], args[2], args[3], args[4])
	case "fmadds":
		code = fmadds(args[0], args[1], args[2], args[3], args[4])
	case "fmaxd":
		code = fmaxd(args[0], args[1], args[2])
	case "fmaxq":
		code = fmaxq(args[0], args[1], args[2])
	case "fmaxs":
		code = fmaxs(args[0], args[1], args[2])
	case "fmind":
		code = fmind(args[0], args[1], args[2])
	case "fminq":
		code = fminq(args[0], args[1], args[2])
	case "fmins":
		code = fmins(args[0], args[1], args[2])
	case "fmsubd":
		code = fmsubd(args[0], args[1], args[2], args[3], args[4])
	case "fmsubq":
		code = fmsubq(args[0], args[1], args[2], args[3], args[4])
	case "fmsubs":
		code = fmsubs(args[0], args[1], args[2], args[3], args[4])
	case "fmuld":
		code = fmuld(args[0], args[1], args[2], args[3])
	case "fmulq":
		code = fmulq(args[0], args[1], args[2], args[3])
	case "fmuls":
		code = fmuls(args[0], args[1], args[2], args[3])
	case "fmvdx":
		code = fmvdx(args[0], args[1])
	case "fmvwx":
		code = fmvwx(args[0], args[1])
	case "fmvxd":
		code = fmvxd(args[0], args[1])
	case "fmvxw":
		code = fmvxw(args[0], args[1])
	case "fnmaddd":
		code = fnmaddd(args[0], args[1], args[2], args[3], args[4])
	case "fnmaddq":
		code = fnmaddq(args[0], args[1], args[2], args[3], args[4])
	case "fnmadds":
		code = fnmadds(args[0], args[1], args[2], args[3], args[4])
	case "fnmsubd":
		code = fnmsubd(args[0], args[1], args[2], args[3], args[4])
	case "fnmsubq":
		code = fnmsubq(args[0], args[1], args[2], args[3], args[4])
	case "fnmsubs":
		code = fnmsubs(args[0], args[1], args[2], args[3], args[4])
	case "fsd":
		code = fsd(args[0], args[1], args[2])
	case "fsgnjd":
		code = fsgnjd(args[0], args[1], args[2])
	case "fsgnjnd":
		code = fsgnjnd(args[0], args[1], args[2])
	case "fsgnjnq":
		code = fsgnjnq(args[0], args[1], args[2])
	case "fsgnjns":
		code = fsgnjns(args[0], args[1], args[2])
	case "fsgnjq":
		code = fsgnjq(args[0], args[1], args[2])
	case "fsgnjs":
		code = fsgnjs(args[0], args[1], args[2])
	case "fsgnjxd":
		code = fsgnjxd(args[0], args[1], args[2])
	case "fsgnjxq":
		code = fsgnjxq(args[0], args[1], args[2])
	case "fsgnjxs":
		code = fsgnjxs(args[0], args[1], args[2])
	case "fsq":
		code = fsq(args[0], args[1], args[2])
	case "fsqrtd":
		code = fsqrtd(args[0], args[1], args[2])
	case "fsqrtq":
		code = fsqrtq(args[0], args[1], args[2])
	case "fsqrts":
		code = fsqrts(args[0], args[1], args[2])
	case "fsubd":
		code = fsubd(args[0], args[1], args[2], args[3])
	case "fsubq":
		code = fsubq(args[0], args[1], args[2], args[3])
	case "fsubs":
		code = fsubs(args[0], args[1], args[2], args[3])
	case "fsw":
		code = fsw(args[0], args[1], args[2])
	case "jal":
		code = jal(args[0], args[1])
	case "jalr":
		code = jalr(args[0], args[1], args[2])
	case "lb":
		code = lb(args[0], args[1], args[2])
	case "lbu":
		code = lbu(args[0], args[1], args[2])
	case "ld":
		code = ld(args[0], args[1], args[2])
	case "lh":
		code = lh(args[0], args[1], args[2])
	case "lhu":
		code = lhu(args[0], args[1], args[2])
	case "lrd":
		code = lrd(args[0], args[1], args[2], args[3])
	case "lrw":
		code = lrw(args[0], args[1], args[2], args[3])
	case "lui":
		code = lui(args[0], args[1])
	case "lw":
		code = lw(args[0], args[1], args[2])
	case "lwu":
		code = lwu(args[0], args[1], args[2])
	case "mul":
		code = mul(args[0], args[1], args[2])
	case "mulh":
		code = mulh(args[0], args[1], args[2])
	case "mulhsu":
		code = mulhsu(args[0], args[1], args[2])
	case "mulhu":
		code = mulhu(args[0], args[1], args[2])
	case "mulw":
		code = mulw(args[0], args[1], args[2])
	case "or":
		code = or(args[0], args[1], args[2])
	case "ori":
		code = ori(args[0], args[1], args[2])
	case "rem":
		code = rem(args[0], args[1], args[2])
	case "remu":
		code = remu(args[0], args[1], args[2])
	case "remuw":
		code = remuw(args[0], args[1], args[2])
	case "remw":
		code = remw(args[0], args[1], args[2])
	case "sb":
		code = sb(args[0], args[1], args[2])
	case "scd":
		code = scd(args[0], args[1], args[2], args[3], args[4])
	case "scw":
		code = scw(args[0], args[1], args[2], args[3], args[4])
	case "sd":
		code = sd(args[0], args[1], args[2])
	case "sh":
		code = sh(args[0], args[1], args[2])
	case "sll":
		code = sll(args[0], args[1], args[2])
	case "slli":
		code = slli(args[0], args[1], args[2])
	case "slliw":
		code = slliw(args[0], args[1], args[2])
	case "sllw":
		code = sllw(args[0], args[1], args[2])
	case "slt":
		code = slt(args[0], args[1], args[2])
	case "slti":
		code = slti(args[0], args[1], args[2])
	case "sltiu":
		code = sltiu(args[0], args[1], args[2])
	case "sltu":
		code = sltu(args[0], args[1], args[2])
	case "sra":
		code = sra(args[0], args[1], args[2])
	case "srai":
		code = srai(args[0], args[1], args[2])
	case "sraiw":
		code = sraiw(args[0], args[1], args[2])
	case "sraw":
		code = sraw(args[0], args[1], args[2])
	case "srl":
		code = srl(args[0], args[1], args[2])
	case "srli":
		code = srli(args[0], args[1], args[2])
	case "srliw":
		code = srliw(args[0], args[1], args[2])
	case "srlw":
		code = srlw(args[0], args[1], args[2])
	case "sub":
		code = sub(args[0], args[1], args[2])
	case "subw":
		code = subw(args[0], args[1], args[2])
	case "sw":
		code = sw(args[0], args[1], args[2])
	case "xor":
		code = xor(args[0], args[1], args[2])
	case "xori":
		code = xori(args[0], args[1], args[2])
	default:
		fmt.Printf("bEncode......no %s", fn)
	}
	return code
}

var cErvMap = map[string][]string{
	"add":      []string{"rd", "rs1", "rs2"},
	"addi":     []string{"rd", "rs1", "imm"},
	"addiw":    []string{"rd", "rs1", "imm"},
	"addw":     []string{"rd", "rs1", "rs2"},
	"amoaddd":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoaddw":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoandd":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoandw":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amomaxd":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amomaxud": []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amomaxuw": []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amomaxw":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amomind":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amominud": []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amominuw": []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amominw":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoord":   []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoorw":   []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoswapd": []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoswapw": []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoxord":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoxorw":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"and":      []string{"rd", "rs1", "rs2"},
	"andi":     []string{"rd", "rs1", "imm"},
	"auipc":    []string{"rd", "imm"},
	"beq":      []string{"rs1", "rs2", "imm"},
	"bge":      []string{"rs1", "rs2", "imm"},
	"bgeu":     []string{"rs1", "rs2", "imm"},
	"blt":      []string{"rs1", "rs2", "imm"},
	"bltu":     []string{"rs1", "rs2", "imm"},
	"bne":      []string{"rs1", "rs2", "imm"},
	"csrrc":    []string{"rd", "csr", "rs1"},
	"csrrci":   []string{"rd", "csr", "uimm"},
	"csrrs":    []string{"rd", "csr", "rs1"},
	"csrrsi":   []string{"rd", "csr", "uimm"},
	"csrrw":    []string{"rd", "csr", "rs1"},
	"csrrwi":   []string{"rd", "csr", "uimm"},
	"div":      []string{"rd", "rs1", "rs2"},
	"divu":     []string{"rd", "rs1", "rs2"},
	"divuw":    []string{"rd", "rs1", "rs2"},
	"divw":     []string{"rd", "rs1", "rs2"},
	"ebreak":   []string{},
	"ecall":    []string{},
	"faddd":    []string{"rd", "rm", "rs1", "rs2"},
	"faddq":    []string{"rd", "rm", "rs1", "rs2"},
	"fadds":    []string{"rd", "rm", "rs1", "rs2"},
	"fclassd":  []string{"rd", "rs1"},
	"fclassq":  []string{"rd", "rs1"},
	"fclasss":  []string{"rd", "rs1"},
	"fcvtdl":   []string{"rd", "rm", "rs1"},
	"fcvtdlu":  []string{"rd", "rm", "rs1"},
	"fcvtdq":   []string{"rd", "rm", "rs1"},
	"fcvtds":   []string{"rd", "rm", "rs1"},
	"fcvtdw":   []string{"rd", "rm", "rs1"},
	"fcvtdwu":  []string{"rd", "rm", "rs1"},
	"fcvtld":   []string{"rd", "rm", "rs1"},
	"fcvtlq":   []string{"rd", "rm", "rs1"},
	"fcvtls":   []string{"rd", "rm", "rs1"},
	"fcvtlud":  []string{"rd", "rm", "rs1"},
	"fcvtluq":  []string{"rd", "rm", "rs1"},
	"fcvtlus":  []string{"rd", "rm", "rs1"},
	"fcvtqd":   []string{"rd", "rm", "rs1"},
	"fcvtql":   []string{"rd", "rm", "rs1"},
	"fcvtqlu":  []string{"rd", "rm", "rs1"},
	"fcvtqs":   []string{"rd", "rm", "rs1"},
	"fcvtqw":   []string{"rd", "rm", "rs1"},
	"fcvtqwu":  []string{"rd", "rm", "rs1"},
	"fcvtsd":   []string{"rd", "rm", "rs1"},
	"fcvtsl":   []string{"rd", "rm", "rs1"},
	"fcvtslu":  []string{"rd", "rm", "rs1"},
	"fcvtsq":   []string{"rd", "rm", "rs1"},
	"fcvtsw":   []string{"rd", "rm", "rs1"},
	"fcvtswu":  []string{"rd", "rm", "rs1"},
	"fcvtwd":   []string{"rd", "rm", "rs1"},
	"fcvtwq":   []string{"rd", "rm", "rs1"},
	"fcvtws":   []string{"rd", "rm", "rs1"},
	"fcvtwud":  []string{"rd", "rm", "rs1"},
	"fcvtwuq":  []string{"rd", "rm", "rs1"},
	"fcvtwus":  []string{"rd", "rm", "rs1"},
	"fdivd":    []string{"rd", "rm", "rs1", "rs2"},
	"fdivq":    []string{"rd", "rm", "rs1", "rs2"},
	"fdivs":    []string{"rd", "rm", "rs1", "rs2"},
	"fence":    []string{"rd", "rs1"},
	"fencei":   []string{"rd", "rs1", "imm"},
	"feqd":     []string{"rd", "rs1", "rs2"},
	"feqq":     []string{"rd", "rs1", "rs2"},
	"feqs":     []string{"rd", "rs1", "rs2"},
	"fld":      []string{"rd", "rs1", "imm"},
	"fled":     []string{"rd", "rs1", "rs2"},
	"fleq":     []string{"rd", "rs1", "rs2"},
	"fles":     []string{"rd", "rs1", "rs2"},
	"flq":      []string{"rd", "rs1", "imm"},
	"fltd":     []string{"rd", "rs1", "rs2"},
	"fltq":     []string{"rd", "rs1", "rs2"},
	"flts":     []string{"rd", "rs1", "rs2"},
	"flw":      []string{"rd", "rs1", "imm"},
	"fmaddd":   []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fmaddq":   []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fmadds":   []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fmaxd":    []string{"rd", "rs1", "rs2"},
	"fmaxq":    []string{"rd", "rs1", "rs2"},
	"fmaxs":    []string{"rd", "rs1", "rs2"},
	"fmind":    []string{"rd", "rs1", "rs2"},
	"fminq":    []string{"rd", "rs1", "rs2"},
	"fmins":    []string{"rd", "rs1", "rs2"},
	"fmsubd":   []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fmsubq":   []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fmsubs":   []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fmuld":    []string{"rd", "rm", "rs1", "rs2"},
	"fmulq":    []string{"rd", "rm", "rs1", "rs2"},
	"fmuls":    []string{"rd", "rm", "rs1", "rs2"},
	"fmvdx":    []string{"rd", "rs1"},
	"fmvwx":    []string{"rd", "rs1"},
	"fmvxd":    []string{"rd", "rs1"},
	"fmvxw":    []string{"rd", "rs1"},
	"fnmaddd":  []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fnmaddq":  []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fnmadds":  []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fnmsubd":  []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fnmsubq":  []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fnmsubs":  []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fsd":      []string{"rs1", "rs2", "imm"},
	"fsgnjd":   []string{"rd", "rs1", "rs2"},
	"fsgnjnd":  []string{"rd", "rs1", "rs2"},
	"fsgnjnq":  []string{"rd", "rs1", "rs2"},
	"fsgnjns":  []string{"rd", "rs1", "rs2"},
	"fsgnjq":   []string{"rd", "rs1", "rs2"},
	"fsgnjs":   []string{"rd", "rs1", "rs2"},
	"fsgnjxd":  []string{"rd", "rs1", "rs2"},
	"fsgnjxq":  []string{"rd", "rs1", "rs2"},
	"fsgnjxs":  []string{"rd", "rs1", "rs2"},
	"fsq":      []string{"rs1", "rs2", "imm"},
	"fsqrtd":   []string{"rd", "rm", "rs1"},
	"fsqrtq":   []string{"rd", "rm", "rs1"},
	"fsqrts":   []string{"rd", "rm", "rs1"},
	"fsubd":    []string{"rd", "rm", "rs1", "rs2"},
	"fsubq":    []string{"rd", "rm", "rs1", "rs2"},
	"fsubs":    []string{"rd", "rm", "rs1", "rs2"},
	"fsw":      []string{"rs1", "rs2", "imm"},
	"jal":      []string{"rd", "imm"},
	"jalr":     []string{"rd", "rs1", "imm"},
	"lb":       []string{"rd", "rs1", "imm"},
	"lbu":      []string{"rd", "rs1", "imm"},
	"ld":       []string{"rd", "rs1", "imm"},
	"lh":       []string{"rd", "rs1", "imm"},
	"lhu":      []string{"rd", "rs1", "imm"},
	"lrd":      []string{"rd", "rs1", "rl", "aq"},
	"lrw":      []string{"rd", "rs1", "rl", "aq"},
	"lui":      []string{"rd", "imm"},
	"lw":       []string{"rd", "rs1", "imm"},
	"lwu":      []string{"rd", "rs1", "imm"},
	"mul":      []string{"rd", "rs1", "rs2"},
	"mulh":     []string{"rd", "rs1", "rs2"},
	"mulhsu":   []string{"rd", "rs1", "rs2"},
	"mulhu":    []string{"rd", "rs1", "rs2"},
	"mulw":     []string{"rd", "rs1", "rs2"},
	"or":       []string{"rd", "rs1", "rs2"},
	"ori":      []string{"rd", "rs1", "imm"},
	"rem":      []string{"rd", "rs1", "rs2"},
	"remu":     []string{"rd", "rs1", "rs2"},
	"remuw":    []string{"rd", "rs1", "rs2"},
	"remw":     []string{"rd", "rs1", "rs2"},
	"sb":       []string{"rs1", "rs2", "imm"},
	"scd":      []string{"rd", "rs1", "rs2", "rl", "aq"},
	"scw":      []string{"rd", "rs1", "rs2", "rl", "aq"},
	"sd":       []string{"rs1", "rs2", "imm"},
	"sh":       []string{"rs1", "rs2", "imm"},
	"sll":      []string{"rd", "rs1", "rs2"},
	"slli":     []string{"rd", "rs1", "shamt"},
	"slliw":    []string{"rd", "rs1", "shamt"},
	"sllw":     []string{"rd", "rs1", "rs2"},
	"slt":      []string{"rd", "rs1", "rs2"},
	"slti":     []string{"rd", "rs1", "imm"},
	"sltiu":    []string{"rd", "rs1", "imm"},
	"sltu":     []string{"rd", "rs1", "rs2"},
	"sra":      []string{"rd", "rs1", "rs2"},
	"srai":     []string{"rd", "rs1", "shamt"},
	"sraiw":    []string{"rd", "rs1", "shamt"},
	"sraw":     []string{"rd", "rs1", "rs2"},
	"srl":      []string{"rd", "rs1", "rs2"},
	"srli":     []string{"rd", "rs1", "shamt"},
	"srliw":    []string{"rd", "rs1", "shamt"},
	"srlw":     []string{"rd", "rs1", "rs2"},
	"sub":      []string{"rd", "rs1", "rs2"},
	"subw":     []string{"rd", "rs1", "rs2"},
	"sw":       []string{"rs1", "rs2", "imm"},
	"xor":      []string{"rd", "rs1", "rs2"},
	"xori":     []string{"rd", "rs1", "imm"},
}
