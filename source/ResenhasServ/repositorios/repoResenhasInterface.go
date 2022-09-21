package repositorio

type RepositorioResenha interface {
	Insert(Resenha)
	Get() []Resenha
	Read(string, string) (Resenha, error)
	Delete(string, string) error
}

type Resenha struct {
	Usuario   string
	NomeLivro string
	Descricao string
}
