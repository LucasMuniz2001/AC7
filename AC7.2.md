# Lista Duplamente Encadeada Ordenada

Uma lista duplamente encadeada ordenada é uma estrutura de dados em que os nós contêm um valor e estão organizados de forma ordenada (por exemplo, em ordem crescente ou decrescente). Cada nó possui referências tanto para o nó anterior quanto para o próximo nó na lista.

## Estrutura de Nó

Cada nó em uma lista duplamente encadeada deve conter três campos:
- Dado: para armazenar o valor do nó.
- Anterior: para apontar para o nó anterior.
- Próximo: para apontar para o próximo nó.

### Exibição de um Nó

Procedimento ExibirListaDuplamenteEncadeada(Início):
    NóAtual = Início
    Enquanto NóAtual não é Nulo:
        Exibir NóAtual.Dado
        NóAtual = NóAtual.Próximo

#### Busca de um Nó

Função BuscarNó(Início, Valor):
    NóAtual = Início
    Enquanto NóAtual não é Nulo:
        Se NóAtual.Dado é igual a Valor:
            Retornar NóAtual
        NóAtual = NóAtual.Próximo
    Retornar Nulo

    
#### Inserção de um Nó

Procedimento InserirNóOrdenado(Início, NovoNó):
    Se Início é Nulo:
        Início = NovoNó
    Senão Se NovoNó.Dado < Início.Dado:
        NovoNó.Próximo = Início
        Início.Anterior = NovoNó
        Início = NovoNó
    Senão:
        NóAtual = Início
        Enquanto NóAtual.Próximo não é Nulo E NovoNó.Dado > NóAtual.Dado:
            NóAtual = NóAtual.Próximo
        NovoNó.Próximo = NóAtual
        NovoNó.Anterior = NóAtual.Anterior
        NóAtual.Anterior.Próximo = NovoNó
        NóAtual.Anterior = NovoNó

##### Remoção de um Nó

Procedimento RemoverNó(Início, Valor):
    NóParaRemover = BuscarNó(Início, Valor)
    Se NóParaRemover é Nulo:
        Retornar  // O nó não foi encontrado
    Se NóParaRemover.Anterior não é Nulo:
        NóParaRemover.Anterior.Próximo = NóParaRemover.Próximo
    Se NóParaRemover.Próximo não é Nulo:
        NóParaRemover.Próximo.Anterior = NóParaRemover.Anterior
    Se NóParaRemover é igual a Início:
        Início = NóParaRemover.Próximo
    Liberar NóParaRemover
