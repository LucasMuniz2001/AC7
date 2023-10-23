package main

import (
	"fmt"
)

// Definindo a estrutura do nó da lista circular
type No struct {
	Data int
	Proximo *No
}

// Função para exibir os nós em uma lista circular
func exibeListaCircular(inicio *Node) {
	if inicio == nil {
		fmt.Println("A lista circular está vazia.")
		return
	}

	current := inicio
	for {
		fmt.Printf("%d -> ", current.Data)
		atual = atual.Proximo
		if atual == inicio {
			break
		}
	}
	fmt.Println()
}

// Função para inserir um nó no início da lista circular
func insereInicio(inicio *No, data int) *No {
	novoNo := &No{Data: data}
	if start == nil {
		novoNo.Proximo = novoNo
		return novoNo
	}

	novoNo.Proximo = inicio
	atual := inicio
	for current.Proximo != start {
		atual = atual.Proximo
	}
	atual.Proximo = novoNo

	return novoNo
}

// Função para excluir o primeiro nó da lista circular
func exlcuiPrimeiroNo(inicio *No) *No {
	if inicio == nil {
		fmt.Println("A lista circular está vazia.")
		return nil
	}

	if inicio.Proximo == inicio {
		fmt.Println("Único nó na lista foi removido.")
		return nil
	}

	atual := inicio
	for atual.Proximo != inicio {
		atual = atual.Proximo
	}
	atual.Proximo = inicio.Proximo
	fmt.Println("Primeiro nó removido.")
	return inicio.Proximo
}

func main() {
	var inicio *No

	// Inserindo alguns nós no início da lista circular
	inicio = insereNoComeço(inicio, 10)
	inicio = insereNoComeço(inicio, 20)
	inicio = insereNoComeço(inicio, 30)

	// Exibindo a lista circular
	fmt.Println("Lista Circular:")
	exibeListaCircular(inicio)

	// Excluindo o primeiro nó
	inicio = excluiPrimeiroNoinicio)

	// Exibindo a lista após a exclusão
	fmt.Println("Lista após a exclusão:")
	exibeListaCircular(inicio)
}
