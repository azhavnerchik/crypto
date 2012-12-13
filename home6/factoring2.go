package main

//import "log"
import "fmt"

//import "math"
import "math/big"
import "github.com/cznic/mathutil"

func main() {
	N, _ := new(big.Int).SetString("648455842808071669662824265346772278726343720706976263060439070378797308618081116462714015276061417569195587321840254520655424906719892428844841839353281972988531310511738648965962582821502504990264452100885281673303711142296421027840289307657458645233683357077834689715838646088239640236866252211790085787877", 0)

	A := mathutil.SqrtBig(N)
	A.Add(A, big.NewInt(1))

	res := -1
	one := big.NewInt(1)
	for A.Cmp(N) == -1 && res != 0 {
		res = check_factor(A, N)
		A.Add(A, one)
	}
}

func check_factor(A, N *big.Int) int {
	x := new(big.Int).Mul(A, A)
	x = mathutil.SqrtBig(x.Sub(x, N))

	p := new(big.Int).Sub(A, x)
	q := new(big.Int).Add(A, x)

	res := new(big.Int).Mul(p, q).Cmp(N)
	if res == 0 {
		fmt.Println("Result:", p)
	}

	return res
}
