package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

// Aula - 2
func (c *ContaCorrente) Sacar(saque float64) string {
	podeSacar := saque > 0 && saque <= c.saldo
	if podeSacar {
		c.saldo -= saque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

// Aula - 3
func (c *ContaCorrente) Depositar(deposito float64) (string, float64) {
	if deposito > 0 {
		c.saldo += deposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do depósito menor que zero", c.saldo
	}
}

func main() {
	// Aula - 1
	// contaDoGuilherme := ContaCorrente{titular: "Guilherme",
	// 	numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	// contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	// fmt.Println(contaDoGuilherme)
	// fmt.Println(contaDaBruna)

	// var contaDaCris *ContaCorrente
	// contaDaCris = new(ContaCorrente)
	// contaDaCris.titular = "Cris"
	// contaDaCris.saldo = 500

	// fmt.Println(*contaDaCris)

	// Aula - 2
	// contaDoMarcos := ContaCorrente{}
	// contaDoMarcos.titular = "Marcos"
	// contaDoMarcos.saldo = 500.

	// fmt.Println(contaDoMarcos.saldo)
	// fmt.Println(contaDoMarcos.Sacar(250))
	// fmt.Println(contaDoMarcos.saldo)

	// Aula - 3
	contaDoMarcos := ContaCorrente{}
	contaDoMarcos.titular = "Marcos"
	contaDoMarcos.Depositar(500)
	fmt.Println("Saldo inicial:", contaDoMarcos.saldo)

	fmt.Println(contaDoMarcos.Sacar(250))
	fmt.Println("Saldo da Conta:", contaDoMarcos.saldo)

	status, valor := contaDoMarcos.Depositar(2000)
	fmt.Println(status, valor)
}
