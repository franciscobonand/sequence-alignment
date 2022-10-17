# Alinhamento de Sequências

## Execução

Primeiramente, é necessário instalar a linguagem Golang. Para isso, consultar o [passo-a-passo oficial de instalação](https://go.dev/doc/install)

Em seguida, vá até o local no qual esse projeto se encotra em seu computador pelo terminal. É necessário estar na raiz do projeto, onde o arquivo `main.go` se encontra.
Com isso, basta executar o programa com o comando `go run .`. As seguintes maneiras de execução são válidas:

- `go run . simple <match> <mismatch> <gap> <sequence1> <sequence2>`
  - Executa o programa utilizando os valores de pontuação fornecidos como entrada.
  - **Os valores de match, mismatch e gap devem ser números inteiros**
- `go run . blosum <gap> <sequence1> <sequence2>`
  - Executa o programa utilizando os valores de pontuação com base na matriz Blosum62.
  - **O valore de gap deve ser um número inteiro**

## Múltiplas Sequências

Esse projeto é capaz de lidar com o alinhamento de múltiplas sequências, basta informar as sequências adicionais como entrada na execução do programa (`... <sequence1> <sequence2> <sequence3> ...`).

### Algoritmo utilizado

Primeiramente é gerada uma pontuação para cada sequência fornecida, tendo como base o complemento da distância de Hamming (número de caracteres iguais na mesma posição), de forma a ordenar essas sequências da maior pontuação para a menor.
Uma vez que as sequências estejam ordenadas, utiliza-se o alinhamento par-a-par progressivamente, da seguinte maneira:

- Primeira sequência é alinhada com a segunda
- Segunda sequência é alinhada com a terceira
- Terceira sequência é alinhada com a quarta, e assim por diante

Esse algoritmo foi empregado dado que sua complexidade O(k^2n^2), onde k é o número de caracteres por sequência e n o número de sequências, é inferior a um algoritmo baseado em programação dinâmica, o qual possui complexidade geométrica associada ao número de sequências dadas como entrada.

### Ressalvas

- Esse alinhamento não é ótimo, dado que erros em um alinhamento inicial não serão corrigidos em outras etapas de alinhamento.
- Para sequências muito diferentes, ocorre o problema de as sequências "alinhadas" resultantes apresentarem uma pequena diferença no número de caracteres.
  - O ideal seria que, após o alinhamento de um par de sequências, a forma alinhada da sequência fosse utilizada no próximo alinhamento progressivo. No entanto, houveram dúvidas a respeito da melhor maneira de alinhar uma sequência previamente alinhada, a qual contém gaps ("-"), com uma sequência não-alinhada.

[Repositório do GitHub do projeto](https://github.com/franciscobonand/sequence-alignment)
