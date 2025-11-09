package media

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

)


func DownloadMedia(url, format string) (error) {
	format = strings.ToLower(format)

	// define opções do yt-dlp
	var args []string
	if format == "mp3" {
		args = []string{"-x", "--audio-format", "mp3", url}
	} else if format == "mp4" {
		args = []string{"-f", "mp4", url}
	} else {
		return fmt.Errorf("formato não suportado: %s", format)
	}

	fmt.Println("Iniciando download com yt-dlp...")
	cmd := exec.Command("yt-dlp", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("erro ao executar yt-dlp: %w", err)
	}

	fmt.Println("Download concluído com sucesso!")
	return nil
}