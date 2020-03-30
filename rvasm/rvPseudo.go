package rvasm

// 单行伪指令对应 指令
var mapPseudoRv = map[string]string{
	"fence":  "fence",
	"nop":    "addi",
	"beqz":   "beq",
	"bnez":   "bne",
	"bgez":   "bge",
	"bltz":   "blt",
	"neg":    "sub",
	"negw":   "subw",
	"snez":   "sltu",
	"sgtz":   "slt",
	"csrr":   "csrrs",
	"sltz":   "slt",
	"sext.w": "addiw",
	"mv":     "addi",
	"not":    "xori",
	"seqz":   "sltiu",
	"fmvs":   "fsgnjs",
	"fabss":  "fsgnjxs",
	"fnegs":  "fsgnjns",
	"fmvd":   "fsgnjd",
	"fabsd":  "fsgnjxd",
	"fnegd":  "fsgnjnd",
	"bgtz":   "blt",
	"blez":   "bge",
	"csrw":   "csrrw",
	"csrs":   "csrrs",
	"csrc":   "csrrc",
	"csrwi":  "csrrwi",
	"csrsi":  "csrrsi",
	"csrci":  "csrrci",
	"bgt":    "blt",
	"ble":    "bge",
	"bgtu":   "bltu",
	"bleu":   "bgeu",
	"j":      "jal",
	"jal":    "jal",
	"jalr":   "jalr",
	"jr":     "jalr",
	"ret":    "jalr",
	// "rdinstret[h]": "csrrs",
	// "rdcycle[h]":   "csrrs",
	// "rdtime[h]":    "csrrs",
	// "frcsr":        "csrrs",
	// "fscsr":        "csrrw",
	// "fscsr":        "csrrw",
	// "frrm":         "csrrs",
	// "fsrm":         "csrrw",
	// "fsrm":         "csrrw",
	// "frflags":      "csrrs",
	// "fsflags":      "csrrw",
	// "fsflags":      "csrrw",
}

// 单行伪指令 形式参数
var mapPseudo = map[string][]string{
	"fence":  []string{},                     //	fence iorw,iorw
	"nop":    []string{},                     //	addi x0,x0,0
	"beqz":   []string{"rs", "offset"},       //	beq rs,x0,offset
	"bnez":   []string{"rs", "offset"},       //	bne rs,x0,offset
	"bgez":   []string{"rs", "offset"},       //	bge rs,x0,offset
	"bltz":   []string{"rs", "offset"},       //	blt rs,x0,offset
	"neg":    []string{"rd", "rs"},           //	sub rd,x0,rs
	"negw":   []string{"rd", "rs"},           //	subw rd,x0,rs
	"snez":   []string{"rd", "rs"},           //	sltu rd,x0,rs
	"sgtz":   []string{"rd", "rs"},           //	slt rd,x0,rs
	"csrr":   []string{"rd", "csr"},          //	csrrs rd,csr,x0
	"sltz":   []string{"rd", "rs"},           //	slt rd,rs,x0
	"sext.w": []string{"rd", "rs"},           //	addiw rd,rs,0
	"mv":     []string{"rd", "rs"},           //	addi rd,rs,0
	"not":    []string{"rd", "rs"},           //	xori rd,rs,-1
	"seqz":   []string{"rd", "rs"},           //	sltiu rd,rs,1
	"fmv.s":  []string{"rd", "rs"},           //	fsgnj.s rd,rs,rs
	"fabs.s": []string{"rd", "rs"},           //	fsgnjx.s rd,rs,rs
	"fneg.s": []string{"rd", "rs"},           //	fsgnjn.s rd,rs,rs
	"fmv.d":  []string{"rd", "rs"},           //	fsgnj.d rd,rs,rs
	"fabs.d": []string{"rd", "rs"},           //	fsgnjx.d rd,rs,rs
	"fneg.d": []string{"rd", "rs"},           //	fsgnjn.d rd,rs,rs
	"bgtz":   []string{"rs", "offset"},       //	blt x0,rs,offset
	"blez":   []string{"rs", "offset"},       //	bge x0,rs,offset
	"jal":    []string{"offset"},             //	jal x1,offset
	"j":      []string{"offset"},             //	jal x0,offset
	"csrw":   []string{"csr", "rs"},          //	csrrw x0,csr,rs
	"csrs":   []string{"csr", "rs"},          //	csrrs x0,csr,rs
	"csrc":   []string{"csr", "rs"},          //	csrrc x0,csr,rs
	"csrwi":  []string{"csr", "imm"},         //	csrrwi x0,csr,imm
	"csrsi":  []string{"csr", "imm"},         //	csrrsi x0,csr,imm
	"csrci":  []string{"csr", "imm"},         //	csrrci x0,csr,imm
	"bgt":    []string{"rs", "rt", "offset"}, //	blt rt,rs,offset
	"ble":    []string{"rs", "rt", "offset"}, //	bge rt,rs,offset
	"bgtu":   []string{"rs", "rt", "offset"}, //	bltu rt,rs,offset
	"bleu":   []string{"rs", "rt", "offset"}, //	bgeu rt,rs,offset
	"jalr":   []string{"rs"},                 //	jalr x1,0(rs)
	"jr":     []string{"rs"},                 //	jalr x0,0(rs)
	"ret":    []string{},                     //	jalr x0,0(x1)
	// "rdinstret[h]": []string{"rd"},                 //	csrrs rd,instret[h],x0
	// "rdcycle[h]":   []string{"rd"},                 //	csrrs rd,cycle[h],x0
	// "rdtime[h]":    []string{"rd"},                 //	csrrs rd,time[h],x0
	// "frcsr":        []string{"rd"},                 //	csrrs rd,fcsr,x0
	// "fscsr":        []string{"rd", "rs"},           //	csrrw rd,fcsr,rs
	// "fscsr":        []string{"rs"},                 //	csrrw x0,fcsr,rs
	// "frrm":         []string{"rd"},                 //	csrrs rd,frm,x0
	// "fsrm":         []string{"rd", "rs"},           //	csrrw rd,frm,rs
	// "fsrm":         []string{"rs"},                 //	csrrw x0,frm,rs
	// "frflags":      []string{"rd"},                 //	csrrs rd,fflags,x0
	// "fsflags":      []string{"rd", "rs"},           //	csrrw rd,fflags,rs
	// "fsflags":      []string{"rs"},                 //	csrrw x0,fflags,rs

}

// 指令 形式参数
var mapRv = map[string][]string{
	"lui":      []string{"rd", "imm"},
	"auipc":    []string{"rd", "imm"},
	"jal":      []string{"rd", "imm"},
	"jalr":     []string{"rd", "rs1", "imm"},
	"lb":       []string{"rd", "rs1", "imm"},
	"lh":       []string{"rd", "rs1", "imm"},
	"lw":       []string{"rd", "rs1", "imm"},
	"lbu":      []string{"rd", "rs1", "imm"},
	"lhu":      []string{"rd", "rs1", "imm"},
	"addi":     []string{"rd", "rs1", "imm"},
	"slti":     []string{"rd", "rs1", "imm"},
	"sltiu":    []string{"rd", "rs1", "imm"},
	"xori":     []string{"rd", "rs1", "imm"},
	"ori":      []string{"rd", "rs1", "imm"},
	"andi":     []string{"rd", "rs1", "imm"},
	"lwu":      []string{"rd", "rs1", "imm"},
	"ld":       []string{"rd", "rs1", "imm"},
	"addiw":    []string{"rd", "rs1", "imm"},
	"flw":      []string{"rd", "rs1", "imm"},
	"fld":      []string{"rd", "rs1", "imm"},
	"flq":      []string{"rd", "rs1", "imm"},
	"fencei":   []string{"rd", "rs1", "imm"},
	"csrrw":    []string{"rd", "csr", "rs1"},
	"csrrs":    []string{"rd", "csr", "rs1"},
	"csrrc":    []string{"rd", "csr", "rs1"},
	"csrrwi":   []string{"rd", "csr", "uimm"},
	"csrrsi":   []string{"rd", "csr", "uimm"},
	"csrrci":   []string{"rd", "csr", "uimm"},
	"fmadds":   []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fmaddd":   []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fmaddq":   []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fmsubs":   []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fmsubd":   []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fmsubq":   []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fnmsubs":  []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fnmsubd":  []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fnmsubq":  []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fnmadds":  []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fnmaddd":  []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"fnmaddq":  []string{"rd", "rm", "rs1", "rs2", "rs3"},
	"amoaddw":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoswapw": []string{"rd", "rs1", "rs2", "rl", "aq"},
	"lrw":      []string{"rd", "rs1", "rl", "aq"},
	"scw":      []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoxorw":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoorw":   []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoandw":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amominw":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amomaxw":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amominuw": []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amomaxuw": []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoaddd":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoswapd": []string{"rd", "rs1", "rs2", "rl", "aq"},
	"lrd":      []string{"rd", "rs1", "rl", "aq"},
	"scd":      []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoxord":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoord":   []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amoandd":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amomind":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amomaxd":  []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amominud": []string{"rd", "rs1", "rs2", "rl", "aq"},
	"amomaxud": []string{"rd", "rs1", "rs2", "rl", "aq"},
	"slli":     []string{"rd", "rs1", "shamt"},
	"srli":     []string{"rd", "rs1", "shamt"},
	"srai":     []string{"rd", "rs1", "shamt"},
	"slliw":    []string{"rd", "rs1", "shamt"},
	"srliw":    []string{"rd", "rs1", "shamt"},
	"sraiw":    []string{"rd", "rs1", "shamt"},
	"add":      []string{"rd", "rs1", "rs2"},
	"sub":      []string{"rd", "rs1", "rs2"},
	"sll":      []string{"rd", "rs1", "rs2"},
	"slt":      []string{"rd", "rs1", "rs2"},
	"sltu":     []string{"rd", "rs1", "rs2"},
	"xor":      []string{"rd", "rs1", "rs2"},
	"srl":      []string{"rd", "rs1", "rs2"},
	"sra":      []string{"rd", "rs1", "rs2"},
	"or":       []string{"rd", "rs1", "rs2"},
	"and":      []string{"rd", "rs1", "rs2"},
	"mul":      []string{"rd", "rs1", "rs2"},
	"mulh":     []string{"rd", "rs1", "rs2"},
	"mulhsu":   []string{"rd", "rs1", "rs2"},
	"mulhu":    []string{"rd", "rs1", "rs2"},
	"div":      []string{"rd", "rs1", "rs2"},
	"divu":     []string{"rd", "rs1", "rs2"},
	"rem":      []string{"rd", "rs1", "rs2"},
	"remu":     []string{"rd", "rs1", "rs2"},
	"addw":     []string{"rd", "rs1", "rs2"},
	"subw":     []string{"rd", "rs1", "rs2"},
	"sllw":     []string{"rd", "rs1", "rs2"},
	"srlw":     []string{"rd", "rs1", "rs2"},
	"sraw":     []string{"rd", "rs1", "rs2"},
	"mulw":     []string{"rd", "rs1", "rs2"},
	"divw":     []string{"rd", "rs1", "rs2"},
	"divuw":    []string{"rd", "rs1", "rs2"},
	"remw":     []string{"rd", "rs1", "rs2"},
	"remuw":    []string{"rd", "rs1", "rs2"},
	"fles":     []string{"rd", "rs1", "rs2"},
	"fled":     []string{"rd", "rs1", "rs2"},
	"fleq":     []string{"rd", "rs1", "rs2"},
	"flts":     []string{"rd", "rs1", "rs2"},
	"fltd":     []string{"rd", "rs1", "rs2"},
	"fltq":     []string{"rd", "rs1", "rs2"},
	"feqs":     []string{"rd", "rs1", "rs2"},
	"feqd":     []string{"rd", "rs1", "rs2"},
	"feqq":     []string{"rd", "rs1", "rs2"},
	"fsgnjs":   []string{"rd", "rs1", "rs2"},
	"fsgnjd":   []string{"rd", "rs1", "rs2"},
	"fsgnjq":   []string{"rd", "rs1", "rs2"},
	"fsgnjns":  []string{"rd", "rs1", "rs2"},
	"fsgnjnd":  []string{"rd", "rs1", "rs2"},
	"fsgnjnq":  []string{"rd", "rs1", "rs2"},
	"fsgnjxs":  []string{"rd", "rs1", "rs2"},
	"fsgnjxd":  []string{"rd", "rs1", "rs2"},
	"fsgnjxq":  []string{"rd", "rs1", "rs2"},
	"fmins":    []string{"rd", "rs1", "rs2"},
	"fmind":    []string{"rd", "rs1", "rs2"},
	"fminq":    []string{"rd", "rs1", "rs2"},
	"fmaxs":    []string{"rd", "rs1", "rs2"},
	"fmaxd":    []string{"rd", "rs1", "rs2"},
	"fmaxq":    []string{"rd", "rs1", "rs2"},
	"fadds":    []string{"rd", "rm", "rs1", "rs2"},
	"faddd":    []string{"rd", "rm", "rs1", "rs2"},
	"faddq":    []string{"rd", "rm", "rs1", "rs2"},
	"fsubs":    []string{"rd", "rm", "rs1", "rs2"},
	"fsubd":    []string{"rd", "rm", "rs1", "rs2"},
	"fsubq":    []string{"rd", "rm", "rs1", "rs2"},
	"fmuls":    []string{"rd", "rm", "rs1", "rs2"},
	"fmuld":    []string{"rd", "rm", "rs1", "rs2"},
	"fmulq":    []string{"rd", "rm", "rs1", "rs2"},
	"fdivs":    []string{"rd", "rm", "rs1", "rs2"},
	"fdivd":    []string{"rd", "rm", "rs1", "rs2"},
	"fdivq":    []string{"rd", "rm", "rs1", "rs2"},
	"fcvtsq":   []string{"rd", "rm", "rs1"},
	"fcvtsd":   []string{"rd", "rm", "rs1"},
	"fcvtds":   []string{"rd", "rm", "rs1"},
	"fcvtdq":   []string{"rd", "rm", "rs1"},
	"fcvtqs":   []string{"rd", "rm", "rs1"},
	"fcvtqd":   []string{"rd", "rm", "rs1"},
	"fsqrts":   []string{"rd", "rm", "rs1"},
	"fsqrtd":   []string{"rd", "rm", "rs1"},
	"fsqrtq":   []string{"rd", "rm", "rs1"},
	"fcvtws":   []string{"rd", "rm", "rs1"},
	"fcvtwus":  []string{"rd", "rm", "rs1"},
	"fcvtls":   []string{"rd", "rm", "rs1"},
	"fcvtlus":  []string{"rd", "rm", "rs1"},
	"fcvtwd":   []string{"rd", "rm", "rs1"},
	"fcvtwud":  []string{"rd", "rm", "rs1"},
	"fcvtld":   []string{"rd", "rm", "rs1"},
	"fcvtlud":  []string{"rd", "rm", "rs1"},
	"fcvtwq":   []string{"rd", "rm", "rs1"},
	"fcvtwuq":  []string{"rd", "rm", "rs1"},
	"fcvtlq":   []string{"rd", "rm", "rs1"},
	"fcvtluq":  []string{"rd", "rm", "rs1"},
	"fcvtsw":   []string{"rd", "rm", "rs1"},
	"fcvtswu":  []string{"rd", "rm", "rs1"},
	"fcvtsl":   []string{"rd", "rm", "rs1"},
	"fcvtslu":  []string{"rd", "rm", "rs1"},
	"fcvtdw":   []string{"rd", "rm", "rs1"},
	"fcvtdwu":  []string{"rd", "rm", "rs1"},
	"fcvtdl":   []string{"rd", "rm", "rs1"},
	"fcvtdlu":  []string{"rd", "rm", "rs1"},
	"fcvtqw":   []string{"rd", "rm", "rs1"},
	"fcvtqwu":  []string{"rd", "rm", "rs1"},
	"fcvtql":   []string{"rd", "rm", "rs1"},
	"fcvtqlu":  []string{"rd", "rm", "rs1"},
	"fmvwx":    []string{"rd", "rs1"},
	"fmvxw":    []string{"rd", "rs1"},
	"fmvdx":    []string{"rd", "rs1"},
	"fmvxd":    []string{"rd", "rs1"},
	"fclasss":  []string{"rd", "rs1"},
	"fclassd":  []string{"rd", "rs1"},
	"fclassq":  []string{"rd", "rs1"},
	"beq":      []string{"rs1", "rs2", "imm"},
	"bne":      []string{"rs1", "rs2", "imm"},
	"blt":      []string{"rs1", "rs2", "imm"},
	"bge":      []string{"rs1", "rs2", "imm"},
	"bltu":     []string{"rs1", "rs2", "imm"},
	"bgeu":     []string{"rs1", "rs2", "imm"},
	"sb":       []string{"rs1", "rs2", "imm"},
	"sh":       []string{"rs1", "rs2", "imm"},
	"sw":       []string{"rs1", "rs2", "imm"},
	"sd":       []string{"rs1", "rs2", "imm"},
	"fsw":      []string{"rs1", "rs2", "imm"},
	"fsd":      []string{"rs1", "rs2", "imm"},
	"fsq":      []string{"rs1", "rs2", "imm"},
	"ebreak":   []string{},
	"ecall":    []string{},
	"fence":    []string{"rd", "rm", "rs1", "pred", "succ"},
}
