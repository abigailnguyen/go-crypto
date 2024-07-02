package generate_token

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GenerateApigeeToken(clientId, clientSecret string) {
	resp, err := http.Post(
		"https://data-test.uat.proptrack.com",
		"application/json",
		strings.NewReader(""),
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", string(body))
}
