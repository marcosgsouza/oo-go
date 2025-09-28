package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	// aulaUm()
	// aulaDois()
	//aulaTres()
	//aulaQuatro()
	aulaCinco()
}

// func aulaUm() {
// 	// Aula - 1
// 	contaDoGuilherme := contas.ContaCorrente{Titular: "Guilherme",
// 		NumeroAgencia: 589, NumeroConta: 123456, Saldo: 125.5}

// 	//contaDaBruna := contas.ContaCorrente{"Bruna", 222, 111222, 200}

// 	contaDaBruna := contas.ContaCorrente{}
// 	contaDaBruna.Titular = "Bruna"
// 	contaDaBruna.NumeroAgencia = 222
// 	contaDaBruna.NumeroConta = 111222
// 	contaDaBruna.Saldo = 200

// 	fmt.Println(contaDoGuilherme)
// 	fmt.Println(contaDaBruna)

// 	var contaDaCris *contas.ContaCorrente
// 	contaDaCris = new(contas.ContaCorrente)
// 	contaDaCris.Titular = "Cris"
// 	contaDaCris.Saldo = 500

// 	fmt.Println(*contaDaCris)
// }

// func aulaDois() {
// 	// Aula - 2
// 	contaDoMarcos := contas.ContaCorrente{}
// 	contaDoMarcos.Titular = "Marcos"
// 	contaDoMarcos.Saldo = 500.

// 	fmt.Println(contaDoMarcos.Saldo)
// 	fmt.Println(contaDoMarcos.Sacar(250))
// 	fmt.Println(contaDoMarcos.Saldo)
// }

// func aulaTres() {
// Aula - 3
// contaDoMarcos := contas.ContaCorrente{}
// contaDoMarcos.Titular = "Marcos"
// contaDoMarcos.Depositar(500)
// fmt.Println("Saldo inicial:", contaDoMarcos.Saldo)

// fmt.Println(contaDoMarcos.Sacar(250))
// fmt.Println("Saldo da Conta:", contaDoMarcos.Saldo)

// status, valor := contaDoMarcos.Depositar(2000)
// fmt.Println(status, valor)

//------------------------------------------------------------------------------

// contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
// contaDaThiago := contas.ContaCorrente{Titular: "Thiago", Saldo: 100}

// transf := contaDaThiago.Transferencia(200, &contaDaSilvia)

// fmt.Println(transf)
// fmt.Println(contaDaSilvia)
// fmt.Println(contaDaThiago)
// }

func aulaQuatro() {
	contaDaBruna := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "bruna",
		Cpf:       "1445555",
		Profissao: "contadora",
	}, NumeroAgencia: 111, NumeroConta: 554255}
	contaDaBruna.Depositar(500)

	fmt.Println(contaDaBruna)

	clienteBruno := clientes.Titular{Nome: "Bruno", Cpf: "33336665", Profissao: "Desenvolvedor"}
	contaDoBruno := contas.ContaCorrente{Titular: clienteBruno, NumeroAgencia: 333, NumeroConta: 3565444}
	contaDoBruno.Depositar(1000)

	fmt.Println(contaDoBruno)
}

type verificaConta interface {
	Sacar(valor float64) string
}

func aulaCinco() {

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	pagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.Obtersaldo())

	contaDoMarcos := contas.ContaCorrente{}
	contaDoMarcos.Depositar(10000)
	pagarBoleto(&contaDoMarcos, 500)

	fmt.Println(contaDoMarcos.Obtersaldo())
}

func pagarBoleto(conta verificaConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}
