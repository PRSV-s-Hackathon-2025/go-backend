package initialize

import (
	"fmt"

	"github.com/PRSV-Hackathon-2025/go-backend/global"
	"github.com/spf13/viper"
)

func LoadSetting() {
	viper := viper.New()

	viper.AddConfigPath("./environment/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}

	if err := viper.Unmarshal(&global.Setting); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
