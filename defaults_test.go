package defaults

import (
	"context"
	"slices"
	"testing"

	"github.com/lox/keyring/v2"
)

func TestProvidersOrder(t *testing.T) {
	file := keyring.Provider{
		Backend: keyring.FileBackend,
		Open: func(context.Context, keyring.OpenOptions) (keyring.Keyring, error) {
			return nil, keyring.ErrUnavailable
		},
	}

	got := Providers(file)
	backends := make([]keyring.Backend, 0, len(got))
	for _, provider := range got {
		if provider.Open == nil {
			t.Fatalf("%s provider has nil Open", provider.Backend)
		}
		backends = append(backends, provider.Backend)
	}

	want := []keyring.Backend{
		keyring.WinCredBackend,
		keyring.KeychainBackend,
		keyring.SecretServiceBackend,
		keyring.KWalletBackend,
		keyring.KeyCtlBackend,
		keyring.PassBackend,
		keyring.FileBackend,
	}
	if !slices.Equal(backends, want) {
		t.Fatalf("Providers backends = %v, want %v", backends, want)
	}
}

func TestKeychainTrustApplicationKeepsProviderOrder(t *testing.T) {
	got := Providers(KeychainTrustApplication(false))

	backends := make([]keyring.Backend, 0, len(got))
	keychainCount := 0
	for _, provider := range got {
		backends = append(backends, provider.Backend)
		if provider.Backend == keyring.KeychainBackend {
			keychainCount++
		}
	}

	want := []keyring.Backend{
		keyring.WinCredBackend,
		keyring.KeychainBackend,
		keyring.SecretServiceBackend,
		keyring.KWalletBackend,
		keyring.KeyCtlBackend,
		keyring.PassBackend,
	}
	if !slices.Equal(backends, want) {
		t.Fatalf("Providers backends = %v, want %v", backends, want)
	}
	if keychainCount != 1 {
		t.Fatalf("expected one keychain provider, got %d", keychainCount)
	}
}
