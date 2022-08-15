package cindle

import (
	controlador "Cindle/controladores"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"

	"gopkg.in/yaml.v3"
)

type Cindle struct {
	ControladorUsuarios controlador.ControladorUsuarios
}

func (c *Cindle) EfetuarLogin(email, senha string) (string, error) {
	return c.ControladorUsuarios.EfetuarLogin(email, senha)
}

func (c *Cindle) CadastrarUsuario(email, senha, nome string) error {
	return c.ControladorUsuarios.EfetuarCadastro(email, senha, nome)
}

type Server struct {
	sv              *http.Server
	TipoRepositorio string `yaml:"tipo-repositorio"`
	c               *Cindle
}

func (s *Server) LerArquivoConfiguracao() error {
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
}

func StartServer() {
	s := &Server{}
	err := s.LerArquivoConfiguracao()
	if err != nil {
		panic(err)
	}

	s.c = &Cindle{
		ControladorUsuarios: controlador.InicializarControladorUsuarios(s.TipoRepositorio),
	}

	http.HandleFunc("/login", s.ControladorTelaLogin)
	http.HandleFunc("/cadastro", s.ControladorTelaCadastro)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	sv := &http.Server{
		Addr: "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(sv.ListenAndServe())
}

func (s *Server) ControladorTelaLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Falha no form")
		return
	}

	if len(r.Form) == 0 || len(r.Form["email"]) != 1 || len(r.Form["senha"]) != 1 { //checar se apenas acessou a pagina ou tentou logar
		p := "./index.html"
		w.Header().Set("Content-type", "text/html")
		http.ServeFile(w, r, p)
		return
	}

	email := r.Form["email"][0]
	senha := r.Form["senha"][0]

	nome, err := s.c.EfetuarLogin(email, senha)
	if err != nil {
		http.Redirect(w, r, "/falhouLogin", http.StatusSeeOther)
		return
	}
	data := struct {
		NomeUsuario string
	}{
		NomeUsuario: nome,
	}
	t, _ := template.ParseFiles("logou.html")
	//t.ExecuteTemplate()
	//w.Header().Set("Content-type", "text/html")
	t.Execute(w, data)
}

func (s *Server) ControladorTelaCadastro(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() //Pegar dados do post
	if err != nil {
		log.Println("Falha no form cadastro")
		return
	}

	err = s.c.CadastrarUsuario(r.Form.Get("email"), r.Form.Get("senha"), r.Form.Get("nome"))
	if err != nil {
		log.Printf("erro ao cadastrar: %v", err)
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
