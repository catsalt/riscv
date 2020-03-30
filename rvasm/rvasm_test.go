// 1.go
package rvasm

import (
	"fmt"
	"testing"
)

// zPtest...
func TestCptcs(t *testing.T) {
	fmt.Println("...")
	// zPtestA()
	// zPtestB()
	// zPtestC()

}

// zParse, zPfile
func TestAparse(t *testing.T) {
	zPfile("test", "test")
	// c := bEncode("addi", []uint32{1, 2, 3})
	// fmt.Println(c)
	// fmt.Println(addi(1, 2, 3))
}

// type F func(string) byte

// func BenchmarkA(b *testing.B) {
// 	fn := func(fname, s string, f F) {
// 		b.Run(fname, func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				f(s)
// 			}
// 		})
// 		fmt.Println(fname, s, f(s))
// 	}
// }
