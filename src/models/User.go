package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"` //retira o ID caso n venha
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Valida e formata o usuario recebido
func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("nome é obrigatório e não pode estar em branco")
	}

	if user.Nick == "" {
		return errors.New("nick é obrigatório e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("email é obrigatório e não pode estar em branco")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("e-mail inserido é inválido")
	}

	if step == "register" && user.Password == "" {
		return errors.New("senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		hashPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashPassword)
	}
	return nil
}
