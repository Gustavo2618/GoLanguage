package main

import (
	"fmt"
	"math/rand"
	"time"
)

// randomFloatInRange gera um número float64 aleatório entre min e max.
func randomFloatInRange(min, max float64) float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Float64()*(max-min)
}

func main() {
	// Usa a hora atual como semente para o gerador de números aleatórios.
	rand := randomFloatInRange(0.00, 0.03)
	fmt.Printf("gerado numero aleatorio: %.04f", rand)
	// Gera e exibe três números float64 aleatórios no intervalo [min, max].

}

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"text/tabwriter"
// 	"time"
// )

// // randomInRange gera um número aleatório entre min e max.
// func randomInRange(r *rand.Rand, min, max int) int {
// 	return r.Intn(max-min+1) + min
// }

// func main() {
// 	// Usa a hora atual como semente para o gerador de números aleatórios.
// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))

// 	// Tabwriter ajuda a gerar uma saída alinhada.
// 	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
// 	defer w.Flush()
// 	show := func(name string, v1, v2, v3 any) {
// 		fmt.Fprintf(w, "%s\t%v\t%v\t%v\n", name, v1, v2, v3)
// 	}

// 	// Gera números aleatórios no intervalo [min, max].
// 	min, max := 5, 15
// 	show("RandomInRange", randomInRange(r, min, max), randomInRange(r, min, max), randomInRange(r, min, max))

// 	// Float32 e Float64 são valores no intervalo [0, 1).
// 	show("Float32", r.Float32(), r.Float32(), r.Float32())
// 	show("Float64", r.Float64(), r.Float64(), r.Float64())

// 	// ExpFloat64 tem uma média de 1 mas decai exponencialmente.
// 	show("ExpFloat64", r.ExpFloat64(), r.ExpFloat64(), r.ExpFloat64())

// 	// NormFloat64 tem uma média de 0 e um desvio padrão de 1.
// 	show("NormFloat64", r.NormFloat64(), r.NormFloat64(), r.NormFloat64())

// 	// Int31, Int63 e Uint32 geram valores com a largura dada.
// 	show("Int31", r.Int31(), r.Int31(), r.Int31())
// 	show("Int63", r.Int63(), r.Int63(), r.Int63())
// 	show("Uint32", r.Uint32(), r.Uint32(), r.Uint32())

// 	// Intn, Int31n e Int63n limitam sua saída para ser < n.
// 	show("Intn(10)", r.Intn(10), r.Intn(10), r.Intn(10))
// 	show("Int31n(10)", r.Int31n(10), r.Int31n(10), r.Int31n(10))
// 	show("Int63n(10)", r.Int63n(10), r.Int63n(10), r.Int63n(10))

// 	// Perm gera uma permutação aleatória dos números [0, n).
// 	show("Perm", r.Perm(5), r.Perm(5), r.Perm(5))
// }
