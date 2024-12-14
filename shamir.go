package main

import (
	"fmt"
	"math/rand"
)

const p = 997

func mod(a, b int) int {
	res := a % b
	if res < 0 {
		res += b
	}
	return res
}

func mulInv(a, p int) int {
	for i := 1; i < p; i++ {
		if mod(a*i, p) == 1 {
			return i
		}
	}
	return -1
}

func lagrangeInterpolation(shares []int, xCoords []int) int {
	secret := 0
	n := len(shares)

	for i := 0; i < n; i++ {
		li := 1
		for j := 0; j < n; j++ {
			if i != j {
				li = mod(li*(0-xCoords[j]), p) // x = 0 в лагранже
				li = mod(li*mulInv(xCoords[i]-xCoords[j], p), p)
			}
		}
		secret = mod(secret+mod(shares[i]*li, p), p)
	}

	return secret
}

func splitSecret(secret, n, k int) ([]int, []int) {
	fmt.Println("[Генерация многочлена и долей]")
	coefficients := make([]int, k)
	coefficients[0] = secret
	for i := 1; i < k; i++ {
		coefficients[i] = rand.Intn(p-1) + 1
	}

	fmt.Print("Многочлен: f(x) = ")
	for i := k - 1; i >= 0; i-- {
		if i == 0 {
			fmt.Printf("%d", coefficients[i])
		} else if i == 1 {
			fmt.Printf("%dx + ", coefficients[i])
		} else {
			fmt.Printf("%dx^%d + ", coefficients[i], i)
		}
	}
	fmt.Println()

	shares := make([]int, n)
	xCoords := make([]int, n)
	for i := 1; i <= n; i++ {
		xCoords[i-1] = i
		y := 0
		for j := 0; j < k; j++ {
			y = mod(y+coefficients[j]*pow(i, j, p), p)
		}
		shares[i-1] = y
		fmt.Printf("Доля %d: x = %d, y = %d\n", i, xCoords[i-1], y)
	}

	return shares, xCoords
}

func pow(base, exp, m int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result = mod(result*base, m)
		}
		base = mod(base*base, m)
		exp /= 2
	}
	return result
}

func main() {
	secret := 127
	n := 10
	k := 9
	//kTrue := 4
	//kFalse := 2

	shares, xCoords := splitSecret(secret, n, k)
	fmt.Println("\n[Итог]")
	fmt.Println("Доли участников (y):", shares)
	fmt.Println("Координаты участников (x):", xCoords)

	fmt.Println("\n----------------")

	reconstructedSecret := lagrangeInterpolation(shares[:k], xCoords[:k])
	fmt.Printf("\nВосстановленный секрет: %d\n", reconstructedSecret)
}
