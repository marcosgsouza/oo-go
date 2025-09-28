package contas

import cliente "banco/clientes"

type ContaCorrente struct {
	Titular       cliente.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(saque float64) string {
	podeSacar := saque > 0 && saque <= c.saldo
	if podeSacar {
		c.saldo -= saque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(deposito float64) (string, float64) {
	if deposito > 0 {
		c.saldo += deposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do depósito menor que zero", c.saldo
	}
}

func (c *ContaCorrente) Transferencia(valorTransferencia float64, conta *ContaCorrente) bool {
	if valorTransferencia > 0 && valorTransferencia <= c.saldo {
		c.saldo -= valorTransferencia
		conta.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) Obtersaldo() float64 {
	return c.saldo
}
