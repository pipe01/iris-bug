package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.UseGlobal(middleware("first"))
	app.UseGlobal(middleware("second"))

	app.Get("/{id prefix(one)}", func(ctx iris.Context) { ctx.WriteString("first route") })
	app.Get("/{id prefix(two)}", func(ctx iris.Context) { ctx.WriteString("second route") })
	app.Get("/{id prefix(three)}", func(ctx iris.Context) { ctx.WriteString("third route") })

	app.Listen(":5555")
}

func middleware(str string) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		ctx.WriteString(fmt.Sprintf("Called %s middleware\n", str))
		ctx.Next()
	}
}
