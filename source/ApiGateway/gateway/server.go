package gateway

import (
	servicos "ApiGateway/servicos"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Cindle struct {
	UsuariosServico *servicos.UsuariosServ
	ResenhasServico *servicos.ResenhasServ
	LivrosSevico    *servicos.LivrosServ
}

func (s *Server) ControladorLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		s.c.UsuariosServico.AcessarPaginaLogin(w, r)
	} else if r.Method == "POST" {
		s.c.UsuariosServico.RealizarLogin(w, r)
	}
}

func (s *Server) ControladorLivros(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Cookie("Username")
	token, _ := r.Cookie("ApiToken")
	perms := s.c.UsuariosServico.ChecarPermissoes(user.Value, token.Value, w, r)
	if perms[0] == "false" {
		s.c.UsuariosServico.IrParaAssinatura(w, r)
	}
	if r.Method == "GET" {
		s.c.LivrosSevico.ListarLivros(w, r)
	}
}

func (s *Server) ControladorCadastroLivros(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Cookie("Username")
	token, _ := r.Cookie("ApiToken")
	perms := s.c.UsuariosServico.ChecarPermissoes(user.Value, token.Value, w, r)
	if perms[1] == "false" {
		s.ControladorLivros(w, r)
	}
	if r.Method == "GET" {
		s.c.LivrosSevico.PaginaAdicionarLivro(w, r)
	}
	if r.Method == "POST" {
		s.c.LivrosSevico.AdicionarLivro(w, r)
	}
}

func (s *Server) ControladorCadastro(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		s.c.UsuariosServico.RealizarCadastro(w, r)
	}
}

func (s *Server) ControladorResenha(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Cookie("Username")
	token, _ := r.Cookie("ApiToken")
	perms := s.c.UsuariosServico.ChecarPermissoes(user.Value, token.Value, w, r)
	if perms[0] == "false" {
		s.c.UsuariosServico.IrParaAssinatura(w, r)
	}
	if r.Method == "GET" {
		s.c.ResenhasServico.ListarResenhas(w, r)
	} else {
		r.ParseForm()
		livro := r.Form.Get("livro")
		if s.c.LivrosSevico.VerificarExistenciaLivro(livro) {
			resenha := r.Form.Get("resenha")
			fmt.Println("Livro existente")
			s.c.ResenhasServico.AdicionarResenha(user.Value, livro, resenha, w, r)
			return
		}
		fmt.Println("Livro nao existente")
		s.c.ResenhasServico.ListarResenhas(w, r)
	}
}

func (s *Server) ControladorAssinatura(w http.ResponseWriter, r *http.Request) {
	// user, _ := r.Cookie("Username")
	// token, _ := r.Cookie("ApiToken")
	// perms := s.c.UsuariosServico.ChecarPermissoes(user.Value, token.Value, w, r)
	// fmt.Println(perms)
	// if perms[0] == "false" {
	// s.c.UsuariosServico.IrParaAssinatura(w, r)
	// }
	if r.Method == "GET" {
		s.c.UsuariosServico.IrParaAssinatura(w, r)
	} else {
		s.c.UsuariosServico.RealizarAssinatura(w, r)
	}
}

func RedirecionarLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

/*
func (c *Cindle) EfetuarLogin(user repositorio.Usuario) (string, error) {
	return c.ControladorUsuarios.EfetuarLogin(user)
}

func (c *Cindle) CadastrarUsuario(user repositorio.Usuario) error {
	return c.ControladorUsuarios.EfetuarCadastro(user)
}*/

type Server struct {
	sv *http.Server
	c  *Cindle
}

func StartServer() {
	s := &Server{}
	/*err := s.LerArquivoConfiguracao()
	if err != nil {
		panic(err)
	}*/

	s.c = &Cindle{
		UsuariosServico: servicos.InicializarUsuariosServ("http://cred:8000"),
		ResenhasServico: servicos.InicializarResenhasServ("http://resenhas:8002"),
		LivrosSevico:    servicos.InicializarLivrosServ("http://livros:8003"),
	}
	http.HandleFunc("/login", s.ControladorLogin)
	http.HandleFunc("/livros", s.ControladorLivros)
	http.HandleFunc("/cadastrarlivros", s.ControladorCadastroLivros)
	http.HandleFunc("/cadastro", s.ControladorCadastro)
	http.HandleFunc("/resenha", s.ControladorResenha)
	http.HandleFunc("/assinatura", s.ControladorAssinatura)
	http.HandleFunc("/", RedirecionarLogin)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	sv := &http.Server{
		Addr: ":7999",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(sv.ListenAndServe())
}
