package repositorio

type RepoLivros struct {
	RepoLivros RepositorioLivro
}

func InicializarRepositorioLivros(cfg string) RepositorioLivro {
	if cfg == "array" {
		return IniciarRepoLivroArray()
	} else {

		return IniciarRepoLivroArray()
	}
}
