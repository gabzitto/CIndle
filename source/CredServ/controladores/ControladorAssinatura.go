package controlador

import (
	"fmt"
	"regexp"
)

type ControladorAssinatura interface {
	EfetuarPagamentoAssinatura(Cartao) error
}

type Cartao struct {
	Numero   string
	Cvv      string
	Validade string
}

type ControladorAssinaturaImp struct {
}

func InicializarControladorAssinaturas(cfg string) *ControladorAssinaturaImp {
	return &ControladorAssinaturaImp{}
}

func (c *ControladorAssinaturaImp) EfetuarPagamentoAssinatura(cc Cartao) error {
	err := c.validarNumero(cc.Numero)
	if err != nil {
		return err
	}
	err = c.validarCVV(cc.Cvv)
	if err != nil {
		return err
	}
	err = c.validarValidade(cc.Validade)
	if err != nil {
		return err
	}
	return nil
}

func (c *ControladorAssinaturaImp) validarNumero(numero string) error {
	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)
	if re.MatchString(numero) {
		return nil
	}
	return fmt.Errorf("Numero de cartao invalido")
}

func (c *ControladorAssinaturaImp) validarCVV(cvv string) error {
	if len(cvv) > 4 || len(cvv) < 3 {
		return fmt.Errorf("CVV invalido")
	}
	re := regexp.MustCompile(`[^0-9 ]+`)
	cvv = re.ReplaceAllString(cvv, "")
	if len(cvv) > 4 || len(cvv) < 3 {
		return fmt.Errorf("CVV invalido")
	}
	return nil
}

func (c *ControladorAssinaturaImp) validarValidade(validade string) error {
	return nil
}
