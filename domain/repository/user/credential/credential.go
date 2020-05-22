package credential

import (
	"github.com/team-gleam/kiwi-basket/domain/model/user/credential"
	"github.com/team-gleam/kiwi-basket/domain/model/user/token"
	"github.com/team-gleam/kiwi-basket/domain/model/user/username"
)

type ICredentialRepository interface {
	Append(credential.Auth) error
	Remove(username.Username) error
	Exists(token.Token) (bool, error)
	GetByToken(token.Token) (credential.Auth, error)
	GetByUsername(username.Username) (credential.Auth, error)
}
