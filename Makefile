run:
	go run main.go

folder_init:
	mkdir router
	mkdir controller
	mkdir service
	mkdir repository
	mkdir model
	mkdir database
	mkdir docs
	mkdir library

rest_init:
	@echo "\nInstalling Gin.............."
	go get -u github.com/gin-gonic/gin
	@echo "\nInstalling Gorm.............."
	go get -u gorm.io/gorm
	@echo "\nInstalling Driver Postgres.............."
	go get -u gorm.io/driver/postgres
	@echo "Done!"
