package repositorio

import "fmt"

type RepositorioUsuarioMap struct {
	usuarios map[string]Usuario
}

func (r *RepositorioUsuarioMap) Insert(User Usuario) {
	r.usuarios[User.Email] = User
}

func IniciarRepoUsuarioMap() *RepositorioUsuarioMap {
	repo := RepositorioUsuarioMap{
		usuarios: make(map[string]Usuario),
	}
	return &repo
}

func (r *RepositorioUsuarioMap) Delete(email string) (Usuario, error) {
	user, exist := r.usuarios[email]
	if exist {
		r.usuarios[email] = Usuario{}
		return user, nil
	}
	return Usuario{}, fmt.Errorf("usuario nao encontrado")
}

func (r *RepositorioUsuarioMap) Read(email string) (Usuario, error) {
	user, exist := r.usuarios[email]
	if exist {
		return user, nil
	}
	return Usuario{}, fmt.Errorf("usuario nao encontrado")
}

func (r *RepositorioUsuarioMap) Update(email string, User Usuario) error {
	_, exist := r.usuarios[email]
	if exist {
		r.usuarios[email] = User
		return nil
	}
	return fmt.Errorf("usuario nao encontrado")
}

func (r *RepositorioUsuarioMap) Upsert(User Usuario) error {
	_, exist := r.usuarios[User.Email]
	if exist {
		return r.Update(User.Email, User)
	}
	r.Insert(User)
	return nil
}

func (r *RepositorioUsuarioMap) Get() map[string]Usuario {
	return r.usuarios
}
