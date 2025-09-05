package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("123456789012345678901234567890", 10) // 30-значное число
	b.SetString("987654321098765432109876543210", 10) // 30-значное число

	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n", b.String())
	fmt.Println()

	result := new(big.Int)

	result.Add(a, b)
	fmt.Printf("Сложение: a + b = %s\n", result.String())

	result.Sub(a, b)
	fmt.Printf("Вычитание: a - b = %s\n", result.String())

	result.Mul(a, b)
	fmt.Printf("Умножение: a * b = %s\n", result.String())

	if b.Sign() != 0 { // проверка, что b не ноль
		quotient := new(big.Int)
		remainder := new(big.Int)

		quotient.DivMod(a, b, remainder)
		fmt.Printf("Деление: a / b = %s (остаток %s)\n",
			quotient.String(), remainder.String())
	} else {
		fmt.Println("Деление на ноль невозможно")
	}

	exponent := big.NewInt(3)
	result.Exp(a, exponent, nil)
	fmt.Printf("Степень: a^3 = %s\n", result.String())
}
