package servicos

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type UsuariosServ struct {
	caminho string
	client  *http.Client
}

func InicializarUsuariosServ(caminho string) *UsuariosServ {
	return &UsuariosServ{caminho: caminho,
		client: http.DefaultClient}
}

func (u *UsuariosServ) AcessarPaginaLogin(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get(u.caminho + "/login")
	responderRequisicao(resp, w)
}

func (u *UsuariosServ) RealizarLogin(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Post(u.caminho+"/login", "application/x-www-form-urlencoded", r.Body)
	responderRequisicao(resp, w)
}

func (u *UsuariosServ) RealizarAssinatura(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	req, err := http.NewRequest("POST", u.caminho+"/assinatura", bytes.NewReader(body))
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
	resp, _ := u.client.Do(req)
	responderRequisicao(resp, w)
}

func (u *UsuariosServ) IrParaAssinatura(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", u.caminho+"/assinatura", nil)
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
	resp, _ := u.client.Do(req)
	responderRequisicao(resp, w)
}

func (u *UsuariosServ) RealizarCadastro(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Post(u.caminho+"/cadastro", "application/x-www-form-urlencoded", r.Body)
	responderRequisicao(resp, w)
}

func (u *UsuariosServ) ChecarPermissoes(email, token string, w http.ResponseWriter, r *http.Request) []string { //[assinante, admin]
	resp, _ := http.Post(u.caminho+"/verificarusuario", "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(fmt.Sprintf("%s,%s", email, token))))
	perm, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Error reading request")
	}
	return strings.Split(string(perm), ",")

}

func montarRequisicao(metodo, caminho, token, user string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(metodo, caminho, body)
	req.Header.Set("Cookie", fmt.Sprintf("Username=%s; ApiToken=%s", user, token))
	if metodo == "post" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	return req
}

func responderRequisicao(resp *http.Response, wr http.ResponseWriter) {
	wr.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	wr.Header().Set("Content-Length", resp.Header.Get("Content-Length"))

	io.Copy(wr, resp.Body)
	resp.Body.Close()
}
