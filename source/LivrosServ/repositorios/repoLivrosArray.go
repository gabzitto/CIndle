package repositorio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type RepositorioLivroArray struct {
	livros []Livro
}

func IniciarRepoLivroArray() *RepositorioLivroArray {
	repo := RepositorioLivroArray{
		livros: []Livro{},
	}
	dados, _ := ioutil.ReadFile("dados.txt")
	json.Unmarshal(dados, &repo.livros)
	return &repo
}

func (r *RepositorioLivroArray) Insert(li Livro) {
	defer r.persistir()
	for i := range r.livros {
		if r.livros[i].Nome == "" {
			r.livros[i] = li
			return
		}
	}
	r.livros = append(r.livros, li)
}

func (r *RepositorioLivroArray) Delete(nome string) error {
	defer r.persistir()
	for i := range r.livros {
		if r.livros[i].Nome == nome {
			r.livros[i] = Livro{}
			return nil
		}
	}
	return fmt.Errorf("livro nao encontrado")
}

func (r *RepositorioLivroArray) Read(nome string) (Livro, error) {
	defer r.persistir()
	for i := range r.livros {
		if r.livros[i].Nome == nome {
			return r.livros[i], nil
		}
	}
	return Livro{}, fmt.Errorf("Livro nao encontrado")
}

func (r *RepositorioLivroArray) Get() []Livro {
	return r.livros
}

func (r *RepositorioLivroArray) persistir() {
	dados, _ := json.Marshal(r.livros)
	ioutil.WriteFile("dados.txt", dados, 0660)
}
