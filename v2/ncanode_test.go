package ncanode_test

import (
	"encoding/base64"
	"io/ioutil"
	"path/filepath"
)

const _defaultPassword = "Qwerty12"

const _defaultXML = "<?xml version=\"1.0\" encoding=\"utf-8\"?><root><name>NCANode</name></root>"

const _defaultRaw = "YXNkYXNk"

func base64content(name string) (string, error) {
	path := filepath.Join("../testdata", name)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return base64.RawStdEncoding.EncodeToString(content), nil
}

func filecontent(name string) (string, error) {
	path := filepath.Join("../testdata", name)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
