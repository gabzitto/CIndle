package apiExterna

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type API struct {
	Livros map[string]Detalhes
}

type Detalhes struct {
	Descricao string `json:"Descricao"`
	Autor     string `json:"Autor"`
	Genero    string `json:"Genero"`
}

func StartAPI() {
	router := httprouter.New()
	a := API{
		Livros: make(map[string]Detalhes),
	}
	router.GET("/livros/:livro", a.DetalhesLivro)

	a.Livros["Harry Potter e a Pedra Filosofal"] = Detalhes{Autor: "J.K. Rowling", Genero: "Fantasia", Descricao: "Harry adentra um mundo mágico que jamais imaginara, vivendo diversas aventuras com seus novos amigos, Rony Weasley e Hermione Granger."}
	a.Livros["A história secreta"] = Detalhes{Autor: "Donna Tartt", Genero: "Suspense", Descricao: "Um sofisticado grupo de alunos de grego resolve reproduzir as orgias dionisíacas da Antiguidade. Uma obra cujos temas são cumplicidade e decepção, inocência e corrupção moral, responsabilidade e culpa."}
	a.Livros["Clean Code"] = Detalhes{Autor: "Robert Cecil Martin", Genero: "Tecnologia", Descricao: "O renomado especialista em software, Robert C. Martin, apresenta um paradigma revolucionário com Código limpo: Habilidades Práticas do Agile Software."}
	sv := &http.Server{
		Addr: "localhost:8001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      router,
	}

	log.Fatal(sv.ListenAndServe())
}

func (a *API) DetalhesLivro(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(ps.ByName("livro"))
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(a.Livros[ps.ByName("livro")])
	w.Write((data))
}
