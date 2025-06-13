package handler

import "gin-project/internal/usecase"

type ApiHandler struct {
	ApiUsecase usecase.CardUsecase
}
