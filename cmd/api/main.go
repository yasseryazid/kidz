package main

import (
	"github.com/yasseryazid/boilerpart/internal/infra/db"
	"github.com/yasseryazid/boilerpart/internal/infra/repository/talent/memory"
	transporthttp "github.com/yasseryazid/boilerpart/internal/transport/http"
	"github.com/yasseryazid/boilerpart/internal/transport/http/health"
	"github.com/yasseryazid/boilerpart/internal/transport/http/talent"
	usecase "github.com/yasseryazid/boilerpart/internal/usecase/talent"
)

func main() {
	dbConn, err := db.OpenFromEnv()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	talentRepo := memory.NewTalentRepository()
	talentService := usecase.NewService(talentRepo)
	r := transporthttp.NewRouter(
		health.NewRegistrar(dbConn),
		talent.NewRegistrar(talentService),
	)
	r.Run(":6060")
}
