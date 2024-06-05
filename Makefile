run:
	go run main.go

folder_init:
	mkdir -p router
	mkdir -p controller
	mkdir -p service
	mkdir -p repository
	mkdir -p model
	mkdir -p database
	mkdir -p docs
	mkdir -p library
	cp .env.example .env

rest_init:
	@echo "\nInstalling Gin.............."
	go get -u github.com/gin-gonic/gin
	@echo "\nInstalling Gorm.............."
	go get -u gorm.io/gorm
	@echo "\nInstalling Driver Postgres.............."
	go get -u gorm.io/driver/postgres
	@echo "Done!"
