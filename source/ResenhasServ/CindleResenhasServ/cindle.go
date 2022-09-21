package CindleResenhasServ

import (
	controlador "CindleResenhasServ/controladores"
	repositorio "CindleResenhasServ/repositorios"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Cindle struct {
	ControladorResenha controlador.ControladorResenhas
}

func (c *Cindle) CadastrarResenha(re repositorio.Resenha) error {
	return c.ControladorResenha.CadastrarResenha(re)
}

func (c *Cindle) ListarResenhas() []repositorio.Resenha {
	return c.ControladorResenha.ListarResenha()
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

	s.TipoRepositorio = "array"

	s.c = &Cindle{
		ControladorResenha: controlador.InicializarControladorResenhas(s.TipoRepositorio),
	}

	http.HandleFunc("/resenha", s.ControladorTelaResenhas)

	sv := &http.Server{
		Addr:         ":8002",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(sv.ListenAndServe())
}

type DadosResenha struct {
	Resenhas []repositorio.Resenha
}

func resenhaValida(re repositorio.Resenha) bool {
	return !(len(re.Usuario) == 0 || len(re.NomeLivro) == 0 || len(re.Descricao) == 0)
}

func (s *Server) ControladorTelaResenhas(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("Cadastrado")
		err := r.ParseForm()
		if err != nil {
			log.Println("Falha no form")
			return
		}
		fmt.Println("form", r.Form)
		res := repositorio.Resenha{
			r.Form["usuario"][0],
			r.Form["livro"][0],
			r.Form["resenha"][0],
		}
		if resenhaValida(res) {
			s.c.CadastrarResenha(res)
		} else {
			fmt.Println("faltou dado")
		}
	}
	dados := DadosResenha{
		Resenhas: s.c.ListarResenhas(),
	}
	fmt.Println("servindo")
	fmt.Println("dados", dados)
	t, _ := template.ParseFiles("./resenhas.html")
	t.Execute(w, dados)
}
