package entities

type Wallet struct {
	PrivateKey          string `json:"privateKey"`
	PublicKey           string `json:"publicKey"`
	CompressedPublicKey string `json:"compressedPublicKey"`
	Namespace           string `json:"namespace,omitempty"`
	Type                string `json:"type,omitempty"`
}
