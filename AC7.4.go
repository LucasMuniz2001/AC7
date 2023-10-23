package main

import "fmt"


type No struct {
	Data     int
	Anterior *Node
	Proximo  *Node
}


type ListaDuplamenteEncadeada struct {
	Primeiro *No
	Ultimo   *No
}

// Função para exibir os nós em uma lista duplamente encadeada
func (lista *ListaDuplamenteEncadeada) Exibir() {
	if lista.Primeiro == nil {
		fmt.Println("A lista está vazia.")
		return
	}

	nodoAtual := lista.Primeiro
	for noAtual != nil {
		fmt.Printf("%d <-> ", nodoAtual.Data)
		noAtual = noAtual.Proximo
	}
	fmt.Println()
}

// Função para buscar um nó na lista
func (lista *ListaDuplamenteEncadeada) Buscar(valor int) *No {
	noAtual := lista.Primeiro
	for noAtual != nil {
		if noAtual.Data == valor {
			return noAtual
		}
		noAtual = noAtual.Proximo
	}
	return nil // Nó não encontrado
}

// Função para inserir um nó no início da lista duplamente encadeada
func (lista *ListaDuplamenteEncadeada) InsereNoInicio(valor int) {
	novoNo := &No{Data: valor}

	if lista.Primeiro == nil {
		lista.Primeiro = novoNodo
		lista.Ultimo = novoNodo
	} else {
		novoNodo.Proximo = lista.Primeiro
		lista.Primeiro.Anterior = novoNodo
		lista.Primeiro = novoNodo
	}
}

// Função para remover um nó da lista
func (lista *ListaDuplamenteEncadeada) Remover(no *No) {
	if no == nil {
		fmt.Println("Nó não encontrado.")
		return
	}

	if no == lista.Primeiro {
		lista.Primeiro = no.Proximo
	} else {
		no.Anterior.Proximo = no.Proximo
	}

	if no == lista.Ultimo {
		lista.Ultimo = no.Anterior
	} else {
		no.Proximo.Anterior = no.Anterior
	}

	fmt.Println("Nó removido:", no.Data)
}

func main() {
	lista := ListaDuplamenteEncadeada{}

	lista.InsereNoInicio(30)
	lista.InsereNoInicio(20)
	lista.InsereNoInicio(10)

	lista.Exibir()

	valorBuscado := 20
	noBuscado := lista.Buscar(valorBuscado)
	if noBuscado != nil {
		fmt.Printf("Nó com valor %d encontrado.\n", valorBuscado)
	} else {
		fmt.Printf("Nó com valor %d não encontrado.\n", valorBuscado)
	}

	valorRemovido := 20
	noRemover := lista.Buscar(valorRemovido)
	lista.Remover(noARemover)

	lista.Exibir()
}
