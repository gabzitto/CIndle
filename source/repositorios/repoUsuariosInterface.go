package repositorio

type RepositorioUsuario interface {
	insert(Usuario)
	Update(string, Usuario) error
	Read(string) (Usuario, error)
	Delete(string) (Usuario, error)
	Upsert(Usuario) error
}

type Usuario struct {
	Email string
	Senha string
	Nome  string
	Tipo  string
}
