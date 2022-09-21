package repositorio

type RepositorioUsuario interface {
	Insert(Usuario)
	Update(string, Usuario) error
	Read(string) (Usuario, error)
	Delete(string) (Usuario, error)
	Upsert(Usuario) error
	Get() []Usuario
}

type Usuario struct {
	Email     string
	Senha     string
	Nome      string
	Assinante bool
	Admin     bool
}
