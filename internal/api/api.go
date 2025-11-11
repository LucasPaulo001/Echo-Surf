package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func TestAPI(urlAPI, method, body, headersAPI string, jsonFormat bool) error {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	var reqBody io.Reader
	if body != "" {
		reqBody = bytes.NewBuffer([]byte(body))
	}

	req, err := http.NewRequest(method, urlAPI, reqBody)
	if err != nil {
		return fmt.Errorf("Erro ao criar requisição: %v", err)
	}

	if headersAPI != "" {
		for _, h := range strings.Split(headersAPI, ";") {
			pair := strings.SplitN(h, ":", 2)
			if len(pair) == 2 {
				req.Header.Set(strings.TrimSpace(pair[0]), strings.TrimSpace(pair[1]))
			}
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Erro ao executar requisição: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("\nStatus: %d %s\n", resp.StatusCode, resp.Status)
	fmt.Println("Headers:")
	for k, v := range resp.Header {
		fmt.Printf(" %s: %s\n", k, strings.Join(v, ", "))
	}

	fmt.Println("\nCorpo da resposta:")
	bodyBytes, _ := io.ReadAll(resp.Body)

	if jsonFormat {
		var pretty bytes.Buffer
		err := json.Indent(&pretty, bodyBytes, "", " ")
		if err != nil {
			fmt.Println(pretty.String())
			return nil
		}
	}

	fmt.Println(string(bodyBytes))
	return nil
}