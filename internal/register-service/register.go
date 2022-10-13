package registerservice

import (
	"fmt"
	config "generate-deveui-cli/internal/configuration"
	"io/ioutil"
	"net/http"
	"strings"
)

func (wp *workerPool) RegisterToLorawan(devEui string) error {
	bodyString := fmt.Sprintf(`{"deveui": %s}`, devEui)
	request, err := http.NewRequest(
		http.MethodPost,
		config.Config.Lorawan.Baseurl+config.Config.Lorawan.Endpoint,
		strings.NewReader(bodyString))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	r, err := wp.client.Do(request)
	if err != nil {
		return err
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("server: could not read request body: %s", err)
	}
	if r.StatusCode > 300 {
		return fmt.Errorf("failed to register deveui: %s, err: %s", devEui, string(reqBody))
	}
	return nil
}
