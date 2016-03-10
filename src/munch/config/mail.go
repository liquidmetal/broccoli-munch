// Functions related to mail settings

package config

func (config *Config) parseConfigMailValues(t *ConfigReader) {
	// Mail settings
	config.mail.publickey = t.Mail["publickey"]
	config.mail.privatekey = t.Mail["privatekey"]
}

func (config *Config) GetMailPublicKey() string {
	return config.mail.publickey
}

func (config *Config) GetMailPrivateKey() string {
	return config.mail.privatekey
}
