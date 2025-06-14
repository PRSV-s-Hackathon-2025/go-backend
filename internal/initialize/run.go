package initialize

import (
	"fmt"

	"github.com/PRSV-Hackathon-2025/go-backend/global"
)

func Run() {
	LoadSetting()
	InitRedis()
	r := InitRouter()

	port := fmt.Sprintf(":%d", global.Setting.Server.Port)

	r.Run(port)
}
