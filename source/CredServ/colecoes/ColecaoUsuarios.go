package colecao

import (
	repositorio "CindleUserServ/repositorios"
	"fmt"
)

type ColecaoUsuarios interface {
	ValidarCredenciais(repositorio.Usuario) (string, error)
	CriarCredenciais(repositorio.Usuario) error
	MudarStatusAssinatura(string) error
	PuxarUsuario(string) (repositorio.Usuario, error)
}

type ColecaoUsuariosImp struct {
	RepositorioUsuarios *repositorio.RepoUsuarios
}

func InicializarColecaoUsuarios(cfg string) *ColecaoUsuariosImp {
	return &ColecaoUsuariosImp{
		RepositorioUsuarios: repositorio.InicializarRepositorioUsuario(cfg),
	}
}

func (c *ColecaoUsuariosImp) PuxarUsuario(email string) (repositorio.Usuario, error) {
	u, err := c.RepositorioUsuarios.RepoUsuarios.Read(email)
	if err == nil {
		return u, err
	}
	return repositorio.Usuario{}, err
}

func (c *ColecaoUsuariosImp) ValidarCredenciais(user repositorio.Usuario) (string, error) {
	u, err := c.RepositorioUsuarios.RepoUsuarios.Read(user.Email)
	if err != nil {
		fmt.Println("nao achou ngm!")
		return "", err
	}
	if u.Senha != user.Senha {
		fmt.Println("senha errada!")
		return "", fmt.Errorf("senha incorreta")
	}
	return u.Email, nil
}

func (c *ColecaoUsuariosImp) CriarCredenciais(user repositorio.Usuario) error {
	_, err := c.RepositorioUsuarios.RepoUsuarios.Read(user.Email)
	if err == nil {
		fmt.Println("ja tem")
		return fmt.Errorf("Credenciais ja cadastradsa")
	}
	c.RepositorioUsuarios.RepoUsuarios.Insert(user)
	return nil
}

func (c *ColecaoUsuariosImp) MudarStatusAssinatura(email string) error {
	fmt.Println("email do assinante", email)
	fmt.Println("Todo mundo", c.RepositorioUsuarios.RepoUsuarios.Get())
	u, err := c.RepositorioUsuarios.RepoUsuarios.Read(email)
	if err != nil {
		fmt.Println("nao achou ngm!")
		return err
	}
	u.Assinante = true
	c.RepositorioUsuarios.RepoUsuarios.Update(u.Email, u)
	return nil
}
