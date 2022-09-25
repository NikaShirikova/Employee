package handler

import "go.uber.org/zap"

var LoggerZap *zap.Logger

type statusResponse struct {
	Status string `json:"status"`
}
