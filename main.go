package main

import (
	"soal-eksplorasi/configs"
	"soal-eksplorasi/routes"
)

func main() {
	configs.InitDBConnection()

	e := routes.New()
	e.Start(":8080")
}
