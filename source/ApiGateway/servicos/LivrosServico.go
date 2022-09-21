package servicos

import (
	"fmt"
	"net/http"
)

type LivrosServ struct {
	caminho string
	client  *http.Client
}

func InicializarLivrosServ(caminho string) *LivrosServ {
	return &LivrosServ{caminho: caminho,
		client: http.DefaultClient}
}

func (re *LivrosServ) ListarLivros(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get(re.caminho + "/livros")
	responderRequisicao(resp, w)
}

func (re *LivrosServ) AdicionarLivro(w http.ResponseWriter, r *http.Request) {

	req, err := http.NewRequest("POST", re.caminho+"/cadastrarlivros", r.Body)
	fmt.Println("adicionando livro")
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := re.client.Do(req)
	responderRequisicao(resp, w)
}

func (re *LivrosServ) PaginaAdicionarLivro(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", re.caminho+"/cadastrarlivros", nil)

	if err != nil {
		return
	}
	resp, _ := re.client.Do(req)
	responderRequisicao(resp, w)
}
