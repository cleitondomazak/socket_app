package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

	// conecta no socket
	conexao, _ := net.Dial("tcp", "127.0.0.1:12345")

	for {
		// Le a mensagem inserida
		leitor := bufio.NewReader(os.Stdin)
		fmt.Print("Mensagem: ")
		texto, _ := leitor.ReadString('\n')

		// envia mensagem para o servidor
		fmt.Fprintf(conexao, texto+"\n")

		messagem, _ := bufio.NewReader(conexao).ReadString('\n')
		fmt.Print("Mensagem enviada para o servidor: " + messagem)
	}
}
