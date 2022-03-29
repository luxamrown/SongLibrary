package main

import (
	"enigmacamp.com/song/config"
	"enigmacamp.com/song/delivery/api"
	"github.com/gin-gonic/gin"
)

type AppServer interface {
	Run()
}

type appServer struct {
	routerEngine *gin.Engine
	cfg          config.Config
}

func (p *appServer) initHandlers() {
	p.v1()
}

func (p *appServer) v1() {
	songApiGroup := p.routerEngine.Group("/song")
	api.NewSongApi(songApiGroup, p.cfg.UseCaseManager)
}

func (p *appServer) Run() {
	p.initHandlers()
	err := p.routerEngine.Run(p.cfg.ApiConfig.Url)
	if err != nil {
		panic(err)
	}
}

func Server() AppServer {
	r := gin.Default()
	c := config.NewConfig(".", "config")
	return &appServer{
		routerEngine: r,
		cfg:          c,
	}
}
