package gapi

import (
	"fmt"

	db "github.com/quocbaodoan/simplebank/db/sqlc"
	"github.com/quocbaodoan/simplebank/pb"
	"github.com/quocbaodoan/simplebank/token"
	"github.com/quocbaodoan/simplebank/util"
	"github.com/quocbaodoan/simplebank/worker"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
