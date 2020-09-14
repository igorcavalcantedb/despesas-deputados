package controller

import (
	"despesas-deputados/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Get() float64 {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("https://dadosabertos.camara.leg.br/api/v2/deputados/160592/despesas")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	//bodyString := string(bodyBytes)
	//fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var todoStruct model.Despesas
	json.Unmarshal(bodyBytes, &todoStruct)
	//fmt.Printf("API Response as struct %+v\n", todoStruct)
	var somaTotal float64
	for _, despesa := range todoStruct.Dados {
		fmt.Println("Mes: ", despesa.Mes, "Valor:", despesa.ValorDocumento)
		somaTotal += despesa.ValorDocumento

	}
	return somaTotal
}
