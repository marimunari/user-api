package utils

import (
	"net/http"
	"user-api/models"
)

var (
	ErrUserNotFound               = models.ErrorResponse{APIResponse: models.APIResponse{Code: http.StatusNotFound, Message: "Usuário não encontrado"}}
	ErrEmailExists                = models.ErrorResponse{APIResponse: models.APIResponse{Code: http.StatusBadRequest, Message: "Usuário já cadastrado"}}
	ErrInvalidCredentials         = models.ErrorResponse{APIResponse: models.APIResponse{Code: http.StatusUnauthorized, Message: "Credenciais inválidas"}}
	ErrFailedToParse              = models.ErrorResponse{APIResponse: models.APIResponse{Code: http.StatusBadRequest, Message: "Falha ao analisar o corpo da solicitação"}}
	ErrFailedToCreate             = models.ErrorResponse{APIResponse: models.APIResponse{Code: http.StatusInternalServerError, Message: "Falha ao criar usuário"}}
	ErrFailedToGenerateToken      = models.ErrorResponse{APIResponse: models.APIResponse{Code: http.StatusInternalServerError, Message: "Falha ao gerar token"}}
	ErrFailedToHashPassword       = models.ErrorResponse{APIResponse: models.APIResponse{Code: http.StatusInternalServerError, Message: "Falha ao hash da senha"}}
	ErrAuthorizationHeaderMissing = models.ErrorResponse{APIResponse: models.APIResponse{Code: http.StatusUnauthorized, Message: "Cabeçalho de autorização ausente"}}
	ErrUnauthorized               = models.ErrorResponse{APIResponse: models.APIResponse{Code: http.StatusUnauthorized, Message: "Não autorizado"}}
	ErrTokenNotProvided           = models.ErrorResponse{APIResponse: models.APIResponse{Code: http.StatusUnauthorized, Message: "Token não fornecido"}}
	ErrTokenBlacklisted           = models.ErrorResponse{APIResponse: models.APIResponse{Code: http.StatusUnauthorized, Message: "Token está na lista negra"}}
	ErrInvalidToken               = models.ErrorResponse{APIResponse: models.APIResponse{Code: http.StatusUnauthorized, Message: "Token inválido"}}
)

var (
	SuccessUserRegistered = models.SuccessResponse{APIResponse: models.APIResponse{Code: http.StatusCreated, Message: "Usuário registrado com sucesso"}}
	SuccessLogin          = models.SuccessResponse{APIResponse: models.APIResponse{Code: http.StatusOK, Message: "Login bem-sucedido"}}
	SuccessLogout         = models.SuccessResponse{APIResponse: models.APIResponse{Code: http.StatusOK, Message: "Logout bem-sucedido"}}
	SuccessGetUserDetail  = models.SuccessResponse{APIResponse: models.APIResponse{Code: http.StatusOK, Message: "Detalhes do usuário obtidos com sucesso"}}
)
