package repositorio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type RepositorioUsuarioArray struct {
	usuarios []Usuario
}

func IniciarRepoUsuarioArray() *RepositorioUsuarioArray {
	repo := RepositorioUsuarioArray{
		usuarios: []Usuario{},
	}
	dados, _ := ioutil.ReadFile("dados.txt")
	json.Unmarshal(dados, &repo.usuarios)
	return &repo
}

func (r *RepositorioUsuarioArray) Insert(User Usuario) {
	defer r.persistir()
	for i := range r.usuarios {
		if r.usuarios[i].Email == "" {
			r.usuarios[i] = User
			return
		}
	}
	r.usuarios = append(r.usuarios, User)
}

func (r *RepositorioUsuarioArray) Delete(email string) (Usuario, error) {
	defer r.persistir()
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
	defer r.persistir()
	for i := range r.usuarios {
		if r.usuarios[i].Email == email {
			r.persistir()
			return r.usuarios[i], nil
		}
	}
	r.persistir()
	return Usuario{}, fmt.Errorf("usuario nao encontrado")

}

func (r *RepositorioUsuarioArray) Update(email string, User Usuario) error {
	defer r.persistir()
	for i := range r.usuarios {
		if r.usuarios[i].Email == email {
			r.usuarios[i] = User
			r.persistir()
			return nil
		}
	}
	r.persistir()
	return fmt.Errorf("usuario nao encontrado")
}

func (r *RepositorioUsuarioArray) Upsert(User Usuario) error {
	defer r.persistir()
	for i := range r.usuarios {
		if r.usuarios[i].Email == User.Email {
			fmt.Println("atualizando")
			return r.Update(User.Email, User)
		}
	}
	fmt.Println("inserindo")
	r.Insert(User)
	return nil
}

func (r *RepositorioUsuarioArray) Get() []Usuario {
	return r.usuarios
}

func (r *RepositorioUsuarioArray) GetP() *[]Usuario {
	return &r.usuarios
}

func (r *RepositorioUsuarioArray) persistir() {
	dados, _ := json.Marshal(r.usuarios)
	ioutil.WriteFile("dados.txt", dados, 0660)
}
