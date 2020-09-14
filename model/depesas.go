package model

type Despesas struct {
	Dados []struct {
		Ano               int     `json:"ano"`
		Mes               int     `json:"mes"`
		TipoDespesa       string  `json:"tipoDespesa"`
		CodDocumento      int     `json:"codDocumento"`
		TipoDocumento     string  `json:"tipoDocumento"`
		CodTipoDocumento  int     `json:"codTipoDocumento"`
		DataDocumento     string  `json:"dataDocumento"`
		NumDocumento      string  `json:"numDocumento"`
		ValorDocumento    float64 `json:"valorDocumento"`
		URLDocumento      string  `json:"urlDocumento"`
		NomeFornecedor    string  `json:"nomeFornecedor"`
		CnpjCpfFornecedor string  `json:"cnpjCpfFornecedor"`
		ValorLiquido      float64 `json:"valorLiquido"`
		ValorGlosa        float64 `json:"valorGlosa"`
		NumRessarcimento  string  `json:"numRessarcimento"`
		CodLote           int     `json:"codLote"`
		Parcela           int     `json:"parcela"`
	} `json:"dados"`
	Links []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
}
