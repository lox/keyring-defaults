package defaults

import (
	keychain "github.com/lox/keyring-keychain"
	keyctl "github.com/lox/keyring-keyctl"
	kwallet "github.com/lox/keyring-kwallet"
	pass "github.com/lox/keyring-pass"
	secretservice "github.com/lox/keyring-secretservice"
	wincred "github.com/lox/keyring-wincred"
	"github.com/lox/keyring/v2"
)

// Providers returns the default first-party provider order.
func Providers(extra ...keyring.Provider) []keyring.Provider {
	return append([]keyring.Provider{
		wincred.Provider(),
		keychain.Provider(keychain.TrustApplication(false)),
		secretservice.Provider(),
		kwallet.Provider(),
		keyctl.Provider(),
		pass.Provider(),
	}, extra...)
}
