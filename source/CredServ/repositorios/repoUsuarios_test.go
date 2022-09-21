package repositorio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestSave(t *testing.T) {
	a := IniciarRepoUsuarioArray()
	dados, _ := ioutil.ReadFile("dados.txt")
	json.Unmarshal(dados, &a.usuarios)
	u := Usuario{
		"bruno@",
		"123456",
		"Caio",
		false,
		false,
	}
	a.Insert(u)
	a.persistir()

	t.Fatalf("failed")
	//fmt.Println(json.Marshal())
}

func TestLoad(t *testing.T) {
	a := IniciarRepoUsuarioArray()
	dados, _ := ioutil.ReadFile("dados.txt")
	json.Unmarshal(dados, &a.usuarios)
	fmt.Println(a.Get())
	t.Fatalf("failed")
}
