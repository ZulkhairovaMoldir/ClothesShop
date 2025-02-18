package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Println("Не удалось загрузить .env файл, используем переменные окружения.")
	}

	if os.Getenv("SECRET_KEY") == "" {
		log.Println("SECRET_KEY не установлен, используем значение по умолчанию.")
		os.Setenv("SECRET_KEY", "MySuperSecretKeyThatShouldBeLongAndRandom123!")
	}
}
