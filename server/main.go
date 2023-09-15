package main

import (
	"log"

	"github.com/Feruz666/aggr/config"
	"github.com/Feruz666/aggr/repository"
	"github.com/Feruz666/aggr/worker"
	"github.com/hibiken/asynq"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	store, err := repository.NewStore()
	if err != nil {
		log.Println(" [repository.NewStore] ", " error ", err)
	}
	p := worker.TaskProcessor{Store: store}

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: cfg.Rds.RedisAddr},
		asynq.Config{
			//Specify how many concurrent workers to use
			Concurrency: 8,
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(worker.TypeTaskMsg, p.HandleMessageSendTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}

}
