
# Lista Circular 

Uma lista circular é uma estrutura de dados onde o último nó da lista aponta de volta para o primeiro nó, formando um ciclo.

## Exibição dos Nós em uma Lista Circular

Procedimento ExibirListaCircular(Início):
    Se Início é Nulo:
        Retornar

   NóAtual = Início
    Faça:
        Exibir NóAtual.Dado
        NóAtual = NóAtual.Próximo
    Enquanto NóAtual não é igual a Início

### Inserção de um Nó no início
Procedimento InserirNoInicio(Início, Valor):
    NovoNó = CriarNó(Valor)
    Se Início é Nulo:
        Início = NovoNó
        NovoNó.Próximo = Início
    Senão:
        NovoNó.Próximo = Início.Próximo
        Início.Próximo = NovoNó

#### Exclusão do primeiro nó
Procedimento ExcluirPrimeiroNó(Início):
    Se Início é Nulo:
        Retornar

   Se Início.Próximo é igual a Início:
        // A lista só possui um nó
        Liberar Início
        Início = Nulo
    Senão:
        NóASerExcluído = Início.Próximo
        Início.Próximo = NóASerExcluído.Próximo
        Liberar NóASerExcluído
