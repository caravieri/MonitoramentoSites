package main

import (
	"fmt"
	"net/http" //um pacote  para realizar requisições Get e Post.
	"os"
	"reflect"  //Pacote reflect para mostra o tipo da variavél usando TypeOf
)

func main() {

	
	exibeIntroducao()

	//comentarios
	// if
	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// }else if comando == 2{
	// 	fmt.Println("Exibindo Logs...")
	// }else if comando == 0{
	// 	fmt.Println("Saindo...")
	// }else{
	// 	fmt.Println("Não conheço este comando")
	// }
	////////////////////////////////////////////////////////
	//slice 
	// func exibeNomes() { 
	// 	nomes := []string {"Guilherme", "Brenner", "Brigida", "Igor", "Bruno"}
	// 	fmt.Println(nomes)
	// 	fmt.Println(reflect.TypeOf(nomes))	
	// 	fmt.Println("O meu Slice tem", len(nomes), "itens")
	// 	fmt.Println("O meu Slice tem capacidade para", cap(nomes), "itens")
	// 	nomes = append(nomes, "Juliano")
	// 	fmt.Println(nomes)
	// 	fmt.Println(reflect.TypeOf(nomes))	
	// 	fmt.Println("O meu Slice tem", len(nomes), "itens")
	// 	fmt.Println("O meu Slice tem capacidade para", cap(nomes), "itens")	
	// }
	
	for {
		exibeMenu()
		comando := leComando()
	
		// switch
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1) //exibe o status do erro
		}
	}
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func exibeIntroducao() {
	nome := "Guilherme"
	versao := 1.1

	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

// entender melhor
func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido) //%d = representa um inteiro
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	var sites [4] string //array
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://consorciomagalu.com.br/"
	sites[2] = "https://www.alura.com.br/"
	fmt.Println(sites)
	fmt.Println(reflect.TypeOf(sites))

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)
	
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "esta com problemas. Status Cod:", resp.StatusCode, '\n')
	}
}

	