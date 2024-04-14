package main

import "github.com/sema0205/avito-backend-assignment-2024/internal/app"

const configPath = "config/config.yaml"

func main() {
	app.Run(configPath)
}
