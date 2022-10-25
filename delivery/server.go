package delivery

import (
	"fmt"
	"warung_nasi_padang/config"
	"warung_nasi_padang/delivery/controller"
	"warung_nasi_padang/delivery/middleware"
	"warung_nasi_padang/manager"
	"warung_nasi_padang/usecase"
	"warung_nasi_padang/utils/authenticator"

	"github.com/gin-gonic/gin"
)

type Server struct {
	useCaseManager manager.UseCaseManager
	authUseCase usecase.AuthUseCase
	tokenService authenticator.AccessToken
	engine		*gin.Engine
	host		string
}

func (s *Server) initController() {
	tokenMdw := middleware.NewTokenValidator(s.tokenService)
	controller.NewMenuController(s.engine, s.useCaseManager.MenuUseCase(),s.authUseCase,tokenMdw)
	controller.NewTransaksiController(s.engine, s.useCaseManager.TransaksiUseCase(),tokenMdw)
}

func(s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	c := config.NewConfig()
	r := gin.Default()
	token := authenticator.NewAccessToken(c.TokenConfig)
	authUseCase := usecase.NewAuthUseCase(token)
	infra := manager.NewInfraManager(c)
	repo := manager.NewRepositoryManager(infra)
	usecase := manager.NewUseCaseManager(repo)
	if c.ApiHost == "" || c.ApiPort == "" {
		panic("No Host or Port define")
	}

	host := fmt.Sprintf("%s:%s", c.ApiHost, c.ApiPort)
	return &Server{useCaseManager: usecase, engine: r, host: host, authUseCase: authUseCase, tokenService: token}
}