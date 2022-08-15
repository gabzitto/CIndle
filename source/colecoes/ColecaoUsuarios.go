package colecao

import (
	repositorio "Cindle/repositorios"
	"fmt"
)

type ColecaoUsuarios interface {
	ValidarCredenciais(string, string) (string, error)
	CriarCredenciais(string, string, string) error
	mudarStatusAssinatura(string, Cartao) error
}

type Cartao struct {
	numero   string
	validade string
	cvv      string
}

type ColecaoUsuariosImp struct {
	RepositorioUsuarios *repositorio.RepoUsuarios
}

func InicializarColecaoUsuarios(cfg string) *ColecaoUsuariosImp {
	return &ColecaoUsuariosImp{
		RepositorioUsuarios: repositorio.InicializarRepositorioUsuario(cfg),
	}
}

func (c *ColecaoUsuariosImp) ValidarCredenciais(email, senha string) (string, error) {
	u, err := c.RepositorioUsuarios.RepoUsuarios.Read(email)
	if err != nil {
		fmt.Println("nao achou ngm!")
		return "", err
	}
	if u.Senha != senha {
		fmt.Println("senha errada!")
		return "", fmt.Errorf("senha incorreta")
	}
	return u.Nome, nil
}

func (c *ColecaoUsuariosImp) CriarCredenciais(email, senha, nome string) error {
	user := repositorio.Usuario{
		Nome:  nome,
		Email: email,
		Senha: senha,
	}
	err := c.RepositorioUsuarios.RepoUsuarios.Upsert(user)
	if err != nil {
		return err
	}
	return nil
}

func (c *ColecaoUsuariosImp) mudarStatusAssinatura(email string, ct Cartao) error {
	return nil
}
