package pkg

import (
	"fmt"
	"net/http"

	"github.com/Unleash/unleash-client-go/v3"
	"github.com/Unleash/unleash-client-go/v3/context"
)

func Run() {
	ctx := context.Context{
		UserId:     AppConfig.UserId,
		Properties: map[string]string{"Fund": "fundo"},
	}

	enabled := unleash.IsEnabled(AppConfig.FlagName, unleash.WithContext(ctx))
    fmt.Printf("%s is %v\n", AppConfig.FlagName, enabled)
}

func InitUnleash() {
	if err := unleash.Initialize(
		unleash.WithListener(&unleash.DebugListener{}),
		unleash.WithAppName(AppConfig.AppName),
		unleash.WithUrl(AppConfig.Url),
		unleash.WithProjectName(AppConfig.ProjectName),
		unleash.WithEnvironment(AppConfig.Environment),
		unleash.WithCustomHeaders(http.Header{"Authorization": { AppConfig.Authorization }}),
	); err != nil {
		panic(err)
    }

	unleash.WaitForReady()
}
