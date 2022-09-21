package repositorio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type RepositorioResenhaArray struct {
	resenhas []Resenha
}

func IniciarRepoResenhaArray() *RepositorioResenhaArray {
	repo := RepositorioResenhaArray{
		resenhas: []Resenha{},
	}
	dados, _ := ioutil.ReadFile("dados.txt")
	json.Unmarshal(dados, &repo.resenhas)
	return &repo
}

func (r *RepositorioResenhaArray) Insert(li Resenha) {
	defer r.persistir()
	for i := range r.resenhas {
		if r.resenhas[i].Usuario == "" {
			r.resenhas[i] = li
			return
		}
	}
	r.resenhas = append(r.resenhas, li)
}

func (r *RepositorioResenhaArray) Delete(nomeLivro, usuario string) error {
	defer r.persistir()
	for i := range r.resenhas {
		if r.resenhas[i].NomeLivro == nomeLivro && r.resenhas[i].Usuario == usuario {
			r.resenhas[i] = Resenha{}
			return nil
		}
	}
	return fmt.Errorf("Resenha nao encontrado")
}

func (r *RepositorioResenhaArray) Read(nomeLivro, usuario string) (Resenha, error) {
	defer r.persistir()
	for i := range r.resenhas {
		if r.resenhas[i].NomeLivro == nomeLivro && r.resenhas[i].Usuario == usuario {
			return r.resenhas[i], nil
		}
	}
	return Resenha{}, fmt.Errorf("Resenha nao encontrada")
}

func (r *RepositorioResenhaArray) Get() []Resenha {
	return r.resenhas
}

func (r *RepositorioResenhaArray) persistir() {
	dados, _ := json.Marshal(r.resenhas)
	ioutil.WriteFile("dados.txt", dados, 0660)
}
