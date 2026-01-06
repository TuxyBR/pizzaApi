package models

type Teste struct {
	IDOc                int     `json:"id_oc"`
	IDCot               int     `json:"id_cot"`
	IDMat               int     `json:"id_mat"`
	IDObra              int     `json:"id_obra"`
	DescricaoDaOc       *string `json:"descricao_da_oc"`
	DescricaoDoMaterial *string `json:"descricao_do_material"`
	Atividade           *string `json:"atividade"`
	Quantidade          float64 `json:"quantidade"`
	Unidade             *string `json:"unidade"`
	Fornecedor          *string `json:"fornecedor"`
	PrecoUnitario       float64 `json:"preco_unitario"`
	ValorTotal          float64 `json:"valor_total"`
	DataDeEntrega       *string `json:"data_de_entrega"`
	TipoPgmt            *string `json:"tipo_pgmt"`
	CondPgmt            *string `json:"cond_pgmt"`
	DadosPgmt           *string `json:"dados_pgmt"`
	ObsPgmt             *string `json:"obs_pgmt"`
	Observacoes         *string `json:"observacoes"`
	Autorizador         *string `json:"autorizador"`
	DataAprovada        *string `json:"data_aprovada"`
	DataDaOc            *string `json:"data_da_oc"`
	NomeCancelamento    *string `json:"nome_cancelamento"`
	DataCancelamento    *string `json:"data_cancelamento"`
	NomeCriador         *string `json:"nome_criador"`
	Status              *string `json:"status"`
}
