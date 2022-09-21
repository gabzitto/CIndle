package colecao

import (
	repositorio "CindleLivrosServ/repositorios"
	"fmt"
)

type ColecaoLivros interface {
	CriarLivro(repositorio.Livro) error
	ExcluirLivro(string) error
	PegarLivros() []repositorio.Livro
}

type ColecaoLivrosImp struct {
	RepositorioLivros repositorio.RepositorioLivro
}

func InicializarColecaoLivros(cfg string) *ColecaoLivrosImp {
	return &ColecaoLivrosImp{
		RepositorioLivros: repositorio.InicializarRepositorioLivros(cfg),
	}
}

func (c *ColecaoLivrosImp) CriarLivro(li repositorio.Livro) error {
	_, err := c.RepositorioLivros.Read(li.Nome)
	if err == nil {
		return fmt.Errorf("Livro ja existente")
	}
	fmt.Println("inserindo livro")
	c.RepositorioLivros.Insert(li)
	return nil
}

func (c *ColecaoLivrosImp) ExcluirLivro(livro string) error {
	return c.RepositorioLivros.Delete(livro)
}

func (c *ColecaoLivrosImp) PegarLivros() []repositorio.Livro {
	return c.RepositorioLivros.Get()
}
