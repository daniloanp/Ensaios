package conn

import "errors"

type Configuration struct {
	DbName                  string `json:"dbname"`
	User                    string `json:"user"`
	Password                string `json:"password"`
	Host                    string `json:"host"`
	Port                    uint64 `json:"port"`
	SslMode                 string `json:"sslmode"`
	FallBackApplicationName string `json:"fallback_application_name"`
	ConnectionTimeout       uint64 `json:"connect_timeout"`
	SslCert                 string `json:"sslcert"`
	SslKey                  string `json:"sslkey"`
	SslRootKey              string `json:"sslrootcert"`
}

/* SslMode
* disable - No SSL
* require - Always SSL (skip verification)
* verify-ca - Always SSL (verify that the certificate presented by the server was signed by a trusted CA)
* verify-full - Always SSL (verify that the certification presented by the server was signed by a trusted CA and the server host name matches the one in the certificate)
*/

var configuration *Configuration = nil

//SetConfiguration, you should not change configuration more than once
func SetConfiguration(c *Configuration) error {
	if configuration != nil {
		return errors.New("You should not change configuration more than once")
	}
	configuration = c
	if err := configuration.loadDefaults(); err != nil {
		return err
	}
	postgres = nil
	return nil
}
const (
	sslModeDisable = "disable"
	sslModeRequire = "require"
	sslModeVerifyCa = "verify-ca"
	sslModeVerifyFull = "verify-full"
)

var (
	defaultConfiguration = Configuration {
		DbName: "ensaios",
		User: "postgres",
		Password: "<password>",
		Host: "127.0.0.1",
		Port: 5432,
		SslMode: sslModeDisable,
		FallBackApplicationName: "",
		ConnectionTimeout: 0,
		SslCert: "",
		SslKey: "",
		SslRootKey: "",
	}
)


func validSslMode(sslMode string) bool {
	for _, val := range [...]string{sslModeDisable, sslModeRequire, sslModeVerifyCa, sslModeVerifyFull} {
		if val == sslMode {
			return true
		}
	}
	return false
}


//LoadDefaults prepare if funcion return
func (c *Configuration) loadDefaults() error {
	if c.Port == 0 {
		c.Port = 5432
	}
	if c.SslMode == "" {
		c.SslMode = sslModeDisable
	} else if !validSslMode(c.SslMode){
		return errors.New("Invalid ssl mode")
	}
	if c.DbName == "" {
		c.DbName = defaultConfiguration.DbName
	}
	if c.User == "" {
		c.User = defaultConfiguration.User
	}
	if c.Password == "" {
		c.Password = defaultConfiguration.Password
	}
	if c.Host == "" {
		c.Host = defaultConfiguration.Host
	}
	if c.Port == 0 {
		c.Port = defaultConfiguration.Port
	}
	if c.SslMode == "" {
		c.SslMode = defaultConfiguration.SslMode
	} else if !validSslMode(c.SslMode) {
		return errors.New("Invalid SslMode passed!")
	}
	if c.FallBackApplicationName == "" {
		c.FallBackApplicationName = defaultConfiguration.FallBackApplicationName
	}
	if c.ConnectionTimeout == 0 {
		c.ConnectionTimeout = defaultConfiguration.ConnectionTimeout
	}
	if c.SslCert == "" {
		c.SslCert = defaultConfiguration.SslCert
	}

	if c.SslKey == "" {
		c.SslKey = defaultConfiguration.SslKey
	}
	if c.SslRootKey == "" {
		c.SslRootKey = defaultConfiguration.SslRootKey
	}
	return nil
}



//* dbname - The name of the database to connect to
//* user - The user to sign in as
//* password - The user's password
//* host - The host to connect to. Values that start with / are for unix domain sockets. (default is localhost)
//* port - The port to bind to. (default is 5432)
//* sslmode - Whether or not to use SSL (default is require, this is not the default for libpq)
//* fallback_application_name - An application_name to fall back to if one isn't provided.
//* connect_timeout - Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
//* sslcert - Cert file location. The file must contain PEM encoded data.
//* sslkey - Key file location. The file must contain PEM encoded data.
//* sslrootcert - The location of the root certificate file. The file must contain PEM encoded data.
