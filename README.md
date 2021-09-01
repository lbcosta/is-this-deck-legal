# Programação Script 2021.1 - AP2

- **Imagine um programa de computador e desenvolva um prototipo para ele com alguma linguagem script.**
  > Explicação do programa na sessão **Sobre a Aplicação**
- **Crie um documento respondendo as seguintes perguntas sobre o protótipo:**
  - **O seu protótipo é descartável ou evolucionário? Justifique.**
    <br>
    Descartável. A lógica inicial da aplicação está no protótipo. Porém, em geral, não é tão amigável para o usuário/público alvo usar a linha de comando. É póssivel criar uma página HTML com a linguagem Go, mas esse tipo de aplicação se beneficiaria melhor de uma linguagem com mais recursos para construção de interface de usuário, como, por exemplo, o Javascript e uma de suas bibliotecas de Front-End, como o React.
  - **Quais as características da linguagem escolhida fazem ela apropriada para o desenvolvimento de um protótipo?**
    <br>
    A linguagem é construída de uma forma popular entre desenvolvedores (C like), é simples, com tipagem simples, com suporte à API externa usada para os própositos da aplicação e com forma simples de leitura de arquivos, requisições HTTP e interação com linha de comando.
  - **Que dificuldades você teve no desenvolvimento do protótipo?**
    <br>
    Por ser a primeira vez usando a linguagem Go, tive a dificuldade natural que desenvolvedores tem em usar uma nova linguagem, como sintaxe (uso de parênteses, ponto e vírgula etc) e aprender a usar as bibliotecas padrões da linguagem.
  - **Quais as vantagens de ter um protótipo desenvolvido antes da implementação do programa efetivo?**
    <br>
    Validar a ideia principal da aplicação antes de pensar sobre os detalhes. Um protótipo ajuda os desenvolvedores a entenderem melhor se a lógica principal da aplicação realmente funciona, que tipos de erros podem acontecer, que tipos de APIs e recursos vão ser essênciais. Possuindo um protótipo que valide a ideia, a implementação do programa efetivo acontece mais rápida e com menos erros ao longo do caminho.

## Sobre a Aplicação - _Is This Deck Legal?_

Magic: The Gathering (MTG), um jogo de cartas extremamente popular no mundo inteiro, possui várias maneiras diferentes de ser jogado usando as mesmas regras básicas. Essas maneiras são chamadas de _formatos_ e são caracterizados, principalmente, por quais coleções de cartas podem ou não ser usadas no formato.

Muitas vezes, jogadores de MTG que já jogam algum formato há algum tempo, tendem a experimentar outros formatos de jogo. Nesse momento, a transição de um formato para outro se torna muito mais fácil quando é possível usar as mesmas cartas de um formato em outro. Fazendo-se, assim, desnecessário adquirir muitas cartas novas ou decks/baralhos inteiros novos.

_Is This Deck Legal_ é uma ferramenta que ajuda os jogadores a identificar quais cartas de um deck são legais em um determinado formato. Ele funciona da seguinte maneira: O programa lê um arquivo de texto com a lista de cartas separadas por linha, você escolhe o formato para checar a legalidade das cartas e, então, o programa consulta a [API de Magic: The Gathering](https://docs.magicthegathering.io/) e te retorna uma lista de quais cartas são válidas e quais cartas não são válidas no formato escolhido.

## Sobre o repositório e como executar a aplicação

A aplicação requer um único argumento na linha de comando: o nome do arquivo que possui a lista de cartas a serem checadas.

Primeiramente é necessário fazer o build da aplicação:

```bash
go build main.go
```

Depois é só executar passando o nome do arquivo (que deve estar na mesma pasta do arquivo main)

```bash
./main deck.csv
```

Um exemplo de input para o programa está dentro do repositório. O arquivo **deck.csv**. Um exemplo de output para esse arquivo seria:

![example-output](https://github.com/lbcosta/is-this-deck-legal/raw/main/output-example.png)
