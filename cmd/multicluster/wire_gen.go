// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
	"multicluster/internal/biz"
	"multicluster/internal/conf"
	"multicluster/internal/data"
	"multicluster/internal/server"
	"multicluster/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	client := data.NewEntClient(confData, logger)
	dataData, cleanup, err := data.NewData(client, logger)
	if err != nil {
		return nil, nil, err
	}
	clusterRepo := data.NewClusterRepo(dataData, logger)
	clusterUsecase := biz.NewClusterUsecase(clusterRepo, logger)
	componentRepo := data.NewComponentRepo(dataData, logger)
	componentUsecase := biz.NewComponentUsecase(componentRepo, logger)
	multiClusterService := service.NewMultiClusterService(clusterUsecase, componentUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, multiClusterService, logger)
	grpcServer := server.NewGRPCServer(confServer, multiClusterService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
