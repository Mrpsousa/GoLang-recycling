package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] {
		req, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error getting: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error reading: %v\n", err)
		}
		var data ViaCEP

		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error unmarshaling: %v\n", err)
		}

		fmt.Println(data)

		file, err := os.Create("dados_cidade.txt")
		if err != nil {
			fmt.Fprint(os.Stderr, "Error creating file: %v\n", err)
		}

		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))
	}
}
