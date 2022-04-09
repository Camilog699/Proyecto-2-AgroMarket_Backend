package app

import (
	env "github.com/joho/godotenv"
)

// Esto va en la linea 9 env.Load("/home/ubuntu/services/.env") -- Normalmente va con .env
func LoadEnv() {
	err := env.Load(".env")
	if err != nil {
		panic(err)
	}
}
