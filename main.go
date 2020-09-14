package main

import (
	"despesas-deputados/controller"
	"despesas-deputados/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("TEste")
	//var api model.Deputado

	imprimirInf(get())
	obterGastos()

}
func get() model.Deputado {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("https://dadosabertos.camara.leg.br/api/v2/deputados/160592")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	//bodyString := string(bodyBytes)
	//fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var todoStruct model.Deputado
	json.Unmarshal(bodyBytes, &todoStruct)
	//fmt.Printf("API Response as struct %+v\n", todoStruct)
	return todoStruct
}

func imprimirInf(deputado model.Deputado) {
	fmt.Println("Nome:" + deputado.Dados.NomeCivil)
	fmt.Println(" -:" + deputado.Dados.UltimoStatus.Nome)
	fmt.Println("Partido:" + deputado.Dados.UltimoStatus.SiglaPartido)

}
func obterGastos() {
	fmt.Println("Total : R$", controller.Get())

}
