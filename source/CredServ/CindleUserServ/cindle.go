package CindleUserServ

import (
	controlador "CindleUserServ/controladores"
	repositorio "CindleUserServ/repositorios"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"

	"gopkg.in/yaml.v3"
)

type Cindle struct {
	ControladorUsuarios    controlador.ControladorUsuarios
	ControladorAssinaturas controlador.ControladorAssinatura
}

func (c *Cindle) EfetuarLogin(user repositorio.Usuario) (string, string, error) {
	return c.ControladorUsuarios.EfetuarLogin(user)
}

func (c *Cindle) CadastrarUsuario(user repositorio.Usuario) error {
	return c.ControladorUsuarios.EfetuarCadastro(user)
}

func (c *Cindle) VirarAssinante(email string, cartao controlador.Cartao) error {
	err := c.ControladorAssinaturas.EfetuarPagamentoAssinatura(cartao)
	if err != nil {
		return err
	}
	return c.ControladorUsuarios.MudarStatusAssinatura(email)
}

func (c *Cindle) VerificarStatusUsuario(email, token string) (bool, bool) {
	return c.ControladorUsuarios.VerificarStatusUsuario(email, token)
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
		ControladorAssinaturas: controlador.InicializarControladorAssinaturas(s.TipoRepositorio),
		ControladorUsuarios:    controlador.InicializarControladorUsuarios(s.TipoRepositorio),
	}
	s.c.CadastrarUsuario(repositorio.Usuario{
		"admin@admin.ad",
		"admin",
		"Administrador",
		true,
		true,
	})

	http.HandleFunc("/login", s.ControladorTelaLogin)
	http.HandleFunc("/cadastro", s.ControladorTelaCadastro)
	http.HandleFunc("/assinatura", s.ControladorTelaAssinatura)
	http.HandleFunc("/verificarusuario", s.ControladorVerificacaoUsuario)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})

	sv := &http.Server{
		Addr: ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(sv.ListenAndServe())
}

func (s *Server) ControladorVerificacaoUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		req, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Errorf("Error reading request")
		}
		emailToken := strings.Split(string(req), ",")
		Assinante, Admin := s.c.VerificarStatusUsuario(emailToken[0], emailToken[1])
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(fmt.Sprintf("%v,%v", Assinante, Admin)))
	}
}

func (s *Server) ControladorTelaLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		p := "./index.html"
		w.Header().Set("Content-type", "text/html")
		http.ServeFile(w, r, p)
		return
	} else {

		err := r.ParseForm()
		if err != nil {
			log.Println("Falha no form")
			return
		}

		email := r.Form["email"][0]
		senha := r.Form["senha"][0]

		usuarioLogando := repositorio.Usuario{
			Email: email,
			Senha: senha,
		}

		_, token, err := s.c.EfetuarLogin(usuarioLogando)
		if err != nil {
			fmt.Println("Login falhou")
			p := "./index.html"
			w.Header().Set("Content-type", "text/html")
			http.ServeFile(w, r, p)
			return
		}

		Assinante, _ := s.c.VerificarStatusUsuario(email, token)

		dados := struct {
			Usuario   string
			Token     string
			Assinante bool
		}{
			Assinante: Assinante,
			Usuario:   email,
			Token:     token,
		}
		t, _ := template.ParseFiles("./assinatura.html")
		t.Execute(w, dados)
	}
}

func (s *Server) ControladorTelaCadastro(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() //Pegar dados do post
	if err != nil {
		log.Println("Falha no form cadastro")
		return
	}

	novoUsuario := repositorio.Usuario{
		Email: r.Form.Get("email"),
		Senha: r.Form.Get("senha"),
		Nome:  r.Form.Get("nome"),
	}
	err = s.c.CadastrarUsuario(novoUsuario)
	if err != nil {
		log.Printf("erro ao cadastrar: %v", err)
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (s *Server) ControladorTelaAssinatura(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Useremail, err := r.Cookie("Username")
		Usetoken, err1 := r.Cookie("ApiToken")
		if err != nil || err1 != nil {
			fmt.Println("Falta credenciais")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		fmt.Println("cookies", Useremail.Value, Usetoken.Value)
		Assinante, _ := s.c.VerificarStatusUsuario(Useremail.Value, Usetoken.Value)
		dados := struct {
			Usuario   string
			Token     string
			Assinante bool
		}{
			Assinante: Assinante,
			Usuario:   Useremail.Value,
			Token:     Usetoken.Value,
		}
		t, _ := template.ParseFiles("./assinatura.html")
		t.Execute(w, dados)
		return
	}
	err := r.ParseForm() //Pegar dados do post
	if err != nil {
		log.Println("Falha no form cadastro")
		return
	}
	fmt.Println("Form", r.Form)
	cc := controlador.Cartao{
		Numero:   r.Form.Get("numero"),
		Cvv:      r.Form.Get("cvv"),
		Validade: r.Form.Get("validade"),
	}
	email, err := r.Cookie("Username")
	fmt.Println(r.Cookies())
	if err != nil {
		fmt.Println("Cookie de usuario nao encontrado")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	fmt.Println(cc, email)
	err = s.c.VirarAssinante(email.Value, cc)
	if err != nil {
		log.Printf("erro ao virar assinante: %v", err)
	}
	http.Redirect(w, r, "/assinatura", http.StatusSeeOther)
}
