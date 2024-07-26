package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(s float64) float64 {
		delta := v0*v0 - 4*0.5*a*(s0-s)
		if delta < 0 {
			return math.NaN()
		}

		t1 := (-v0 + math.Sqrt(delta)) / (2 * 0.5 * a)
		t2 := (-v0 - math.Sqrt(delta)) / (2 * 0.5 * a)

		if t1 >= 0 {
			return t1
		}
		return t2
	}
}

func main() {
	a1, v0_1, s0_1 := 9.8, 0.0, 0.0
	s1 := 19.6
	displaceFn1 := GenDisplaceFn(a1, v0_1, s0_1)
	fmt.Printf("Teste 1:\\n")
	fmt.Printf("Tempo para deslocamento %.2f (a = %.1f, v0 = %.1f, s0 = %.1f) é %.2f segundos\\n", s1, a1, v0_1, s0_1, displaceFn1(s1))

	a2, v0_2, s0_2 := 5.0, 2.0, 10.0
	s2 := 25.0
	displaceFn2 := GenDisplaceFn(a2, v0_2, s0_2)
	fmt.Printf("Teste 2:\\n")
	fmt.Printf("Tempo para deslocamento %.2f (a = %.1f, v0 = %.1f, s0 = %.1f) é %.2f segundos\\n", s2, a2, v0_2, s0_2, displaceFn2(s2))
}
