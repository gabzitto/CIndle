package autorizacao

import (
	"fmt"
)

type UsuariosLogados struct {
	logados map[string]string // token - nome de usuario
}

func InicializarControladorUsuariosLogados() *UsuariosLogados {
	return &UsuariosLogados{
		logados: make(map[string]string),
	}
}

func (u *UsuariosLogados) PegarUsuario(token string) (string, error) {
	if user, ok := u.logados[token]; ok {
		return user, nil
	}
	return "", fmt.Errorf("Usuario nao logado")
}

func (u *UsuariosLogados) AdicionarUsuario(token, user string) {
	u.logados[token] = user
}
