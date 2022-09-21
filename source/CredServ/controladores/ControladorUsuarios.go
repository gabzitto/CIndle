package controlador

import (
	"CindleUserServ/autorizacao"
	colecao "CindleUserServ/colecoes"
	repositorio "CindleUserServ/repositorios"
	"fmt"
	"math/rand"
	"time"
)

type ControladorUsuarios interface {
	EfetuarLogin(repositorio.Usuario) (string, string, error)
	EfetuarCadastro(repositorio.Usuario) error
	MudarStatusAssinatura(string) error
	VerificarStatusUsuario(string, string) (bool, bool)
}

type ControladorUsuariosImp struct {
	Autorizacao     *autorizacao.UsuariosLogados
	ColecaoUsuarios colecao.ColecaoUsuarios
}

func InicializarControladorUsuarios(cfg string) *ControladorUsuariosImp {
	return &ControladorUsuariosImp{
		Autorizacao:     autorizacao.InicializarControladorUsuariosLogados(),
		ColecaoUsuarios: colecao.InicializarColecaoUsuarios(cfg),
	}
}

func (c *ControladorUsuariosImp) EfetuarLogin(user repositorio.Usuario) (string, string, error) {
	nomeUsuario, err := c.ColecaoUsuarios.ValidarCredenciais(user)
	if err != nil {
		return "", "", err
	}
	token := c.GerarToken()
	c.Autorizacao.AdicionarUsuario(token, user.Email)
	return nomeUsuario, token, nil
}

func (c *ControladorUsuariosImp) EfetuarCadastro(user repositorio.Usuario) error {
	return c.ColecaoUsuarios.CriarCredenciais(user)
}

func (c *ControladorUsuariosImp) MudarStatusAssinatura(email string) error {
	fmt.Println(c.ColecaoUsuarios.PuxarUsuario("caio@caio.caio"))
	return c.ColecaoUsuarios.MudarStatusAssinatura(email)

}

func (c *ControladorUsuariosImp) VerificarStatusUsuario(email, token string) (bool, bool) {
	u, err := c.ColecaoUsuarios.PuxarUsuario(email)
	if err == nil {
		return u.Assinante, u.Admin
	}
	return false, false
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (c *ControladorUsuariosImp) GerarToken() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 15)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
