// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AppLogDecryptionAlgorithm undocumented
type AppLogDecryptionAlgorithm int

const (
	// AppLogDecryptionAlgorithmVAes256 undocumented
	AppLogDecryptionAlgorithmVAes256 AppLogDecryptionAlgorithm = 0
)

// AppLogDecryptionAlgorithmPAes256 returns a pointer to AppLogDecryptionAlgorithmVAes256
func AppLogDecryptionAlgorithmPAes256() *AppLogDecryptionAlgorithm {
	v := AppLogDecryptionAlgorithmVAes256
	return &v
}