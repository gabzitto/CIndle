package servicos

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ResenhasServ struct {
	caminho string
	client  *http.Client
}

func InicializarResenhasServ(caminho string) *ResenhasServ {
	return &ResenhasServ{caminho: caminho,
		client: http.DefaultClient}
}

func (re *ResenhasServ) ListarResenhas(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get(re.caminho + "/resenha")
	responderRequisicao(resp, w)
}

func (re *ResenhasServ) AdicionarResenha(user string, w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadAll(r.Body)
	payload := fmt.Sprintf("usuario=%s&%s", user, string(bytes))

	req, err := http.NewRequest("POST", re.caminho+"/resenha", strings.NewReader(payload))

	if err != nil {
		return
	}
	token, err := r.Cookie("ApiToken")
	userEmail, err1 := r.Cookie("Username")
	if err != nil || err1 != nil {
		fmt.Println("Cookie nao presente")
		return
	}
	req.Header.Set("Cookie", fmt.Sprintf("%s; %s", userEmail, token))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := re.client.Do(req)
	responderRequisicao(resp, w)
}
