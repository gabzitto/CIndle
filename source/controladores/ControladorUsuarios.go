package controlador

import (
	colecao "Cindle/colecoes"
)

type ControladorUsuarios interface {
	EfetuarLogin(string, string) (string, error)
	EfetuarCadastro(string, string, string) error
	mudarStatusAssinatura(string, colecao.Cartao) error
}

type ControladorUsuariosImp struct {
	ColecaoUsuarios colecao.ColecaoUsuarios
}

func InicializarControladorUsuarios(cfg string) *ControladorUsuariosImp {
	return &ControladorUsuariosImp{
		ColecaoUsuarios: colecao.InicializarColecaoUsuarios(cfg),
	}
}

func (c *ControladorUsuariosImp) EfetuarLogin(email, senha string) (string, error) {
	return c.ColecaoUsuarios.ValidarCredenciais(email, senha)
}

func (c *ControladorUsuariosImp) EfetuarCadastro(email, senha, nome string) error {
	return c.ColecaoUsuarios.CriarCredenciais(email, senha, nome)
}

func (c *ControladorUsuariosImp) mudarStatusAssinatura(email string, ct colecao.Cartao) error {
	return nil
}
