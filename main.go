package main

import (
	"flag"
	"fmt"

	m "./models"
	dbguardian "./utils"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {

	var dbpath string
	var renderpath string

	flag.StringVar(&dbpath, "DBPATH", "", "db path to storage subscribers")
	flag.StringVar(&renderpath, "STATICFILES", "", "location of html and js files")
	flag.Parse()

	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())

	if renderpath == "" {
		app.RegisterView(iris.HTML("./views", ".html"))
		app.StaticWeb("/dist", "./views/dist")
		app.StaticWeb("/public", "./views/public")
	} else {
		app.RegisterView(iris.HTML(renderpath, ".html"))
		app.StaticWeb("/dist", renderpath+"/dist")
		app.StaticWeb("/dist", renderpath+"/public")
	}

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.Gzip(true)
		ctx.View("index.html")
	})

	developerAPI := app.Party("/developers")
	{
		developerAPI.Post("/", func(ctx iris.Context) {
			dbguardian.Initialize(dbpath)
			var newUser m.User
			ctx.ReadJSON(&newUser)

			fmt.Println(&newUser)

			dbguardian.WriteData(&newUser)
			ctx.JSON(iris.Map{"message": "Gracias por suscribirte!", "status": "new user added"})
		})

		developerAPI.Get("/", func(ctx iris.Context) {
			dbguardian.Initialize(dbpath)
			result := dbguardian.ReadData()
			// if err != nil {
			// 	panic(err)
			// }
			fmt.Println(len(result))
			fmt.Println(result)
			ctx.WriteString("Nothing to look here YET")
		})

		developerAPI.Get("/write", func(ctx iris.Context) {
			ctx.WriteString("Nothing to look here")
		})
	}

	app.Run(iris.Addr(":8079"), iris.WithoutServerError(iris.ErrServerClosed))
}
