package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/LucasPaulo001/Echo-Surf/internal/browser"
	"github.com/LucasPaulo001/Echo-Surf/internal/media"
	"github.com/fatih/color"
)


func main() {
	// Cores
	title := color.New(color.FgCyan, color.Bold).SprintFunc()
	info := color.New(color.FgGreen).SprintFunc()
	section := color.New(color.FgYellow, color.Bold).SprintFunc()
	error := color.New(color.FgRed, color.Bold).SprintFunc()

	url := flag.String("url", "", "URL da página a ser acessada (https://example.com).")
	headers := flag.Bool("headers", false, "Exibir headers HTTP")
	linksOnly := flag.Bool("links", false, "Mostrar apenas links da página")
	imagesOnly := flag.Bool("images", false, "Mostrar apenas as imagens")
	downloadYt := flag.String("download", "", "URL do vídeo para download")
	help := flag.Bool("help", false, "Exibir ajuda do CLI")

	flag.Parse()

	if *help {
		fmt.Println(section("\n--- Ajunda ---"))
		fmt.Print("[--url <url> --links]: Lista os links de uma página web\n")
		fmt.Print("[--url <url>]: Lista os dados da página web\n")
		fmt.Print("[--url <url> --images]: Lista as imagens de uma página web\n")
		fmt.Print("[--url <url> --download (mp3 ou mp4)]: Faz download de mídia\n")

		return
	}

	if *url == "" {
		fmt.Println("Uso: echosurf --url https://example.com [--save] [--headers]")
		return
	}

	// Retorno de busca
	fmt.Print("\n" + title("Echo Surf v0.1 - Navegador de Linha de Comando\n\n"))

	fmt.Println("Acessando: ", info(*url))

	page, err := browser.LoadPage(*url)
	if err != nil {
		log.Fatalf("Erro ao carregar página: %v", err)
	}

	fmt.Print(section("\n--- Informações da Página ---\n"))
	fmt.Println("Status: ", info(page.StatusCode))
	fmt.Println("Título: ", info(page.Title))
	fmt.Printf("Links encontrados: %d\n", len(page.Links))
	fmt.Printf("Imagens: %d\n", len(page.Images))

	if *headers {
		fmt.Print(section("\n---- Headers HTTP ----\n"))
		for k, vals := range page.Headers {
			for _, v := range vals {
				fmt.Printf("%s: %s\n", k, v)
			}
		}
	}

	if *linksOnly {
		fmt.Println(section("\n--- Links ---"))
		for _, link := range page.Links {
			fmt.Println("-", link)
		}

		return
	}

	if *imagesOnly {
		fmt.Println(section("\n--- Imagens ---"))
		for _, image := range page.Images {
			fmt.Println("-", image)
		}

		return
	}

	if *downloadYt != "" {
		err := media.DownloadMedia(*url, *downloadYt)
		if err != nil {
			fmt.Print(error("Erro ao baixar mídia: ", err))
			os.Exit(1)
		}

		fmt.Println("Download concluído com sucesso.")
		return
	}
}