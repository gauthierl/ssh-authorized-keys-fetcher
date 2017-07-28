package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"text/template"
)

// GitlabFetchPubKeys from a Gitlab server
func GitlabFetchPubKeys(URLTemplate string, userName string) (pubKeys []string, err error) {
	var URLBuffer bytes.Buffer

	templateEngine, err := template.New("URL").Parse(URLTemplate)
	if err != nil {
		return
	}

	templateData := struct{ UserName string }{UserName: userName}

	err = templateEngine.Execute(&URLBuffer, templateData)
	if err != nil {
		return
	}

	URL := URLBuffer.String()

	response, err := http.Get(URL)
	if err != nil {
		return
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("Gitlab return an error (%s)", URL)
		return
	}

	rd := bufio.NewReader(response.Body)

	for {
		pubKey, silentErr := rd.ReadString('\n')
		if silentErr != nil {
			break
		}

		pubKeys = append(pubKeys, pubKey)
	}

	return
}
