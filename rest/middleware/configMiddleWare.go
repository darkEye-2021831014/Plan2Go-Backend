package middleware

import "plan2go-backend/config"

type ConfigMiddleware struct {
	cnf *config.Config
}

func NewConfigMiddleware(cnf *config.Config) *ConfigMiddleware {
	return &ConfigMiddleware{
		cnf: cnf,
	}
}