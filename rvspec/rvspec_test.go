package rvspec

import (
	"fmt"
	"testing"
)

//187 FENCE _ fm pred succ rs1 000 rd 0001111 FENCE
//一定得按命名规范, Test+Cap
func TestRv(t *testing.T) {
	fmt.Println("....")

	// add := aTasmfunc(1, "0000000 rs2 rs1 000 rd 0110011 ADD")
	// ebreak := aTasmfunc(1, "000000000001 00000 000 00000 1110011 EBREAK")
	// ecall := aTasmfunc(1, "000000000000 00000 000 00000 1110011 ECALL")
	// fence := aTasmfunc(1, "fm pred succ rs1 000 rd 0001111 FENCE")
	// fmt.Println(add)
	// fmt.Println(ebreak)
	// fmt.Println(ecall)
	// fmt.Println(fence)
	zTrfile("rvall188.txt", "riscv.go")
	// zTorder("rvAll188mapkv.txt","rvEncode")
	rvFile("rvall188.txt", "rvEncode.go")

}
