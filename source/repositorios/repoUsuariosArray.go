package repositorio

import (
	"fmt"
)

type RepositorioUsuarioArray struct {
	usuarios []Usuario
}

func IniciarRepoUsuarioArray() *RepositorioUsuarioArray {
	repo := RepositorioUsuarioArray{
		usuarios: []Usuario{},
	}
	return &repo
}

func (r *RepositorioUsuarioArray) insert(User Usuario) {
	for i := range r.usuarios {
		if r.usuarios[i].Email != "" {
			r.usuarios[i] = User
			return
		}
	}
	r.usuarios = append(r.usuarios, User)
}

func (r *RepositorioUsuarioArray) Delete(email string) (Usuario, error) {
	for i := range r.usuarios {
		if r.usuarios[i].Email == email {
			tmp := r.usuarios[i]
			r.usuarios[i] = Usuario{}
			return tmp, nil
		}
	}
	return Usuario{}, fmt.Errorf("usuario nao encontrado")
}

func (r *RepositorioUsuarioArray) Read(email string) (Usuario, error) {
	for i := range r.usuarios {
		if r.usuarios[i].Email == email {
			return r.usuarios[i], nil
		}
	}
	return Usuario{}, fmt.Errorf("usuario nao encontrado")
}

func (r *RepositorioUsuarioArray) Update(email string, User Usuario) error {
	for i := range r.usuarios {
		if r.usuarios[i].Email == email {
			r.usuarios[i] = User
			return nil
		}
	}
	return fmt.Errorf("usuario nao encontrado")
}

func (r *RepositorioUsuarioArray) Upsert(User Usuario) error {
	for i := range r.usuarios {
		if r.usuarios[i].Email == User.Email {
			return r.Update(User.Email, User)
		}
	}
	r.insert(User)
	return nil
}
