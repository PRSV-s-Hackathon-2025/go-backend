package router

import "github.com/PRSV-Hackathon-2025/go-backend/internal/router/public"

type RouterGroup struct {
	Public public.PublicRouter
}

var RouterGroupApp = new(RouterGroup)