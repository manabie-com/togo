package config

import (
	"crypto/ecdsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/yaml.v2"
)

var MonitorIfaces = []string{}

var ConfigVar = Config{}

const (
	GOBLIN_FT_CONFIG_PATH = "/etc/goblin/ft"
)

var (
	KeyPair      *tls.Certificate
	CertPool     *x509.CertPool
	JwtPublicKey *ecdsa.PublicKey
	// Limiter      *ratelimit.RequestLimit
	LocalNetwork []*net.IPNet
)

func InitJWT() {
	data, err := ioutil.ReadFile(ConfigVar.AuthConfig.JwtFile)
	if err != nil {
		panic(fmt.Errorf("error reading the jwt public key: %s", err))
	}

	JwtPublicKey, err = jwt.ParseECPublicKeyFromPEM(data)
	if err != nil {
		panic(fmt.Errorf("error parsing the jwt public key: %s", err))
	}
}

type DatabaseCfg struct {
	DbHost         string `yaml:"db_host"`
	DbType         string `yaml:"db_type"`
	DbName         string `yaml:"db_name"`
	DbUser         string `yaml:"db_user"`
	DbUserPassword string `yaml:"db_userpasswd"`
}

type AuthCfg struct {
	JwtFile string `yaml:"jw_public_key"`
	AuthUrl string `yaml:"auth_url"`
}

type TlsCfg struct {
	TlsServerKey string `yaml:"tls_server_key"`
	TlsServerCrt string `yaml:"tls_server_crt"`
	TlsServerCA  string `yaml:"tls_server_ca"`
}

type InternalCfg struct {
	PcapExtractDir string `yaml:"pcap_extract_dir"`
	PromtUrl       string `yaml:"promt_url"`
	RedisUrl       string `yaml:"redis_url"`
	AlertUrl       string `yaml:"alert_url"`
	PushGwUrl      string `yaml:"pushgw_url"`
	ZipDir         string `yaml:"zip_dir"`
	ZeekExtractDir string `yaml:"zeek_extract_dir"`
}

type DiskQuotaCfg struct {
	PcapQuota       int `yaml:"pcap_quota_gb"`
	FileCarvedQuota int `yaml:"file_carved_quota_gb"`
}

type EthernetCfg struct {
	Name        string `yaml:"name"`
	PcapDir     string `yaml:"pcap_dir"`
	TimelineDir string `yaml:"timeline_dir"`
}

type Config struct {
	DatabaseConfig DatabaseCfg `yaml:"database"`
	AuthConfig     AuthCfg     `yaml:"auth"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) error {

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&ConfigVar); err != nil {
		return err
	}

	return nil
}

func GetDbUrl() map[string]string {
	return map[string]string{
		"Host":     ConfigVar.DatabaseConfig.DbHost,
		"User":     ConfigVar.DatabaseConfig.DbUser,
		"Name":     ConfigVar.DatabaseConfig.DbName,
		"Password": ConfigVar.DatabaseConfig.DbUserPassword,
		"Type":     ConfigVar.DatabaseConfig.DbType,
	}
}

func AuthUrl() string {
	return ConfigVar.AuthConfig.AuthUrl
}
