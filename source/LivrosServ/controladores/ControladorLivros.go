package controlador

import (
	"CindleLivrosServ/adaptador"
	colecao "CindleLivrosServ/colecoes"
	repositorio "CindleLivrosServ/repositorios"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ControladorLivros interface {
	AcessarLivros() []repositorio.Livro
	CadastrarLivro(repositorio.Livro) error
	ValidarLivro(repositorio.Livro) error
	BuscarDetalhes(string) repositorio.Livro
	DeletarLivro(string) error
}

type ControladorLivrosImp struct {
	ColecaoLivros colecao.ColecaoLivros
	client        *http.Client
}

func InicializarControladorLivros(cfg string) *ControladorLivrosImp {
	return &ControladorLivrosImp{
		ColecaoLivros: colecao.InicializarColecaoLivros(cfg),
		client:        http.DefaultClient,
	}
}

func (c *ControladorLivrosImp) AcessarLivros() []repositorio.Livro {
	return c.ColecaoLivros.PegarLivros()
}

func (c *ControladorLivrosImp) BuscarDetalhes(nomeLivro string) repositorio.Livro {
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:8001/livros/%s", nomeLivro), nil)
	resp, err := c.client.Do(req)
	if err != nil {
		return repositorio.Livro{Nome: nomeLivro}
	}
	data, _ := ioutil.ReadAll(resp.Body)
	ada := adaptador.Detalhes{}
	json.Unmarshal(data, &ada)
	fmt.Println(ada.Genero)
	return repositorio.Livro{
		Nome:      nomeLivro,
		Genero:    ada.Genero,
		Autor:     ada.Autor,
		Descricao: ada.Descricao,
	}
}

func (c *ControladorLivrosImp) CadastrarLivro(li repositorio.Livro) error {
	return c.ColecaoLivros.CriarLivro(li)
}

func (c *ControladorLivrosImp) ValidarLivro(li repositorio.Livro) error {
	if len(li.Nome) < 2 {
		return fmt.Errorf("Livro invalido")
	}
	return nil
}

func (c *ControladorLivrosImp) DeletarLivro(nome string) error {
	return c.ColecaoLivros.ExcluirLivro(nome)
}
