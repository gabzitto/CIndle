package repositorio

type RepoUsuarios struct {
	RepoUsuarios RepositorioUsuario
}

func InicializarRepositorioUsuario(cfg string) *RepoUsuarios {
	rp := RepoUsuarios{}
	if cfg == "array" {
		rp.RepoUsuarios = IniciarRepoUsuarioArray()
	} else {
		//rp.RepoUsuarios = IniciarRepoUsuarioMap()
	}
	return &rp
}
