package crypto

import (
	"github.com/op/go-logging"
	"google.golang.org/grpc/examples/helloworld/crypto/primitives"
)

var (
	log = logging.MustGetLogger("crypto")
)

// Init initializes the crypto layer. It load from viper the security level
// and the logging setting.
func Init() (err error) {
	// Init security level
	securityLevel := 256
	hashAlgorithm := "SHA3"

	log.Debugf("Working at security level [%d]", securityLevel)
	if err = primitives.InitSecurityLevel(hashAlgorithm, securityLevel); err != nil {
		log.Errorf("Failed setting security level: [%s]", err)

		return
	}

	return
}
