package CindleLivrosServ

import (
	controlador "CindleLivrosServ/controladores"
	repositorio "CindleLivrosServ/repositorios"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Cindle struct {
	ControladorLivros controlador.ControladorLivros
}

func (c *Cindle) ListarLivros() []repositorio.Livro {
	return c.ControladorLivros.AcessarLivros()
}

func (c *Cindle) CadastrarLivro(li repositorio.Livro) error {
	err := c.ControladorLivros.ValidarLivro(li)
	if err != nil {
		fmt.Println("raelly")
		return err
	}
	if li.Autor == "" || li.Descricao == "" || li.Genero == "" {
		livroDetalhado := c.ControladorLivros.BuscarDetalhes(li.Nome)
		if li.Autor == "" {
			li.Autor = livroDetalhado.Autor
		}
		if li.Descricao == "" {
			li.Descricao = livroDetalhado.Descricao
		}
		if li.Genero == "" {
			li.Genero = livroDetalhado.Genero
		}
		return c.ControladorLivros.CadastrarLivro(livroDetalhado)
	}
	return c.ControladorLivros.CadastrarLivro(li)
}

type Server struct {
	sv              *http.Server
	TipoRepositorio string `yaml:"tipo-repositorio"`
	c               *Cindle
}

/*func (s *Server) LerArquivoConfiguracao() error {
	yamlFile, err := ioutil.ReadFile("cindle.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, s)
	if err != nil {
		return err
	}
	fmt.Println(s)
	return nil
}*/

func StartServer() {
	s := &Server{}
	/*err := s.LerArquivoConfiguracao()
	if err != nil {
		panic(err)
	}*/

	s.TipoRepositorio = "array"

	s.c = &Cindle{
		ControladorLivros: controlador.InicializarControladorLivros(s.TipoRepositorio),
	}

	http.HandleFunc("/livros", s.ControladorTelaLivro)
	http.HandleFunc("/cadastrarlivros", s.ControladorTelaCadastroLivro)
	http.HandleFunc("/listarLivros", s.ControladorListagemLivros)

	sv := &http.Server{
		Addr: ":8003",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(sv.ListenAndServe())
}

type DadosLivros struct {
	Livros []repositorio.Livro
}

func (s *Server) ControladorListagemLivros(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		livros := s.c.ListarLivros()
		livrosString := ""
		for _, l := range livros {
			livrosString += l.Nome
			livrosString += "#"
		}
		w.Write([]byte(livrosString))
	}
}

func (s *Server) ControladorTelaCadastroLivro(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("Cadastrado")
		err := r.ParseForm()
		if err != nil {
			log.Println("Falha no form")
			return
		}
		fmt.Println("form", r.Form)
		li := repositorio.Livro{
			Nome:      r.Form.Get("livro"),
			Genero:    r.Form.Get("genero"),
			Descricao: r.Form.Get("descricao"),
			Autor:     r.Form.Get("autor"),
			Link:      r.Form.Get("link"),
		}
		s.c.CadastrarLivro(li)
	}
	p := "./cadastrarLivro.html"
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)
}

func (s *Server) ControladorTelaLivro(w http.ResponseWriter, r *http.Request) {
	dados := DadosLivros{
		Livros: s.c.ListarLivros(),
	}
	fmt.Println("servindo")
	fmt.Println("dados", dados)
	t, _ := template.ParseFiles("./livros.html")
	t.Execute(w, dados)
}
