package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/LucasPaulo001/Echo-Surf/internal/browser"
)


func main() {
	url := flag.String("url", "", "URL da página a ser acessada (https://example.com).")

	headers := flag.Bool("headers", false, "Exibir headers HTTP")

	flag.Parse()

	if *url == "" {
		fmt.Println("Uso: echosurf --url https://example.com [--save] [--headers]")
		return
	}

	fmt.Print("\nEcho Surf v0.1 - Navegador de Linha de Comando\n")

	fmt.Println("Acessando: ", *url)
	page, err := browser.LoadPage(*url)
	
	if err != nil {
		log.Fatalf("Erro ao carregar página: %v", err)
	}

	fmt.Printf("Status: %d\n", page.StatusCode)
	fmt.Printf("Título: %s\n", page.Title)
	fmt.Printf("Links encontrados: %d\n", page.LinkCount)
	fmt.Printf("Imagens: %d\n", page.ImagesCount)

	if *headers {
		fmt.Print("\n---- Headers HTTP ----\n")
		for k, vals := range page.Headers {
			for _, v := range vals {
				fmt.Printf("%s: %s\n", k, v)
			}
		}
	}
}