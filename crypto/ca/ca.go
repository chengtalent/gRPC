package ca

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"database/sql"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3" // This blank import is required to load sqlite3 driver
	"github.com/op/go-logging"
	"google.golang.org/grpc/examples/helloworld/crypto/primitives"
)

// Hash is the common interface implemented by all hash functions.
type Hash interface {
	// Write (via the embedded io.Writer interface) adds more data to the running hash.
	// It never returns an error.
	io.Writer

	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	Sum(b []byte) []byte

	// Reset resets the Hash to its initial state.
	Reset()

	// Size returns the number of bytes Sum will return.
	Size() int

	// BlockSize returns the hash's underlying block size.
	// The Write method must be able to accept any amount
	// of data, but it may operate more efficiently if all writes
	// are a multiple of the block size.
	BlockSize() int
}

var caLogger = logging.MustGetLogger("ca")

// CA is the base certificate authority.
type CA struct {
	db *sql.DB

	path string

	priv *ecdsa.PrivateKey
	cert *x509.Certificate
	raw  []byte
}

// CertificateSpec defines the parameter used to create a new certificate.
type CertificateSpec struct {
	id           string
	commonName   string
	serialNumber *big.Int
	pub          interface{}
	usage        x509.KeyUsage
	NotBefore    *time.Time
	NotAfter     *time.Time
	ext          *[]pkix.Extension
}

var (
	mutex          = &sync.RWMutex{}
	caOrganization string
	caCountry      string
	rootPath       string
	caDir          string
)

// NewCertificateSpec creates a new certificate spec
func NewCertificateSpec(id string, commonName string, serialNumber *big.Int, pub interface{}, usage x509.KeyUsage, notBefore *time.Time, notAfter *time.Time, opt ...pkix.Extension) *CertificateSpec {
	spec := new(CertificateSpec)
	spec.id = id
	spec.commonName = commonName
	spec.serialNumber = serialNumber
	spec.pub = pub
	spec.usage = usage
	spec.NotBefore = notBefore
	spec.NotAfter = notAfter
	spec.ext = &opt
	return spec
}

// NewDefaultPeriodCertificateSpec creates a new certificate spec with notBefore a minute ago and not after 90 days from notBefore.
//
func NewDefaultPeriodCertificateSpec(id string, serialNumber *big.Int, pub interface{}, usage x509.KeyUsage, opt ...pkix.Extension) *CertificateSpec {
	return NewDefaultPeriodCertificateSpecWithCommonName(id, id, serialNumber, pub, usage, opt...)
}

// NewDefaultPeriodCertificateSpecWithCommonName creates a new certificate spec with notBefore a minute ago and not after 90 days from notBefore and a specifc commonName.
//
func NewDefaultPeriodCertificateSpecWithCommonName(id string, commonName string, serialNumber *big.Int, pub interface{}, usage x509.KeyUsage, opt ...pkix.Extension) *CertificateSpec {
	notBefore := time.Now().Add(-1 * time.Minute)
	notAfter := notBefore.Add(time.Hour * 24 * 90)
	return NewCertificateSpec(id, commonName, serialNumber, pub, usage, &notBefore, &notAfter, opt...)
}

// NewDefaultCertificateSpec creates a new certificate spec with serialNumber = 1, notBefore a minute ago and not after 90 days from notBefore.
//
func NewDefaultCertificateSpec(id string, pub interface{}, usage x509.KeyUsage, opt ...pkix.Extension) *CertificateSpec {
	serialNumber := big.NewInt(1)
	return NewDefaultPeriodCertificateSpec(id, serialNumber, pub, usage, opt...)
}

// NewDefaultCertificateSpecWithCommonName creates a new certificate spec with serialNumber = 1, notBefore a minute ago and not after 90 days from notBefore and a specific commonName.
//
func NewDefaultCertificateSpecWithCommonName(id string, commonName string, pub interface{}, usage x509.KeyUsage, opt ...pkix.Extension) *CertificateSpec {
	serialNumber := big.NewInt(1)
	return NewDefaultPeriodCertificateSpecWithCommonName(id, commonName, serialNumber, pub, usage, opt...)
}

// CacheConfiguration caches the viper configuration
func CacheConfiguration() {
	caOrganization = "pki.ca.subject.organization"
	caCountry = "pki.ca.subject.country"
	rootPath = "/tmp/ethereum"
	caDir = "CA"
}

// GetID returns the spec's ID field/value
//
func (spec *CertificateSpec) GetID() string {
	return spec.id
}

// GetCommonName returns the spec's Common Name field/value
//
func (spec *CertificateSpec) GetCommonName() string {
	return spec.commonName
}

// GetSerialNumber returns the spec's Serial Number field/value
//
func (spec *CertificateSpec) GetSerialNumber() *big.Int {
	return spec.serialNumber
}

// GetPublicKey returns the spec's Public Key field/value
//
func (spec *CertificateSpec) GetPublicKey() interface{} {
	return spec.pub
}

// GetUsage returns the spec's usage (which is the x509.KeyUsage) field/value
//
func (spec *CertificateSpec) GetUsage() x509.KeyUsage {
	return spec.usage
}

// GetNotBefore returns the spec NotBefore (time.Time) field/value
//
func (spec *CertificateSpec) GetNotBefore() *time.Time {
	return spec.NotBefore
}

// GetNotAfter returns the spec NotAfter (time.Time) field/value
//
func (spec *CertificateSpec) GetNotAfter() *time.Time {
	return spec.NotAfter
}

// GetOrganization returns the spec's Organization field/value
//
func (spec *CertificateSpec) GetOrganization() string {
	return caOrganization
}

// GetCountry returns the spec's Country field/value
//
func (spec *CertificateSpec) GetCountry() string {
	return caCountry
}

// GetSubjectKeyID returns the spec's subject KeyID
//
func (spec *CertificateSpec) GetSubjectKeyID() *[]byte {
	return &[]byte{1, 2, 3, 4}
}

// GetSignatureAlgorithm returns the X509.SignatureAlgorithm field/value
//
func (spec *CertificateSpec) GetSignatureAlgorithm() x509.SignatureAlgorithm {
	return x509.ECDSAWithSHA384
}

// GetExtensions returns the sepc's extensions
//
func (spec *CertificateSpec) GetExtensions() *[]pkix.Extension {
	return spec.ext
}

// TableInitializer is a function type for table initialization
type TableInitializer func(*sql.DB) error

func InitializeCommonTables(db *sql.DB) error {
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS Certificates (row INTEGER PRIMARY KEY, id VARCHAR(64), timestamp INTEGER, usage INTEGER, cert BLOB, hash BLOB, kdfkey BLOB)"); err != nil {
		return err
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS Users (row INTEGER PRIMARY KEY, id VARCHAR(64), enrollmentId VARCHAR(100), role INTEGER, metadata VARCHAR(256), token BLOB, state INTEGER, key BLOB)"); err != nil {
		return err
	}
	return nil
}

// NewCA sets up a new CA.
func NewCA(name string, initTables TableInitializer) *CA {
	ca := new(CA)
	ca.path = filepath.Join(rootPath, caDir)

	caLogger.Info(ca.path)

	if _, err := os.Stat(ca.path); err != nil {
		caLogger.Info("Fresh start; creating databases, key pairs, and certificates.")

		if err := os.MkdirAll(ca.path, 0755); err != nil {
			caLogger.Panic(err)
		}
	}

	// open or create certificate database
	db, err := sql.Open("sqlite3", ca.path+"/"+name+".db")
	if err != nil {
		caLogger.Panic(err)
	}

	if err = db.Ping(); err != nil {
		caLogger.Panic(err)
	}

	if err = initTables(db); err != nil {
		caLogger.Panic(err)
	}
	ca.db = db

	// read or create signing key pair
	priv, err := ca.readCAPrivateKey(name)
	if err != nil {
		priv = ca.createCAKeyPair(name)
	}
	ca.priv = priv

	// read CA certificate, or create a self-signed CA certificate
	raw, err := ca.readCACertificate(name)
	if err != nil {
		raw = ca.createCACertificate(name, &ca.priv.PublicKey)
	}
	cert, err := x509.ParseCertificate(raw)
	if err != nil {
		caLogger.Panic(err)
	}

	ca.raw = raw
	ca.cert = cert

	return ca
}

func (ca *CA) IssueCertificate(in []byte, name string) ([]byte, error) {
	raw, err := ca.readCACertificate(name)

	if err != nil {
		block, _ := pem.Decode(in)
		pub, err := x509.ParsePKIXPublicKey(block.Bytes)
		if(err != nil) {
			caLogger.Panic(err)
		}

		pubkey := pub.(*ecdsa.PublicKey)
		raw = ca.createCACertificate(name, (pubkey))
	}

	cooked := pem.EncodeToMemory(
		&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: raw,
		})

	return cooked, nil
}

func (ca *CA) GetCACertificate() ([]byte) {
	raw := ca.cert.Raw

	cooked := pem.EncodeToMemory(
		&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: raw,
		})

	return cooked
}

// Stop Close closes down the CA.
func (ca *CA) Stop() error {
	err := ca.db.Close()
	if err == nil {
		caLogger.Debug("Shutting down CA - Successfully")
	} else {
		caLogger.Debug(fmt.Sprintf("Shutting down CA - Error closing DB [%s]", err))
	}
	return err
}

func (ca *CA) createCAKeyPair(name string) *ecdsa.PrivateKey {
	caLogger.Debug("Creating CA key pair.")

	curve := primitives.GetDefaultCurve()

	priv, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err == nil {
		raw, _ := x509.MarshalECPrivateKey(priv)
		cooked := pem.EncodeToMemory(
			&pem.Block{
				Type:  "ECDSA PRIVATE KEY",
				Bytes: raw,
			})
		err = ioutil.WriteFile(ca.path+"/"+name+".priv", cooked, 0644)
		if err != nil {
			caLogger.Panic(err)
		}

		raw, _ = x509.MarshalPKIXPublicKey(&priv.PublicKey)
		cooked = pem.EncodeToMemory(
			&pem.Block{
				Type:  "ECDSA PUBLIC KEY",
				Bytes: raw,
			})
		err = ioutil.WriteFile(ca.path+"/"+name+".pub", cooked, 0644)
		if err != nil {
			caLogger.Panic(err)
		}
	}
	if err != nil {
		caLogger.Panic(err)
	}

	return priv
}

func (ca *CA) readCAPrivateKey(name string) (*ecdsa.PrivateKey, error) {
	caLogger.Debug("Reading CA private key.")

	cooked, err := ioutil.ReadFile(ca.path + "/" + name + ".priv")
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(cooked)
	return x509.ParseECPrivateKey(block.Bytes)
}

func (ca *CA) createCACertificate(name string, pub *ecdsa.PublicKey) []byte {
	caLogger.Debug("Creating CA certificate.")

	raw, err := ca.newCertificate(name, pub, x509.KeyUsageDigitalSignature|x509.KeyUsageCertSign, nil)
	if err != nil {
		caLogger.Panic(err)
	}

	cooked := pem.EncodeToMemory(
		&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: raw,
		})
	err = ioutil.WriteFile(ca.path+"/"+name+".cert", cooked, 0644)
	if err != nil {
		caLogger.Panic(err)
	}

	return raw
}

func (ca *CA) readCACertificate(name string) ([]byte, error) {
	caLogger.Debug("Reading CA certificate.")

	path := ca.path + "/" + name + ".cert"

	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	cooked, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(cooked)
	return block.Bytes, nil
}

func (ca *CA) createCertificate(id string, pub interface{}, usage x509.KeyUsage, timestamp int64, kdfKey []byte, opt ...pkix.Extension) ([]byte, error) {
	spec := NewDefaultCertificateSpec(id, pub, usage, opt...)
	return ca.createCertificateFromSpec(spec, timestamp, kdfKey, true)
}

func (ca *CA) createCertificateFromSpec(spec *CertificateSpec, timestamp int64, kdfKey []byte, persist bool) ([]byte, error) {
	caLogger.Debug("Creating certificate for " + spec.GetID() + ".")

	raw, err := ca.newCertificateFromSpec(spec)
	if err != nil {
		caLogger.Error(err)
		return nil, err
	}

	if persist {
		err = ca.persistCertificate(spec.GetID(), timestamp, spec.GetUsage(), raw, kdfKey)
	}

	return raw, err
}

func (ca *CA) persistCertificate(id string, timestamp int64, usage x509.KeyUsage, certRaw []byte, kdfKey []byte) error {
	mutex.Lock()
	defer mutex.Unlock()

	hash := primitives.NewHash()
	hash.Write(certRaw)
	var err error

	if _, err = ca.db.Exec("INSERT INTO Certificates (id, timestamp, usage, cert, hash, kdfkey) VALUES (?, ?, ?, ?, ?, ?)", id, timestamp, usage, certRaw, hash.Sum(nil), kdfKey); err != nil {
		caLogger.Error(err)
	}
	return err
}

func (ca *CA) newCertificate(id string, pub interface{}, usage x509.KeyUsage, ext []pkix.Extension) ([]byte, error) {
	spec := NewDefaultCertificateSpec(id, pub, usage, ext...)
	return ca.newCertificateFromSpec(spec)
}

func (ca *CA) newCertificateFromSpec(spec *CertificateSpec) ([]byte, error) {
	notBefore := spec.GetNotBefore()
	notAfter := spec.GetNotAfter()

	parent := ca.cert
	isCA := parent == nil

	tmpl := x509.Certificate{
		SerialNumber: spec.GetSerialNumber(),
		Subject: pkix.Name{
			CommonName:   spec.GetCommonName(),
			Organization: []string{spec.GetOrganization()},
			Country:      []string{spec.GetCountry()},
		},
		NotBefore: *notBefore,
		NotAfter:  *notAfter,

		SubjectKeyId:       *spec.GetSubjectKeyID(),
		SignatureAlgorithm: spec.GetSignatureAlgorithm(),
		KeyUsage:           spec.GetUsage(),

		BasicConstraintsValid: true,
		IsCA: isCA,
	}

	if len(*spec.GetExtensions()) > 0 {
		tmpl.Extensions = *spec.GetExtensions()
		tmpl.ExtraExtensions = *spec.GetExtensions()
	}
	if isCA {
		parent = &tmpl
	}

	raw, err := x509.CreateCertificate(
		rand.Reader,
		&tmpl,
		parent,
		spec.GetPublicKey(),
		ca.priv,
	)
	if isCA && err != nil {
		caLogger.Panic(err)
	}

	return raw, err
}

func (ca *CA) readCertificateByKeyUsage(id string, usage x509.KeyUsage) ([]byte, error) {
	caLogger.Debugf("Reading certificate for %s and usage %v", id, usage)

	mutex.RLock()
	defer mutex.RUnlock()

	var raw []byte
	err := ca.db.QueryRow("SELECT cert FROM Certificates WHERE id=? AND usage=?", id, usage).Scan(&raw)

	if err != nil {
		caLogger.Debugf("readCertificateByKeyUsage() Error: %v", err)
	}

	return raw, err
}

func (ca *CA) readCertificateByTimestamp(id string, ts int64) ([]byte, error) {
	caLogger.Debug("Reading certificate for " + id + ".")

	mutex.RLock()
	defer mutex.RUnlock()

	var raw []byte
	err := ca.db.QueryRow("SELECT cert FROM Certificates WHERE id=? AND timestamp=?", id, ts).Scan(&raw)

	return raw, err
}

func (ca *CA) readCertificates(id string, opt ...int64) (*sql.Rows, error) {
	caLogger.Debug("Reading certificatess for " + id + ".")

	mutex.RLock()
	defer mutex.RUnlock()

	if len(opt) > 0 && opt[0] != 0 {
		return ca.db.Query("SELECT cert, kdfkey FROM Certificates WHERE id=? AND timestamp=? ORDER BY usage", id, opt[0])
	}

	return ca.db.Query("SELECT cert, kdfkey FROM Certificates WHERE id=?", id)
}

func (ca *CA) readCertificateSets(id string, start, end int64) (*sql.Rows, error) {
	caLogger.Debug("Reading certificate sets for " + id + ".")

	mutex.RLock()
	defer mutex.RUnlock()

	return ca.db.Query("SELECT cert, kdfKey, timestamp FROM Certificates WHERE id=? AND timestamp BETWEEN ? AND ? ORDER BY timestamp", id, start, end)
}

func (ca *CA) readCertificateByHash(hash []byte) ([]byte, error) {
	caLogger.Debug("Reading certificate for hash " + string(hash) + ".")

	mutex.RLock()
	defer mutex.RUnlock()

	var raw []byte
	row := ca.db.QueryRow("SELECT cert FROM Certificates WHERE hash=?", hash)
	err := row.Scan(&raw)

	return raw, err
}
