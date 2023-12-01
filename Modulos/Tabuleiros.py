import tkinter as tk

class Tabuleiro:
    def __init__(self, master):
        self.master = master
        self.master.title("Jogo da Velha")
        self.frame = tk.Frame(self.master)
        self.frame.pack()

        self.tabuleiro_estado = [[' ' for _ in range(3)] for _ in range(3)]
        self.jogador_atual = 'X'  # Começa com o jogador X

        # Criando os 9 quadrados grandes
        for i in range(3):
            for j in range(3):
                quadrado = tk.Frame(self.frame, width=200, height=200, bg="black")
                quadrado.grid(row=i, column=j, padx=10, pady=10)

                # Criando os 9 quadrados menores dentro de cada quadrado grande
                for a in range(3):
                    for b in range(3):
                        quadrado_menor = tk.Label(quadrado, text='', width=8, height=4, bg="white", relief=tk.RIDGE)
                        quadrado_menor.grid(row=a, column=b, padx=2, pady=2)
                        quadrado_menor.bind("<Button-1>", lambda event, info=quadrado_menor.info: self.jogar(info))

    def jogar(self, info):
        linha, coluna, sub_linha, sub_coluna = info['linha'], info['coluna'], info['sub_linha'], info['sub_coluna']

        if self.tabuleiro_estado[linha][coluna] == ' ':
            self.tabuleiro_estado[linha][coluna] = self.jogador_atual
            info['text'] = self.jogador_atual

            # Verificar se o jogador atual venceu
            if self.verificar_vitoria(self.jogador_atual):
                print(f"Jogador {self.jogador_atual} venceu!")

            # Alternar para o próximo jogador
            self.alternar_jogador()

    def verificar_vitoria(self, jogador):
        # Verifica se o jogador atual venceu em algum tabuleiro menor
        # Adicione a lógica para verificar linhas, colunas e diagonais

        # Exemplo de lógica para verificar uma linha
        for linha in self.tabuleiro_estado:
            if all(celula == jogador for celula in linha):
                return True

        return False

    def verificar_empate(self):
        # Verifica se todos os tabuleiros menores estão preenchidos
        for linha in self.tabuleiro_estado:
            if any(celula == ' ' for celula in linha):
                return False  # Ainda há pelo menos um espaço vazio

        return True  # Todos os tabuleiros menores estão preenchidos

    def alternar_jogador(self):
        self.jogador_atual = 'O' if self.jogador_atual == 'X' else 'X'

    def verificar_estado(self):
        # Adicione aqui a lógica para verificar o estado do tabuleiro principal (quem venceu, se há empate, etc.)
        pass

def main():
    root = tk.Tk()
    jogo = Tabuleiro(root)
    root.mainloop()

if __name__ == "__main__":
    main()
