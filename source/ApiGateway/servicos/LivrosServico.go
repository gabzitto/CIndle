package servicos

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

func (re *LivrosServ) VerificarExistenciaLivro(nomeLivro string) bool {
	if nomeLivro == "" {
		return false
	}
	resp, _ := http.Get(re.caminho + "/listarLivros")
	bytes, _ := ioutil.ReadAll(resp.Body)
	respString := strings.Split(string(bytes), "#")
	for _, livro := range respString {
		if livro == nomeLivro {
			return true
		}
	}
	return false
}

func (re *LivrosServ) AdicionarLivro(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	req, err := http.NewRequest("POST", re.caminho+"/cadastrarlivros", bytes.NewReader(body))
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
