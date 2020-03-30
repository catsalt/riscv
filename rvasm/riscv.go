package rvasm

//0 FENCE: fm pred succ rs1 000 rd 0001111 FENCE
func fence(rd, rs1 uint32) (code uint32) {
	// code= fm<<28 | pred<<24 | succ<<20 | rs1<<15 | 0b000 | rd<<7 | 0b0001111
	return code
}

//1 ECALL: 000000000000 00000 000 00000 1110011 ECALL
func ecall() (code uint32) {
	code = 0b1110011
	return code
}

//2 EBREAK: 000000000001 00000 000 00000 1110011 EBREAK
func ebreak() (code uint32) {
	code = 0b100000000000001110011
	return code
}

//3 LUI: imm[31:12] rd 0110111 LUI
func lui(rd, imm uint32) (code uint32) {
	code = imm>>12<<12 | rd<<7 | 0b0110111
	return code
}

//4 AUIPC: imm[31:12] rd 0010111 AUIPC
func auipc(rd, imm uint32) (code uint32) {
	code = imm>>12<<12 | rd<<7 | 0b0010111
	return code
}

//5 JAL: imm[20|10:1|11|19:12] rd 1101111 JAL
func jal(rd, imm uint32) (code uint32) {
	code = (imm>>20&0b1<<19|imm>>1&(1<<10-1)<<9|imm>>11&0b1<<8|imm>>12&(1<<8-1))<<12 | rd<<7 | 0b1101111
	return code
}

//6 JALR: imm[11:0] rs1 000 rd 1100111 JALR
func jalr(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b1100111
	return code
}

//7 LB: imm[11:0] rs1 000 rd 0000011 LB
func lb(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b0000011
	return code
}

//8 LH: imm[11:0] rs1 001 rd 0000011 LH
func lh(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b0000011
	return code
}

//9 LW: imm[11:0] rs1 010 rd 0000011 LW
func lw(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b0000011
	return code
}

//10 LBU: imm[11:0] rs1 100 rd 0000011 LBU
func lbu(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b100<<12 | rd<<7 | 0b0000011
	return code
}

//11 LHU: imm[11:0] rs1 101 rd 0000011 LHU
func lhu(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b101<<12 | rd<<7 | 0b0000011
	return code
}

//12 ADDI: imm[11:0] rs1 000 rd 0010011 ADDI
func addi(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b0010011
	return code
}

//13 SLTI: imm[11:0] rs1 010 rd 0010011 SLTI
func slti(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b0010011
	return code
}

//14 SLTIU: imm[11:0] rs1 011 rd 0010011 SLTIU
func sltiu(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b011<<12 | rd<<7 | 0b0010011
	return code
}

//15 XORI: imm[11:0] rs1 100 rd 0010011 XORI
func xori(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b100<<12 | rd<<7 | 0b0010011
	return code
}

//16 ORI: imm[11:0] rs1 110 rd 0010011 ORI
func ori(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b110<<12 | rd<<7 | 0b0010011
	return code
}

//17 ANDI: imm[11:0] rs1 111 rd 0010011 ANDI
func andi(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b111<<12 | rd<<7 | 0b0010011
	return code
}

//18 LWU: imm[11:0] rs1 110 rd 0000011 LWU
func lwu(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b110<<12 | rd<<7 | 0b0000011
	return code
}

//19 LD: imm[11:0] rs1 011 rd 0000011 LD
func ld(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b011<<12 | rd<<7 | 0b0000011
	return code
}

//20 ADDIW: imm[11:0] rs1 000 rd 0011011 ADDIW
func addiw(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b0011011
	return code
}

//21 FLW: imm[11:0] rs1 010 rd 0000111 FLW
func flw(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b0000111
	return code
}

//22 FLD: imm[11:0] rs1 011 rd 0000111 FLD
func fld(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b011<<12 | rd<<7 | 0b0000111
	return code
}

//23 FLQ: imm[11:0] rs1 100 rd 0000111 FLQ
func flq(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b100<<12 | rd<<7 | 0b0000111
	return code
}

//24 FENCE.I: imm[11:0] rs1 001 rd 0001111 FENCE.I
func fencei(rd, rs1, imm uint32) (code uint32) {
	code = imm&(1<<12-1)<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b0001111
	return code
}

//25 BEQ: imm[12|10:5] rs2 rs1 000 imm[4:1|11] 1100011 BEQ
func beq(rs1, rs2, imm uint32) (code uint32) {
	code = (imm>>12&0b1<<6|imm>>5&(1<<6-1))<<25 | rs2<<20 | rs1<<15 | 0b000<<12 | (imm&0b11110|imm>>11&0b1)<<7 | 0b1100011
	return code
}

//26 BNE: imm[12|10:5] rs2 rs1 001 imm[4:1|11] 1100011 BNE
func bne(rs1, rs2, imm uint32) (code uint32) {
	code = (imm>>12&0b1<<6|imm>>5&(1<<6-1))<<25 | rs2<<20 | rs1<<15 | 0b001<<12 | (imm&0b11110|imm>>11&0b1)<<7 | 0b1100011
	return code
}

//27 BLT: imm[12|10:5] rs2 rs1 100 imm[4:1|11] 1100011 BLT
func blt(rs1, rs2, imm uint32) (code uint32) {
	code = (imm>>12&0b1<<6|imm>>5&(1<<6-1))<<25 | rs2<<20 | rs1<<15 | 0b100<<12 | (imm&0b11110|imm>>11&0b1)<<7 | 0b1100011
	return code
}

//28 BGE: imm[12|10:5] rs2 rs1 101 imm[4:1|11] 1100011 BGE
func bge(rs1, rs2, imm uint32) (code uint32) {
	code = (imm>>12&0b1<<6|imm>>5&(1<<6-1))<<25 | rs2<<20 | rs1<<15 | 0b101<<12 | (imm&0b11110|imm>>11&0b1)<<7 | 0b1100011
	return code
}

//29 BLTU: imm[12|10:5] rs2 rs1 110 imm[4:1|11] 1100011 BLTU
func bltu(rs1, rs2, imm uint32) (code uint32) {
	code = (imm>>12&0b1<<6|imm>>5&(1<<6-1))<<25 | rs2<<20 | rs1<<15 | 0b110<<12 | (imm&0b11110|imm>>11&0b1)<<7 | 0b1100011
	return code
}

//30 BGEU: imm[12|10:5] rs2 rs1 111 imm[4:1|11] 1100011 BGEU
func bgeu(rs1, rs2, imm uint32) (code uint32) {
	code = (imm>>12&0b1<<6|imm>>5&(1<<6-1))<<25 | rs2<<20 | rs1<<15 | 0b111<<12 | (imm&0b11110|imm>>11&0b1)<<7 | 0b1100011
	return code
}

//31 SB: imm[11:5] rs2 rs1 000 imm[4:0] 0100011 SB
func sb(rs1, rs2, imm uint32) (code uint32) {
	code = imm>>5&(1<<7-1)<<25 | rs2<<20 | rs1<<15 | 0b000<<12 | imm&0b11111<<7 | 0b0100011
	return code
}

//32 SH: imm[11:5] rs2 rs1 001 imm[4:0] 0100011 SH
func sh(rs1, rs2, imm uint32) (code uint32) {
	code = imm>>5&(1<<7-1)<<25 | rs2<<20 | rs1<<15 | 0b001<<12 | imm&0b11111<<7 | 0b0100011
	return code
}

//33 SW: imm[11:5] rs2 rs1 010 imm[4:0] 0100011 SW
func sw(rs1, rs2, imm uint32) (code uint32) {
	code = imm>>5&(1<<7-1)<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | imm&0b11111<<7 | 0b0100011
	return code
}

//34 SD: imm[11:5] rs2 rs1 011 imm[4:0] 0100011 SD
func sd(rs1, rs2, imm uint32) (code uint32) {
	code = imm>>5&(1<<7-1)<<25 | rs2<<20 | rs1<<15 | 0b011<<12 | imm&0b11111<<7 | 0b0100011
	return code
}

//35 FSW: imm[11:5] rs2 rs1 010 imm[4:0] 0100111 FSW
func fsw(rs1, rs2, imm uint32) (code uint32) {
	code = imm>>5&(1<<7-1)<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | imm&0b11111<<7 | 0b0100111
	return code
}

//36 FSD: imm[11:5] rs2 rs1 011 imm[4:0] 0100111 FSD
func fsd(rs1, rs2, imm uint32) (code uint32) {
	code = imm>>5&(1<<7-1)<<25 | rs2<<20 | rs1<<15 | 0b011<<12 | imm&0b11111<<7 | 0b0100111
	return code
}

//37 FSQ: imm[11:5] rs2 rs1 100 imm[4:0] 0100111 FSQ
func fsq(rs1, rs2, imm uint32) (code uint32) {
	code = imm>>5&(1<<7-1)<<25 | rs2<<20 | rs1<<15 | 0b100<<12 | imm&0b11111<<7 | 0b0100111
	return code
}

//38 SLLI: 000000 shamt rs1 001 rd 0010011 SLLI
func slli(rd, rs1, shamt uint32) (code uint32) {
	code = 0b000000<<26 | shamt<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b0010011
	return code
}

//39 SRLI: 000000 shamt rs1 101 rd 0010011 SRLI
func srli(rd, rs1, shamt uint32) (code uint32) {
	code = 0b000000<<26 | shamt<<20 | rs1<<15 | 0b101<<12 | rd<<7 | 0b0010011
	return code
}

//40 SRAI: 010000 shamt rs1 101 rd 0010011 SRAI
func srai(rd, rs1, shamt uint32) (code uint32) {
	code = 0b010000<<26 | shamt<<20 | rs1<<15 | 0b101<<12 | rd<<7 | 0b0010011
	return code
}

//41 SLLIW: 0000000 shamt rs1 001 rd 0011011 SLLIW
func slliw(rd, rs1, shamt uint32) (code uint32) {
	code = 0b0000000<<25 | shamt<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b0011011
	return code
}

//42 SRLIW: 0000000 shamt rs1 101 rd 0011011 SRLIW
func srliw(rd, rs1, shamt uint32) (code uint32) {
	code = 0b0000000<<25 | shamt<<20 | rs1<<15 | 0b101<<12 | rd<<7 | 0b0011011
	return code
}

//43 SRAIW: 0100000 shamt rs1 101 rd 0011011 SRAIW
func sraiw(rd, rs1, shamt uint32) (code uint32) {
	code = 0b0100000<<25 | shamt<<20 | rs1<<15 | 0b101<<12 | rd<<7 | 0b0011011
	return code
}

//44 ADD: 0000000 rs2 rs1 000 rd 0110011 ADD
func add(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000000<<25 | rs2<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b0110011
	return code
}

//45 SUB: 0100000 rs2 rs1 000 rd 0110011 SUB
func sub(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0100000<<25 | rs2<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b0110011
	return code
}

//46 SLL: 0000000 rs2 rs1 001 rd 0110011 SLL
func sll(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000000<<25 | rs2<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b0110011
	return code
}

//47 SLT: 0000000 rs2 rs1 010 rd 0110011 SLT
func slt(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000000<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b0110011
	return code
}

//48 SLTU: 0000000 rs2 rs1 011 rd 0110011 SLTU
func sltu(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000000<<25 | rs2<<20 | rs1<<15 | 0b011<<12 | rd<<7 | 0b0110011
	return code
}

//49 XOR: 0000000 rs2 rs1 100 rd 0110011 XOR
func xor(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000000<<25 | rs2<<20 | rs1<<15 | 0b100<<12 | rd<<7 | 0b0110011
	return code
}

//50 SRL: 0000000 rs2 rs1 101 rd 0110011 SRL
func srl(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000000<<25 | rs2<<20 | rs1<<15 | 0b101<<12 | rd<<7 | 0b0110011
	return code
}

//51 SRA: 0100000 rs2 rs1 101 rd 0110011 SRA
func sra(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0100000<<25 | rs2<<20 | rs1<<15 | 0b101<<12 | rd<<7 | 0b0110011
	return code
}

//52 OR: 0000000 rs2 rs1 110 rd 0110011 OR
func or(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000000<<25 | rs2<<20 | rs1<<15 | 0b110<<12 | rd<<7 | 0b0110011
	return code
}

//53 AND: 0000000 rs2 rs1 111 rd 0110011 AND
func and(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000000<<25 | rs2<<20 | rs1<<15 | 0b111<<12 | rd<<7 | 0b0110011
	return code
}

//54 MUL: 0000001 rs2 rs1 000 rd 0110011 MUL
func mul(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000001<<25 | rs2<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b0110011
	return code
}

//55 MULH: 0000001 rs2 rs1 001 rd 0110011 MULH
func mulh(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000001<<25 | rs2<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b0110011
	return code
}

//56 MULHSU: 0000001 rs2 rs1 010 rd 0110011 MULHSU
func mulhsu(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000001<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b0110011
	return code
}

//57 MULHU: 0000001 rs2 rs1 011 rd 0110011 MULHU
func mulhu(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000001<<25 | rs2<<20 | rs1<<15 | 0b011<<12 | rd<<7 | 0b0110011
	return code
}

//58 DIV: 0000001 rs2 rs1 100 rd 0110011 DIV
func div(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000001<<25 | rs2<<20 | rs1<<15 | 0b100<<12 | rd<<7 | 0b0110011
	return code
}

//59 DIVU: 0000001 rs2 rs1 101 rd 0110011 DIVU
func divu(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000001<<25 | rs2<<20 | rs1<<15 | 0b101<<12 | rd<<7 | 0b0110011
	return code
}

//60 REM: 0000001 rs2 rs1 110 rd 0110011 REM
func rem(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000001<<25 | rs2<<20 | rs1<<15 | 0b110<<12 | rd<<7 | 0b0110011
	return code
}

//61 REMU: 0000001 rs2 rs1 111 rd 0110011 REMU
func remu(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000001<<25 | rs2<<20 | rs1<<15 | 0b111<<12 | rd<<7 | 0b0110011
	return code
}

//62 ADDW: 0000000 rs2 rs1 000 rd 0111011 ADDW
func addw(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000000<<25 | rs2<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b0111011
	return code
}

//63 SUBW: 0100000 rs2 rs1 000 rd 0111011 SUBW
func subw(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0100000<<25 | rs2<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b0111011
	return code
}

//64 SLLW: 0000000 rs2 rs1 001 rd 0111011 SLLW
func sllw(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000000<<25 | rs2<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b0111011
	return code
}

//65 SRLW: 0000000 rs2 rs1 101 rd 0111011 SRLW
func srlw(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000000<<25 | rs2<<20 | rs1<<15 | 0b101<<12 | rd<<7 | 0b0111011
	return code
}

//66 SRAW: 0100000 rs2 rs1 101 rd 0111011 SRAW
func sraw(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0100000<<25 | rs2<<20 | rs1<<15 | 0b101<<12 | rd<<7 | 0b0111011
	return code
}

//67 MULW: 0000001 rs2 rs1 000 rd 0111011 MULW
func mulw(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000001<<25 | rs2<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b0111011
	return code
}

//68 DIVW: 0000001 rs2 rs1 100 rd 0111011 DIVW
func divw(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000001<<25 | rs2<<20 | rs1<<15 | 0b100<<12 | rd<<7 | 0b0111011
	return code
}

//69 DIVUW: 0000001 rs2 rs1 101 rd 0111011 DIVUW
func divuw(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000001<<25 | rs2<<20 | rs1<<15 | 0b101<<12 | rd<<7 | 0b0111011
	return code
}

//70 REMW: 0000001 rs2 rs1 110 rd 0111011 REMW
func remw(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000001<<25 | rs2<<20 | rs1<<15 | 0b110<<12 | rd<<7 | 0b0111011
	return code
}

//71 REMUW: 0000001 rs2 rs1 111 rd 0111011 REMUW
func remuw(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0000001<<25 | rs2<<20 | rs1<<15 | 0b111<<12 | rd<<7 | 0b0111011
	return code
}

//72 FLE.S: 1010000 rs2 rs1 000 rd 1010011 FLE.S
func fles(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b1010000<<25 | rs2<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b1010011
	return code
}

//73 FLE.D: 1010001 rs2 rs1 000 rd 1010011 FLE.D
func fled(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b1010001<<25 | rs2<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b1010011
	return code
}

//74 FLE.Q: 1010011 rs2 rs1 000 rd 1010011 FLE.Q
func fleq(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b1010011<<25 | rs2<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b1010011
	return code
}

//75 FLT.S: 1010000 rs2 rs1 001 rd 1010011 FLT.S
func flts(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b1010000<<25 | rs2<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b1010011
	return code
}

//76 FLT.D: 1010001 rs2 rs1 001 rd 1010011 FLT.D
func fltd(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b1010001<<25 | rs2<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b1010011
	return code
}

//77 FLT.Q: 1010011 rs2 rs1 001 rd 1010011 FLT.Q
func fltq(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b1010011<<25 | rs2<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b1010011
	return code
}

//78 FEQ.S: 1010000 rs2 rs1 010 rd 1010011 FEQ.S
func feqs(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b1010000<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b1010011
	return code
}

//79 FEQ.D: 1010001 rs2 rs1 010 rd 1010011 FEQ.D
func feqd(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b1010001<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b1010011
	return code
}

//80 FEQ.Q: 1010011 rs2 rs1 010 rd 1010011 FEQ.Q
func feqq(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b1010011<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b1010011
	return code
}

//81 FSGNJ.S: 0010000 rs2 rs1 000 rd 1010011 FSGNJ.S
func fsgnjs(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0010000<<25 | rs2<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b1010011
	return code
}

//82 FSGNJ.D: 0010001 rs2 rs1 000 rd 1010011 FSGNJ.D
func fsgnjd(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0010001<<25 | rs2<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b1010011
	return code
}

//83 FSGNJ.Q: 0010011 rs2 rs1 000 rd 1010011 FSGNJ.Q
func fsgnjq(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0010011<<25 | rs2<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b1010011
	return code
}

//84 FSGNJN.S: 0010000 rs2 rs1 001 rd 1010011 FSGNJN.S
func fsgnjns(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0010000<<25 | rs2<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b1010011
	return code
}

//85 FSGNJN.D: 0010001 rs2 rs1 001 rd 1010011 FSGNJN.D
func fsgnjnd(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0010001<<25 | rs2<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b1010011
	return code
}

//86 FSGNJN.Q: 0010011 rs2 rs1 001 rd 1010011 FSGNJN.Q
func fsgnjnq(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0010011<<25 | rs2<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b1010011
	return code
}

//87 FSGNJX.S: 0010000 rs2 rs1 010 rd 1010011 FSGNJX.S
func fsgnjxs(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0010000<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b1010011
	return code
}

//88 FSGNJX.D: 0010001 rs2 rs1 010 rd 1010011 FSGNJX.D
func fsgnjxd(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0010001<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b1010011
	return code
}

//89 FSGNJX.Q: 0010011 rs2 rs1 010 rd 1010011 FSGNJX.Q
func fsgnjxq(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0010011<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b1010011
	return code
}

//90 FMIN.S: 0010100 rs2 rs1 000 rd 1010011 FMIN.S
func fmins(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0010100<<25 | rs2<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b1010011
	return code
}

//91 FMIN.D: 0010101 rs2 rs1 000 rd 1010011 FMIN.D
func fmind(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0010101<<25 | rs2<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b1010011
	return code
}

//92 FMIN.Q: 0010111 rs2 rs1 000 rd 1010011 FMIN.Q
func fminq(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0010111<<25 | rs2<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b1010011
	return code
}

//93 FMAX.S: 0010100 rs2 rs1 001 rd 1010011 FMAX.S
func fmaxs(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0010100<<25 | rs2<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b1010011
	return code
}

//94 FMAX.D: 0010101 rs2 rs1 001 rd 1010011 FMAX.D
func fmaxd(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0010101<<25 | rs2<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b1010011
	return code
}

//95 FMAX.Q: 0010111 rs2 rs1 001 rd 1010011 FMAX.Q
func fmaxq(rd, rs1, rs2 uint32) (code uint32) {
	code = 0b0010111<<25 | rs2<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b1010011
	return code
}

//96 FADD.S: 0000000 rs2 rs1 rm rd 1010011 FADD.S
func fadds(rd, rm, rs1, rs2 uint32) (code uint32) {
	code = 0b0000000<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//97 FADD.D: 0000001 rs2 rs1 rm rd 1010011 FADD.D
func faddd(rd, rm, rs1, rs2 uint32) (code uint32) {
	code = 0b0000001<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//98 FADD.Q: 0000011 rs2 rs1 rm rd 1010011 FADD.Q
func faddq(rd, rm, rs1, rs2 uint32) (code uint32) {
	code = 0b0000011<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//99 FSUB.S: 0000100 rs2 rs1 rm rd 1010011 FSUB.S
func fsubs(rd, rm, rs1, rs2 uint32) (code uint32) {
	code = 0b0000100<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//100 FSUB.D: 0000101 rs2 rs1 rm rd 1010011 FSUB.D
func fsubd(rd, rm, rs1, rs2 uint32) (code uint32) {
	code = 0b0000101<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//101 FSUB.Q: 0000111 rs2 rs1 rm rd 1010011 FSUB.Q
func fsubq(rd, rm, rs1, rs2 uint32) (code uint32) {
	code = 0b0000111<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//102 FMUL.S: 0001000 rs2 rs1 rm rd 1010011 FMUL.S
func fmuls(rd, rm, rs1, rs2 uint32) (code uint32) {
	code = 0b0001000<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//103 FMUL.D: 0001001 rs2 rs1 rm rd 1010011 FMUL.D
func fmuld(rd, rm, rs1, rs2 uint32) (code uint32) {
	code = 0b0001001<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//104 FMUL.Q: 0001011 rs2 rs1 rm rd 1010011 FMUL.Q
func fmulq(rd, rm, rs1, rs2 uint32) (code uint32) {
	code = 0b0001011<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//105 FDIV.S: 0001100 rs2 rs1 rm rd 1010011 FDIV.S
func fdivs(rd, rm, rs1, rs2 uint32) (code uint32) {
	code = 0b0001100<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//106 FDIV.D: 0001101 rs2 rs1 rm rd 1010011 FDIV.D
func fdivd(rd, rm, rs1, rs2 uint32) (code uint32) {
	code = 0b0001101<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//107 FDIV.Q: 0001111 rs2 rs1 rm rd 1010011 FDIV.Q
func fdivq(rd, rm, rs1, rs2 uint32) (code uint32) {
	code = 0b0001111<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//108 FCVT.S.Q: 0100000 00011 rs1 rm rd 1010011 FCVT.S.Q
func fcvtsq(rd, rm, rs1 uint32) (code uint32) {
	code = 0b0100000<<25 | 0b00011<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//109 FCVT.S.D: 0100000 00001 rs1 rm rd 1010011 FCVT.S.D
func fcvtsd(rd, rm, rs1 uint32) (code uint32) {
	code = 0b0100000<<25 | 0b00001<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//110 FCVT.D.S: 0100001 00000 rs1 rm rd 1010011 FCVT.D.S
func fcvtds(rd, rm, rs1 uint32) (code uint32) {
	code = 0b0100001<<25 | 0b00000<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//111 FCVT.D.Q: 0100001 00011 rs1 rm rd 1010011 FCVT.D.Q
func fcvtdq(rd, rm, rs1 uint32) (code uint32) {
	code = 0b0100001<<25 | 0b00011<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//112 FCVT.Q.S: 0100011 00000 rs1 rm rd 1010011 FCVT.Q.S
func fcvtqs(rd, rm, rs1 uint32) (code uint32) {
	code = 0b0100011<<25 | 0b00000<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//113 FCVT.Q.D: 0100011 00001 rs1 rm rd 1010011 FCVT.Q.D
func fcvtqd(rd, rm, rs1 uint32) (code uint32) {
	code = 0b0100011<<25 | 0b00001<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//114 FSQRT.S: 0101100 00000 rs1 rm rd 1010011 FSQRT.S
func fsqrts(rd, rm, rs1 uint32) (code uint32) {
	code = 0b0101100<<25 | 0b00000<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//115 FSQRT.D: 0101101 00000 rs1 rm rd 1010011 FSQRT.D
func fsqrtd(rd, rm, rs1 uint32) (code uint32) {
	code = 0b0101101<<25 | 0b00000<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//116 FSQRT.Q: 0101111 00000 rs1 rm rd 1010011 FSQRT.Q
func fsqrtq(rd, rm, rs1 uint32) (code uint32) {
	code = 0b0101111<<25 | 0b00000<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//117 FCVT.W.S: 1100000 00000 rs1 rm rd 1010011 FCVT.W.S
func fcvtws(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1100000<<25 | 0b00000<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//118 FCVT.WU.S: 1100000 00001 rs1 rm rd 1010011 FCVT.WU.S
func fcvtwus(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1100000<<25 | 0b00001<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//119 FCVT.L.S: 1100000 00010 rs1 rm rd 1010011 FCVT.L.S
func fcvtls(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1100000<<25 | 0b00010<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//120 FCVT.LU.S: 1100000 00011 rs1 rm rd 1010011 FCVT.LU.S
func fcvtlus(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1100000<<25 | 0b00011<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//121 FCVT.W.D: 1100001 00000 rs1 rm rd 1010011 FCVT.W.D
func fcvtwd(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1100001<<25 | 0b00000<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//122 FCVT.WU.D: 1100001 00001 rs1 rm rd 1010011 FCVT.WU.D
func fcvtwud(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1100001<<25 | 0b00001<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//123 FCVT.L.D: 1100001 00010 rs1 rm rd 1010011 FCVT.L.D
func fcvtld(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1100001<<25 | 0b00010<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//124 FCVT.LU.D: 1100001 00011 rs1 rm rd 1010011 FCVT.LU.D
func fcvtlud(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1100001<<25 | 0b00011<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//125 FCVT.W.Q: 1100011 00000 rs1 rm rd 1010011 FCVT.W.Q
func fcvtwq(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1100011<<25 | 0b00000<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//126 FCVT.WU.Q: 1100011 00001 rs1 rm rd 1010011 FCVT.WU.Q
func fcvtwuq(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1100011<<25 | 0b00001<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//127 FCVT.L.Q: 1100011 00010 rs1 rm rd 1010011 FCVT.L.Q
func fcvtlq(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1100011<<25 | 0b00010<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//128 FCVT.LU.Q: 1100011 00011 rs1 rm rd 1010011 FCVT.LU.Q
func fcvtluq(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1100011<<25 | 0b00011<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//129 FCVT.S.W: 1101000 00000 rs1 rm rd 1010011 FCVT.S.W
func fcvtsw(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1101000<<25 | 0b00000<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//130 FCVT.S.WU: 1101000 00001 rs1 rm rd 1010011 FCVT.S.WU
func fcvtswu(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1101000<<25 | 0b00001<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//131 FCVT.S.L: 1101000 00010 rs1 rm rd 1010011 FCVT.S.L
func fcvtsl(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1101000<<25 | 0b00010<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//132 FCVT.S.LU: 1101000 00011 rs1 rm rd 1010011 FCVT.S.LU
func fcvtslu(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1101000<<25 | 0b00011<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//133 FCVT.D.W: 1101001 00000 rs1 rm rd 1010011 FCVT.D.W
func fcvtdw(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1101001<<25 | 0b00000<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//134 FCVT.D.WU: 1101001 00001 rs1 rm rd 1010011 FCVT.D.WU
func fcvtdwu(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1101001<<25 | 0b00001<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//135 FCVT.D.L: 1101001 00010 rs1 rm rd 1010011 FCVT.D.L
func fcvtdl(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1101001<<25 | 0b00010<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//136 FCVT.D.LU: 1101001 00011 rs1 rm rd 1010011 FCVT.D.LU
func fcvtdlu(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1101001<<25 | 0b00011<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//137 FCVT.Q.W: 1101011 00000 rs1 rm rd 1010011 FCVT.Q.W
func fcvtqw(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1101011<<25 | 0b00000<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//138 FCVT.Q.WU: 1101011 00001 rs1 rm rd 1010011 FCVT.Q.WU
func fcvtqwu(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1101011<<25 | 0b00001<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//139 FCVT.Q.L: 1101011 00010 rs1 rm rd 1010011 FCVT.Q.L
func fcvtql(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1101011<<25 | 0b00010<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//140 FCVT.Q.LU: 1101011 00011 rs1 rm rd 1010011 FCVT.Q.LU
func fcvtqlu(rd, rm, rs1 uint32) (code uint32) {
	code = 0b1101011<<25 | 0b00011<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1010011
	return code
}

//141 FMV.W.X: 1111000 00000 rs1 000 rd 1010011 FMV.W.X
func fmvwx(rd, rs1 uint32) (code uint32) {
	code = 0b1111000<<25 | 0b00000<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b1010011
	return code
}

//142 FMV.X.W: 1110000 00000 rs1 000 rd 1010011 FMV.X.W
func fmvxw(rd, rs1 uint32) (code uint32) {
	code = 0b1110000<<25 | 0b00000<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b1010011
	return code
}

//143 FMV.D.X: 1111001 00000 rs1 000 rd 1010011 FMV.D.X
func fmvdx(rd, rs1 uint32) (code uint32) {
	code = 0b1111001<<25 | 0b00000<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b1010011
	return code
}

//144 FMV.X.D: 1110001 00000 rs1 000 rd 1010011 FMV.X.D
func fmvxd(rd, rs1 uint32) (code uint32) {
	code = 0b1110001<<25 | 0b00000<<20 | rs1<<15 | 0b000<<12 | rd<<7 | 0b1010011
	return code
}

//145 FCLASS.S: 1110000 00000 rs1 001 rd 1010011 FCLASS.S
func fclasss(rd, rs1 uint32) (code uint32) {
	code = 0b1110000<<25 | 0b00000<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b1010011
	return code
}

//146 FCLASS.D: 1110001 00000 rs1 001 rd 1010011 FCLASS.D
func fclassd(rd, rs1 uint32) (code uint32) {
	code = 0b1110001<<25 | 0b00000<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b1010011
	return code
}

//147 FCLASS.Q: 1110011 00000 rs1 001 rd 1010011 FCLASS.Q
func fclassq(rd, rs1 uint32) (code uint32) {
	code = 0b1110011<<25 | 0b00000<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b1010011
	return code
}

//148 FMADD.S: rs3 00 rs2 rs1 rm rd 1000011 FMADD.S
func fmadds(rd, rm, rs1, rs2, rs3 uint32) (code uint32) {
	code = rs3<<27 | 0b00<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1000011
	return code
}

//149 FMADD.D: rs3 01 rs2 rs1 rm rd 1000011 FMADD.D
func fmaddd(rd, rm, rs1, rs2, rs3 uint32) (code uint32) {
	code = rs3<<27 | 0b01<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1000011
	return code
}

//150 FMADD.Q: rs3 11 rs2 rs1 rm rd 1000011 FMADD.Q
func fmaddq(rd, rm, rs1, rs2, rs3 uint32) (code uint32) {
	code = rs3<<27 | 0b11<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1000011
	return code
}

//151 FMSUB.S: rs3 00 rs2 rs1 rm rd 1000111 FMSUB.S
func fmsubs(rd, rm, rs1, rs2, rs3 uint32) (code uint32) {
	code = rs3<<27 | 0b00<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1000111
	return code
}

//152 FMSUB.D: rs3 01 rs2 rs1 rm rd 1000111 FMSUB.D
func fmsubd(rd, rm, rs1, rs2, rs3 uint32) (code uint32) {
	code = rs3<<27 | 0b01<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1000111
	return code
}

//153 FMSUB.Q: rs3 11 rs2 rs1 rm rd 1000111 FMSUB.Q
func fmsubq(rd, rm, rs1, rs2, rs3 uint32) (code uint32) {
	code = rs3<<27 | 0b11<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1000111
	return code
}

//154 FNMSUB.S: rs3 00 rs2 rs1 rm rd 1001011 FNMSUB.S
func fnmsubs(rd, rm, rs1, rs2, rs3 uint32) (code uint32) {
	code = rs3<<27 | 0b00<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1001011
	return code
}

//155 FNMSUB.D: rs3 01 rs2 rs1 rm rd 1001011 FNMSUB.D
func fnmsubd(rd, rm, rs1, rs2, rs3 uint32) (code uint32) {
	code = rs3<<27 | 0b01<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1001011
	return code
}

//156 FNMSUB.Q: rs3 11 rs2 rs1 rm rd 1001011 FNMSUB.Q
func fnmsubq(rd, rm, rs1, rs2, rs3 uint32) (code uint32) {
	code = rs3<<27 | 0b11<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1001011
	return code
}

//157 FNMADD.S: rs3 00 rs2 rs1 rm rd 1001111 FNMADD.S
func fnmadds(rd, rm, rs1, rs2, rs3 uint32) (code uint32) {
	code = rs3<<27 | 0b00<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1001111
	return code
}

//158 FNMADD.D: rs3 01 rs2 rs1 rm rd 1001111 FNMADD.D
func fnmaddd(rd, rm, rs1, rs2, rs3 uint32) (code uint32) {
	code = rs3<<27 | 0b01<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1001111
	return code
}

//159 FNMADD.Q: rs3 11 rs2 rs1 rm rd 1001111 FNMADD.Q
func fnmaddq(rd, rm, rs1, rs2, rs3 uint32) (code uint32) {
	code = rs3<<27 | 0b11<<25 | rs2<<20 | rs1<<15 | rm<<12 | rd<<7 | 0b1001111
	return code
}

//161 AMOADD.W: 00000 aq rl rs2 rs1 010 rd 0101111 AMOADD.W
func amoaddw(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b00000<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b0101111
	return code
}

//162 AMOSWAP.W: 00001 aq rl rs2 rs1 010 rd 0101111 AMOSWAP.W
func amoswapw(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b00001<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b0101111
	return code
}

//163 LR.W: 00010 aq rl 00000 rs1 010 rd 0101111 LR.W
func lrw(rd, rs1, rl, aq uint32) (code uint32) {
	code = 0b00010<<27 | aq<<26 | rl<<25 | 0b00000<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b0101111
	return code
}

//164 SC.W: 00011 aq rl rs2 rs1 010 rd 0101111 SC.W
func scw(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b00011<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b0101111
	return code
}

//165 AMOXOR.W: 00100 aq rl rs2 rs1 010 rd 0101111 AMOXOR.W
func amoxorw(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b00100<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b0101111
	return code
}

//166 AMOOR.W: 01000 aq rl rs2 rs1 010 rd 0101111 AMOOR.W
func amoorw(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b01000<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b0101111
	return code
}

//167 AMOAND.W: 01100 aq rl rs2 rs1 010 rd 0101111 AMOAND.W
func amoandw(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b01100<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b0101111
	return code
}

//168 AMOMIN.W: 10000 aq rl rs2 rs1 010 rd 0101111 AMOMIN.W
func amominw(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b10000<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b0101111
	return code
}

//169 AMOMAX.W: 10100 aq rl rs2 rs1 010 rd 0101111 AMOMAX.W
func amomaxw(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b10100<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b0101111
	return code
}

//170 AMOMINU.W: 11000 aq rl rs2 rs1 010 rd 0101111 AMOMINU.W
func amominuw(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b11000<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b0101111
	return code
}

//171 AMOMAXU.W: 11100 aq rl rs2 rs1 010 rd 0101111 AMOMAXU.W
func amomaxuw(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b11100<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b0101111
	return code
}

//172 AMOADD.D: 00000 aq rl rs2 rs1 011 rd 0101111 AMOADD.D
func amoaddd(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b00000<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b011<<12 | rd<<7 | 0b0101111
	return code
}

//173 AMOSWAP.D: 00001 aq rl rs2 rs1 011 rd 0101111 AMOSWAP.D
func amoswapd(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b00001<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b011<<12 | rd<<7 | 0b0101111
	return code
}

//174 LR.D: 00010 aq rl 00000 rs1 011 rd 0101111 LR.D
func lrd(rd, rs1, rl, aq uint32) (code uint32) {
	code = 0b00010<<27 | aq<<26 | rl<<25 | 0b00000<<20 | rs1<<15 | 0b011<<12 | rd<<7 | 0b0101111
	return code
}

//175 SC.D: 00011 aq rl rs2 rs1 011 rd 0101111 SC.D
func scd(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b00011<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b011<<12 | rd<<7 | 0b0101111
	return code
}

//176 AMOXOR.D: 00100 aq rl rs2 rs1 011 rd 0101111 AMOXOR.D
func amoxord(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b00100<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b011<<12 | rd<<7 | 0b0101111
	return code
}

//177 AMOOR.D: 01000 aq rl rs2 rs1 011 rd 0101111 AMOOR.D
func amoord(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b01000<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b011<<12 | rd<<7 | 0b0101111
	return code
}

//178 AMOAND.D: 01100 aq rl rs2 rs1 011 rd 0101111 AMOAND.D
func amoandd(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b01100<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b011<<12 | rd<<7 | 0b0101111
	return code
}

//179 AMOMIN.D: 10000 aq rl rs2 rs1 011 rd 0101111 AMOMIN.D
func amomind(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b10000<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b011<<12 | rd<<7 | 0b0101111
	return code
}

//180 AMOMAX.D: 10100 aq rl rs2 rs1 011 rd 0101111 AMOMAX.D
func amomaxd(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b10100<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b011<<12 | rd<<7 | 0b0101111
	return code
}

//181 AMOMINU.D: 11000 aq rl rs2 rs1 011 rd 0101111 AMOMINU.D
func amominud(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b11000<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b011<<12 | rd<<7 | 0b0101111
	return code
}

//182 AMOMAXU.D: 11100 aq rl rs2 rs1 011 rd 0101111 AMOMAXU.D
func amomaxud(rd, rs1, rs2, rl, aq uint32) (code uint32) {
	code = 0b11100<<27 | aq<<26 | rl<<25 | rs2<<20 | rs1<<15 | 0b011<<12 | rd<<7 | 0b0101111
	return code
}

//184 CSRRW: csr rs1 001 rd 1110011 CSRRW
func csrrw(rd, csr, rs1 uint32) (code uint32) {
	code = csr<<20 | rs1<<15 | 0b001<<12 | rd<<7 | 0b1110011
	return code
}

//185 CSRRS: csr rs1 010 rd 1110011 CSRRS
func csrrs(rd, csr, rs1 uint32) (code uint32) {
	code = csr<<20 | rs1<<15 | 0b010<<12 | rd<<7 | 0b1110011
	return code
}

//186 CSRRC: csr rs1 011 rd 1110011 CSRRC
func csrrc(rd, csr, rs1 uint32) (code uint32) {
	code = csr<<20 | rs1<<15 | 0b011<<12 | rd<<7 | 0b1110011
	return code
}

//187 CSRRWI: csr uimm 101 rd 1110011 CSRRWI
func csrrwi(rd, csr, uimm uint32) (code uint32) {
	code = csr<<20 | uimm<<15 | 0b101<<12 | rd<<7 | 0b1110011
	return code
}

//188 CSRRSI: csr uimm 110 rd 1110011 CSRRSI
func csrrsi(rd, csr, uimm uint32) (code uint32) {
	code = csr<<20 | uimm<<15 | 0b110<<12 | rd<<7 | 0b1110011
	return code
}

//189 CSRRCI: csr uimm 111 rd 1110011 CSRRCI
func csrrci(rd, csr, uimm uint32) (code uint32) {
	code = csr<<20 | uimm<<15 | 0b111<<12 | rd<<7 | 0b1110011
	return code
}
