package main

import "net"
import "fmt"
import "bufio"
import "strings"

func main() {
	fmt.Println("Subindo o server...")

	// Escuta em qualquer interface
	ln, _ := net.Listen("tcp", ":12345")

	// Aceita conexão na porta definida
	conexao, _ := ln.Accept()

	// Fica rodando até que seja executado ctrl+c
	for {
		// Aguarda mensagem
		messagem, _ := bufio.NewReader(conexao).ReadString('\n')

		// Imprime a mensagem
		fmt.Print("Mensagem recebida:", string(messagem))

		novamensagem := strings.ToUpper(messagem)

		// Deixa o cliente mandar outra mensagem
		conexao.Write([]byte(novamensagem + "\n"))
	}
}
