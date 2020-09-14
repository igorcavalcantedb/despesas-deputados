package model

type Deputado struct {
	Dados struct {
		ID           int    `json:"id"`
		URI          string `json:"uri"`
		NomeCivil    string `json:"nomeCivil"`
		UltimoStatus struct {
			ID            int    `json:"id"`
			URI           string `json:"uri"`
			Nome          string `json:"nome"`
			SiglaPartido  string `json:"siglaPartido"`
			URIPartido    string `json:"uriPartido"`
			SiglaUf       string `json:"siglaUf"`
			IDLegislatura int    `json:"idLegislatura"`
			URLFoto       string `json:"urlFoto"`
			Email         string `json:"email"`
			Data          string `json:"data"`
			NomeEleitoral string `json:"nomeEleitoral"`
			Gabinete      struct {
				Nome     string `json:"nome"`
				Predio   string `json:"predio"`
				Sala     string `json:"sala"`
				Andar    string `json:"andar"`
				Telefone string `json:"telefone"`
				Email    string `json:"email"`
			} `json:"gabinete"`
			Situacao          string      `json:"situacao"`
			CondicaoEleitoral string      `json:"condicaoEleitoral"`
			DescricaoStatus   interface{} `json:"descricaoStatus"`
		} `json:"ultimoStatus"`
		Cpf                 string      `json:"cpf"`
		Sexo                string      `json:"sexo"`
		URLWebsite          interface{} `json:"urlWebsite"`
		RedeSocial          []string    `json:"redeSocial"`
		DataNascimento      string      `json:"dataNascimento"`
		DataFalecimento     interface{} `json:"dataFalecimento"`
		UfNascimento        string      `json:"ufNascimento"`
		MunicipioNascimento string      `json:"municipioNascimento"`
		Escolaridade        string      `json:"escolaridade"`
	} `json:"dados"`
	Links []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
}
