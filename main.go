package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	loadApi(app)

	err := app.Start()

	if err != nil {
		log.Fatal(err)
	}
}

func loadApi(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		g := e.Router.Group("/api/v1")

		g.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/blocks",
			Handler: func(ctx echo.Context) error {
				blocks, err := app.Dao().FindRecordsByExpr("blocks")

				result := []any{}

				if err != nil {
					return err
				}

				for _, block := range blocks {
					result = append(result, block.ColumnValueMap())
				}

				return ctx.JSON(http.StatusOK, result)
			},
		})

		return nil
	})
}
