// Code generated by github.com/kyoh86/appenv.Generator DO NOT EDIT.

package out

import (
	def "github.com/kyoh86/appenv/internal/def"
	keyring "github.com/zalando/go-keyring"
)

const DiscardKeyringService string = ""

type Keyring struct {
	Token *def.Token
}

func loadKeyring(keyringService string) (key Keyring, err error) {
	if keyringService == DiscardKeyringService {
		return
	}
	{
		v, err := keyring.Get(keyringService, "token")
		if err == nil {
			var value def.Token
			if err = value.UnmarshalText([]byte(v)); err != nil {
				return key, err
			}
			key.Token = &value
		}
	}
	return
}

func saveKeyring(keyringService string, key *Keyring) (err error) {
	if keyringService == DiscardKeyringService {
		return
	}
	{
		buf, err := key.Token.MarshalText()
		if err != nil {
			return err
		}
		if err := keyring.Set(keyringService, "token", string(buf)); err != nil {
			return err
		}
	}
	return nil
}
