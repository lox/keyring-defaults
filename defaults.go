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
	providers := []keyring.Provider{
		wincred.Provider(),
		keychain.Provider(),
		secretservice.Provider(),
		kwallet.Provider(),
		keyctl.Provider(),
		pass.Provider(),
	}
	for _, provider := range extra {
		providers = replaceOrAppend(providers, provider)
	}
	return providers
}

// KeychainTrustApplication returns a keychain provider with trust configured.
func KeychainTrustApplication(enabled bool) keyring.Provider {
	return keychain.Provider(keychain.TrustApplication(enabled))
}

func replaceOrAppend(providers []keyring.Provider, provider keyring.Provider) []keyring.Provider {
	for i := range providers {
		if providers[i].Backend == provider.Backend {
			providers[i] = provider
			return providers
		}
	}
	return append(providers, provider)
}
