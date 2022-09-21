package colecao

import (
	repositorio "CindleResenhasServ/repositorios"
	"fmt"
)

type ColecaoResenhas interface {
	CriarResenha(repositorio.Resenha) error
	ExcluirResenha(repositorio.Resenha) error
	PegarResenhas() []repositorio.Resenha
}

type ColecaoResenhasImp struct {
	RepositorioResenhas repositorio.RepositorioResenha
}

func InicializarColecaoResenhas(cfg string) *ColecaoResenhasImp {
	return &ColecaoResenhasImp{
		RepositorioResenhas: repositorio.InicializarRepositorioResenhas(cfg),
	}
}

func (c *ColecaoResenhasImp) CriarResenha(re repositorio.Resenha) error {
	_, err := c.RepositorioResenhas.Read(re.NomeLivro, re.Usuario)
	if err == nil {
		return fmt.Errorf("Resenha ja existente")
	}
	c.RepositorioResenhas.Insert(re)
	return nil
}

func (c *ColecaoResenhasImp) ExcluirResenha(re repositorio.Resenha) error {
	return c.RepositorioResenhas.Delete(re.NomeLivro, re.Usuario)
}

func (c *ColecaoResenhasImp) PegarResenhas() []repositorio.Resenha {
	return c.RepositorioResenhas.Get()
}
