package redisclient

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/plugins/admin/controller"
)

func InitRouter(prefix string) *context.App {
	app := context.NewApp()

	authenticator := auth.SetPrefix(prefix).SetAuthFailCallback(func(ctx *context.Context) {
		ctx.Write(302, map[string]string{
			"Location": prefix + "/login",
		}, ``)
	}).SetPermissionDenyCallback(func(ctx *context.Context) {
		controller.ShowErrorPage(ctx, "permission denied")
	})

	app.GET(prefix+"/redisclient", authenticator.Middleware(Show))

	return app
}
