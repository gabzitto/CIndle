package controlador

import (
	colecao "CindleResenhasServ/colecoes"
	repositorio "CindleResenhasServ/repositorios"
)

type ControladorResenhas interface {
	ListarResenha() []repositorio.Resenha
	CadastrarResenha(repositorio.Resenha) error
	DeletarResenha(repositorio.Resenha) error
}

type ControladorResenhasImp struct {
	ColecaoResenhas colecao.ColecaoResenhas
}

func InicializarControladorResenhas(cfg string) *ControladorResenhasImp {
	return &ControladorResenhasImp{
		ColecaoResenhas: colecao.InicializarColecaoResenhas(cfg),
	}
}

func (c *ControladorResenhasImp) ListarResenha() []repositorio.Resenha {
	return c.ColecaoResenhas.PegarResenhas()
}

func (c *ControladorResenhasImp) CadastrarResenha(re repositorio.Resenha) error {
	return c.ColecaoResenhas.CriarResenha(re)
}

func (c *ControladorResenhasImp) DeletarResenha(re repositorio.Resenha) error {
	return c.ColecaoResenhas.ExcluirResenha(re)
}
