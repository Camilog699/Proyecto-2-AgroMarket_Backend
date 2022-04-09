package godo

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Linea 11 godoenv.Load ("/home/ubuntu/services/.env") -- Normalmente va vacio
func GodoGet(key string) (value string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error cargando el archivo .env")
	}
	value = os.Getenv(key)
	return
}
