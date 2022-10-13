package cli

import (
	"context"
	"fmt"
	"generate-deveui-cli/internal/constants"
	deveuigenerator "generate-deveui-cli/internal/deveui-generator"
	httpclient "generate-deveui-cli/internal/http-client"
	registerservice "generate-deveui-cli/internal/register-service"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func Start(ctx context.Context) {
	log.Print("starting batch run")
	devEuis := deveuigenerator.Generate()
	devEuiQueue := make(chan string, 100)
	httpClient := http.Client{
		Timeout: constants.DEFAULT_HTTP_TIMEOUT * time.Second,
	}
	wp := registerservice.NewRegisterService(
		constants.MAX_NO_OF_WORKERS,
		devEuiQueue,
		httpclient.NewHttpClient(httpClient),
	)
	for _, devEui := range devEuis {
		wp.AddTask(devEui)
	}
	wp.Run(ctx)
	fmt.Println("registered devEuis: ")
	for devEui := range wp.GetResults() {
		fmt.Println(devEui)
	}
	log.Print("completed batch run")
}
