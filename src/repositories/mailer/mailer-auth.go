package mailer

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func LoadCredentialsFromFile(filename string) (*MailerAuth, error) {
	if filepath.Ext(filename) != ".json" {
		return nil, errors.New("Credentials file must be a json file")
	}

	file, err := os.OpenFile(filename, os.O_RDWR, 0o600)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	auth := &MailerAuth{}
	if err := json.Unmarshal(bytes, auth); err != nil {
		return nil, err
	}

	return auth, nil
}

func LoadCredentialsFromFlags(username, password string) *MailerAuth {
	return &MailerAuth{
		Username: username,
		Password: password,
	}
}
