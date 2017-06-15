package main

import (
	"dispatch"
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()

	manager := dispatch.NewManager(200)
	manager.Dispatch()

	app.OnErrorCode(iris.StatusNotFound, func(ctx context.Context) {
		ctx.WriteString("not found")
	})

	from := 1
	step := 200
	app.Get("/start", func(ctx context.Context) {
		for i := from; i < from+step; i++ {
			displayJob := &dispatch.DisplayJob{Title: "title" + fmt.Sprintln(i)}
			outputJob := dispatch.OutputJob{Output: "Output" + fmt.Sprintln(i)}
			//
			manager.JobQueue <- displayJob
			manager.JobQueue <- outputJob
		}
		from = from + step
		ctx.WriteString("input success")
	})

	app.Run(iris.Addr(":8080"))
}
