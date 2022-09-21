package repositorio

type RepositorioLivro interface {
	Insert(Livro)
	Get() []Livro
	Read(string) (Livro, error)
	Delete(string) error
}

type Livro struct {
	Nome      string
	Genero    string
	Descricao string
	Autor     string
	Link      string
}
