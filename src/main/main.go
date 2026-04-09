package main

import (
	config "github.com/oderapi/configs"
	"github.com/oderapi/src/infra/persistence"
	"github.com/oderapi/src/main/factory"
	"github.com/oderapi/src/main/factory/db"
	"github.com/oderapi/src/main/factory/usecase"
	"gofr.dev/pkg/gofr"
)

func main() {
	if err := config.LoadProfile(); err != nil {
		panic(err)
	}

	app := gofr.New()
	postgresDb := db.MakePostgresDB()
	bootstrapSa := usecase.MakeBootstrapSa(postgresDb)

	app.OnStart(func(ctx *gofr.Context) error {
		err := persistence.RunMigrations(db.MakePostgresDB())
		if err != nil {
			ctx.Logger.Fatal(err)
			return err
		}

		input, err := usecase.MakeBootstrapSaInput()
		if err != nil {
			ctx.Logger.Fatal(err)
			return err
		}

		message, err := factory.MakeRunBootstrapSa(input, bootstrapSa)
		if err != nil {
			ctx.Logger.Fatal(err)
			return err
		}

		ctx.Logger.Infof("Boostrap SA: %s", message)
		return nil
	})
	app.Run()
}
