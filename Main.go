package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// Tabuleiro representa um tabuleiro de jogo da velha
type Tabuleiro [3][3]int

// Variaveis para implementar a regra principal do game
var ultimoSubLinha int = -1
var ultimoSubColuna int = -1

// NovoTabuleiro cria e retorna um novo tabuleiro vazio
func NovoTabuleiro() Tabuleiro {
	return Tabuleiro{}
}

// TabuleiroUltimate representa o tabuleiro do Ultimate Tic Tac Toe
type TabuleiroUltimate [3][3]Tabuleiro

// NovoTabuleiroUltimate cria e retorna um novo tabuleiro Ultimate vazio
func NovoTabuleiroUltimate() TabuleiroUltimate {
	var tabuleiroUltimate TabuleiroUltimate
	for i := range tabuleiroUltimate {
		for j := range tabuleiroUltimate[i] {
			tabuleiroUltimate[i][j] = NovoTabuleiro()
		}
	}
	return tabuleiroUltimate
}

// ExibirTabuleiroUltimate exibe o tabuleiro Ultimate no console
func ExibirTabuleiroUltimate(tabuleiroUltimate TabuleiroUltimate) {
	for i := range tabuleiroUltimate {
		for subI := 0; subI < 3; subI++ {
			for j := range tabuleiroUltimate[i] {
				for subJ := 0; subJ < 3; subJ++ {
					valor := tabuleiroUltimate[i][j][subI][subJ]
					simbolo := "-"
					switch valor {
					case 1:
						simbolo = "X"
					case 2:
						simbolo = "O"
					case 3:
						simbolo = "C" // "C" representa coringa
					}
					fmt.Printf(" %s ", simbolo)
				}
				if j < 2 {
					fmt.Print("|")
				}
			}
			fmt.Println()
		}
		if i < 2 {
			fmt.Println("-----------------------------")
		}
	}
}

// JogarUltimate permite que os jogadores façam uma jogada no Ultimate Tic Tac Toe
func JogarUltimate(tabuleiroUltimate *TabuleiroUltimate, jogador, linha, coluna, subLinha, subColuna int) bool {
	// Se for a primeira jogada ou a jogada estiver no quadrado correto
	if (ultimoSubLinha == -1 && ultimoSubColuna == -1) || (linha == ultimoSubLinha && coluna == ultimoSubColuna) {
		if (*tabuleiroUltimate)[linha][coluna][subLinha][subColuna] == 0 || (*tabuleiroUltimate)[linha][coluna][subLinha][subColuna] == 3 {
			(*tabuleiroUltimate)[linha][coluna][subLinha][subColuna] = jogador
			// Atualiza a última jogada
			ultimoSubLinha = subLinha
			ultimoSubColuna = subColuna
			return true
		}
	}
	return false
}

// VerificarVitoriaU verifica se há um vencedor no tic tac toe menor
func VerificarVitoriaU(tabuleiro *Tabuleiro) int {
	// Verifica linhas e colunas para encontrar um vencedor
	for i := 0; i < 3; i++ {
		if tabuleiro[i][0] != 0 && (tabuleiro[i][0] == tabuleiro[i][1] || tabuleiro[i][1] == 3) && (tabuleiro[i][1] == tabuleiro[i][2] || tabuleiro[i][2] == 3) {
			return tabuleiro[i][0]
		}
		if tabuleiro[0][i] != 0 && (tabuleiro[0][i] == tabuleiro[1][i] || tabuleiro[1][i] == 3) && (tabuleiro[1][i] == tabuleiro[2][i] || tabuleiro[2][i] == 3) {
			return tabuleiro[0][i]
		}
	}
	// Verifica diagonais para encontrar um vencedor
	if tabuleiro[0][0] != 0 && (tabuleiro[0][0] == tabuleiro[1][1] || tabuleiro[1][1] == 3) && (tabuleiro[1][1] == tabuleiro[2][2] || tabuleiro[2][2] == 3) {
		return tabuleiro[0][0]
	}
	if tabuleiro[0][2] != 0 && (tabuleiro[0][2] == tabuleiro[1][1] || tabuleiro[1][1] == 3) && (tabuleiro[1][1] == tabuleiro[2][0] || tabuleiro[2][0] == 3) {
		return tabuleiro[0][2]
	}
	// Nenhum vencedor encontrado
	return 0
}

// VerificarVitoriaUltimate verifica se há um vencedor no Ultimate Tic Tac Toe
func VerificarVitoriaUltimate(tabuleiroUltimate *TabuleiroUltimate) int {
	// Verifica cada tabuleiro menor para um vencedor
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			vencedor := VerificarVitoriaU(&tabuleiroUltimate[i][j])
			if vencedor != 0 {
				// Preenche o tabuleiro menor com o número do vencedor
				for subI := 0; subI < 3; subI++ {
					for subJ := 0; subJ < 3; subJ++ {
						tabuleiroUltimate[i][j][subI][subJ] = vencedor
					}
				}
			} else if QuadradoMenorEmpatado(&tabuleiroUltimate[i][j]) {
				// Marca o quadrado menor como coringa apenas se não houver um vencedor
				for subI := 0; subI < 3; subI++ {
					for subJ := 0; subJ < 3; subJ++ {
						tabuleiroUltimate[i][j][subI][subJ] = 3 // 3 representa coringa
					}
				}
			}
		}
	}
	// Verifica o tabuleiro Ultimate para um vencedor
	var tabuleiroVencedor Tabuleiro
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tabuleiroVencedor[i][j] = VerificarVitoriaU(&tabuleiroUltimate[i][j])
		}
	}
	return VerificarVitoriaU(&tabuleiroVencedor)
}

// Verificar se um quadrado menor está empatado
func QuadradoMenorEmpatado(tabuleiro *Tabuleiro) bool {
	// Verifica se todas as células estão preenchidas e não há vencedor
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (*tabuleiro)[i][j] == 0 {
				return false
			}
		}
	}
	return VerificarVitoriaU(tabuleiro) == 0
}

// Verificar se um quadrado menor está completo
func QuadradoMenorCompleto(tabuleiroUltimate *TabuleiroUltimate, subLinha, subColuna int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (*tabuleiroUltimate)[subLinha][subColuna][i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tabuleiroUltimate := NovoTabuleiroUltimate()

	var jogador1 string
	var jogador2 string

	fmt.Print("Digite o nome do Jogador 1 (X): ")
	fmt.Scan(&jogador1)

	fmt.Print("Digite o nome do Jogador 2 (O): ")
	fmt.Scan(&jogador2)

	jogadorAtual := rand.Intn(2) + 1
	nomeAtual := jogador1
	if jogadorAtual == 1 {
		nomeAtual = jogador1
	} else {
		nomeAtual = jogador2
	}

	for {
		ExibirTabuleiroUltimate(tabuleiroUltimate)
		if ultimoSubLinha != -1 && ultimoSubColuna != -1 {
			fmt.Printf("%s (Jogador %d), entre com a sublinha e subcoluna (ex: 2 2): ", nomeAtual, jogadorAtual)
		} else {
			fmt.Printf("%s (Jogador %d), entre com a linha, coluna, sublinha e subcoluna (ex: 1 1 2 2): ", nomeAtual, jogadorAtual)
		}

		scanner.Scan()
		entrada := scanner.Text()
		var posicoes []string

		if ultimoSubLinha != -1 && ultimoSubColuna != -1 {
			posicoes = append(posicoes, strconv.Itoa(ultimoSubLinha+1), strconv.Itoa(ultimoSubColuna+1))
			entradaSub := strings.Split(entrada, " ")
			posicoes = append(posicoes, entradaSub...)
		} else {
			posicoes = strings.Split(entrada, " ")
		}

		// Verifica se a entrada tem o número correto de partes
		if len(posicoes) != 4 {
			fmt.Println("Entrada inválida. Por favor, insira os números corretamente.(Se nn inseriu nada ainda, desconsidere a mensagem e continue o jogo!)")
			continue
		}

		linha, errLinha := strconv.Atoi(posicoes[0])
		coluna, errColuna := strconv.Atoi(posicoes[1])
		subLinha, errSubLinha := strconv.Atoi(posicoes[2])
		subColuna, errSubColuna := strconv.Atoi(posicoes[3])

		// Verifica se houve erro na conversão de algum dos números
		if errLinha != nil || errColuna != nil || errSubLinha != nil || errSubColuna != nil {
			fmt.Println("Entrada inválida. Por favor, insira apenas números.")
			continue
		}
		if JogarUltimate(&tabuleiroUltimate, jogadorAtual, linha-1, coluna-1, subLinha-1, subColuna-1) {
			// Atualiza a última jogada com a sublinha e subcoluna da jogada atual
			ultimoSubLinha = subLinha - 1
			ultimoSubColuna = subColuna - 1

			vencedor := VerificarVitoriaUltimate(&tabuleiroUltimate)
			if vencedor != 0 {
				fmt.Printf("%s venceu o Ultimate Tic Tac Toe!\n", nomeAtual)
				break
			}
			// Se o quadrado menor estiver completo ou o jogador vencer, permita escolher outro quadrado
			if QuadradoMenorCompleto(&tabuleiroUltimate, ultimoSubLinha, ultimoSubColuna) || vencedor != 0 {
				ultimoSubLinha = -1
				ultimoSubColuna = -1
			}
		} else {
			fmt.Println("Posição já ocupada, inválida ou não está no quadrado correto, tente novamente.")
		}

		// Atualiza o jogadorAtual e nomeAtual para o próximo jogador
		if jogadorAtual == 1 {
			jogadorAtual = 2
			nomeAtual = jogador2
		} else {
			jogadorAtual = 1
			nomeAtual = jogador1
		}
	}
}
