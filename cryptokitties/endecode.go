// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cryptokitties

import (
	"math/big"
	// "strings"
	// "strconv"
	"math/rand"
	// ethereum "github.com/ethereum/go-ethereum"
	// "github.com/ethereum/go-ethereum/accounts/abi"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/core/types"
	// "github.com/ethereum/go-ethereum/event"
	"fmt"
	"time"
)

var (
// ALPHABET = "123456789abcdefghijkmnopqrstuvwx"//此实现不是32进制表示，正确的应该是从0-w，在golang中需要将转换完的32进制
// 			   0123456789abcdefghijklmnopqrstuv
// BASE     = 32   //## 32 chars/letters/digits

)

func ConvertTo32(hex string) (ret string) {
	m1 := make(map[string]string)
	m1["0"] = "1"
	m1["1"] = "2"
	m1["2"] = "3"
	m1["3"] = "4"
	m1["4"] = "5"
	m1["5"] = "6"
	m1["6"] = "7"
	m1["7"] = "8"
	m1["8"] = "9"
	m1["9"] = "a"
	m1["a"] = "b"
	m1["b"] = "c"
	m1["c"] = "d"
	m1["d"] = "e"
	m1["e"] = "f"
	m1["f"] = "g"
	m1["g"] = "h"
	m1["h"] = "i"
	m1["i"] = "j"
	m1["j"] = "k"
	m1["k"] = "m"
	m1["l"] = "n"
	m1["m"] = "o"
	m1["n"] = "p"
	m1["o"] = "q"
	m1["p"] = "r"
	m1["q"] = "s"
	m1["r"] = "t"
	m1["s"] = "u"
	m1["t"] = "v"
	m1["u"] = "w"
	m1["v"] = "x"

	g1 := big.NewInt(0)
	g1.SetString(hex, 16)
	// fmt.Println(g1.Text(16))
	// fmt.Println(g1.Text(32))
	ori := g1.Text(32)
	for _, v := range ori {
		ret += m1[string(v)]
	}
	return
}
func iseven(c rune) bool {
	switch c {
	case '2': //123456789abcdefghijkmnopqrstuvwx
	case '4':
	case '6':
	case '8':
	case 'a':
	case 'c':
	case 'e':
	case 'g':
	case 'i':
	case 'k':
	case 'n':
	case 'p':
	case 'r':
	case 't':
	case 'v':
	case 'x':
		return true
	default:
		// freebsd, openbsd,
		// plan9, windows...
		// fmt.Printf("%s.", os)
		return false
	}
	return false
}
func MixGenes(mGenes, sGenes string) (ret string) {
	// 	Here is a quick sketch of the relative odds of getting a specific gene from the parents

	// 75% - either dominant gene [D1] from parent A or B

	// 18.75% (75/4) - chance of getting either 1st recessive [R1] from A or B

	// 4.69% (75/4²) - chance of getting either 2nd recessive [R2] from A or B

	// 1.17% (75/4³) - chance of getting either 3rd recessive [R3] from A or B

	// 25% - chance of getting a mutation given A & B contain the right gene pairs
	// 	ddca578ka4f7949p4d11535kaeea175h846k2243aa9gfdcd
	//	ddca5k78a47f994p4d11553keaea175h846k4223aa9gdfcd
	//  c9am65567ff7b9gg1d1138539f77647577k46784f9gpfcaa
	//  c9am65657ff7b9gg11d183359f77647577k46748f9gpfcaa
	// mGenes:: ddca5k78a47f994p4d11553keaea175h846k4223aa9gdfcd
	// sGenes:: c9am65657ff7b9gg11d183359f77647577k46748f9gpfcaa
	// c9am65657ff7b9gg11d183359f77647577k46748f9gpfcaa
	// c9cm65657477b9gg41d1533k9aea147577644243aa9pfcaa
	rand.Seed(time.Now().UnixNano())
	var babyGenes [48]rune
	for i := 0; i < 12; i++ {
		index := 4 * i
		for j := 3; j > 0; j-- {
			ran := rand.Float64()
			// fmt.Println("ran:", ran)
			if ran < 0.25 {
				chars := []rune(mGenes)
				chars[index+j], chars[index+j-1] = chars[index+j-1], chars[index+j]
				mGenes = string(chars)
			}
			ran2 := rand.Float64()
			// fmt.Println("ran2:", ran2)
			if ran2 < 0.25 {
				chars := []rune(sGenes)
				chars[index+j], chars[index+j-1] = chars[index+j-1], chars[index+j]
				sGenes = string(chars)
			}
		}
	}
	fmt.Println("mGenes::", mGenes)
	fmt.Println("sGenes::", sGenes)
	for i := 0; i < 48; i++ {
		mutation := 0
		if i%4 == 0 {
			gene1 := mGenes[i]
			gene2 := sGenes[i]
			if gene1 > gene2 {
				gene1, gene2 = gene2, gene1
			}

			if (gene2-gene1) == 1 && iseven(rune(gene1)) {
				probability := 0.25
				if gene1 > 23 {
					probability /= 2
				}

				if rand.Float64() < probability {
					mutation = (int(gene1) / 2) + 16
				}
			}
		}
		if mutation > 0 {
			babyGenes[i] = rune(mutation)
		} else {
			if rand.Float64() < 0.5 {
				babyGenes[i] = rune(mGenes[i])
			} else {
				babyGenes[i] = rune(sGenes[i])
			}
		}
	}
	ret = string(babyGenes[:])
	return
}
