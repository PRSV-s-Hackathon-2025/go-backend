package global

import (
	"github.com/PRSV-Hackathon-2025/go-backend/pkg/setting"
	"github.com/redis/go-redis/v9"
)

var (
	Setting setting.Setting
	Redis * redis.Client
)