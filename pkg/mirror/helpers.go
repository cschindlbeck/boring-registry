package mirror

import "github.com/boring-registry/boring-registry/pkg/core"

type OpenTofuPublicKey struct {
	Fingerprint string `json:"fingerprint"`
	ArmoredKey  string `json:"armored_key"`
	Origin      string `json:"origin"`
	OriginURL   string `json:"origin_url"`
}

type OpenTofuProvider struct {
	SigningKeys struct {
		PublicGPGKeys []OpenTofuPublicKey `json:"public_gpg_keys"`
	} `json:"signing_keys"`
}

func mapOpenTofuKeys(publicKeys []OpenTofuPublicKey) []core.GPGPublicKey {
	var mappedKeys []core.GPGPublicKey
	for _, key := range publicKeys {
		mappedKeys = append(mappedKeys, core.GPGPublicKey{
			KeyID:      key.Fingerprint,
			ASCIIArmor: key.ArmoredKey,
			Source:     key.Origin,
			SourceURL:  key.OriginURL,
		})
	}
	return mappedKeys
}
