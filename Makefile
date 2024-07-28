
# Default target: display help
.DEFAULT_GOAL := help

# Target to initialize project folders
folder_init:
	mkdir -p router controller
	cp -r builder-maker/database database
	cp -r builder-maker/controller_auth controller
	cp .env.example .env
	@echo "Project folders initialized."

run_local:
	go run main.go

# Define default values for variables
name ?= customer

# Define a variable to hold the module path
module_path := $(shell go list -m)

# Example target that prints the module path
show_path:
	@echo "$(module_path)"

# Convert name to title case (capitalize each word)
title_case_name := $(shell echo "$(name)" | awk '{ \
  n = split(tolower($$0), a, " "); \
  for (i = 1; i <= n; i++) { \
    a[i] = toupper(substr(a[i],1,1)) tolower(substr(a[i],2)); \
  } \
  print a[1]; \
  for (i = 2; i <= n; i++) { \
    printf " %s", a[i]; \
  } \
  print ""; \
}')

# Target to create temp.txt based on the name variable
temp.txt:
	@printf '    %sCtrl := $(name).New%sController(db.GormDb)\n' "$(name)" "$(title_case_name)" > temp.txt
	@printf '    %sRoutes := router.Group("$(name)")\n' "$(name)" >> temp.txt
	@printf '    {\n' >> temp.txt
	@printf '        %sRoutes.POST("/add", %sCtrl.Register)\n' "$(name)" "$(name)" >> temp.txt
	@printf '        %sRoutes.GET("/detail", %sCtrl.Login)\n' "$(name)" "$(name)" >> temp.txt
	@printf '    }\n' >> temp.txt

import_controller.txt:
	@printf '\n    "%s/controller/%s"\n' "$(module_path)" "$(name)" > import_controller.txt

# Replace Example
replace_controller_example:
	@sed -i '' 's/example/$(name)/g' controller/$(name)/$(name).go
	@sed -i '' 's/Example/$(title_case_name)/g' controller/$(name)/$(name).go
	@echo "Replaced 'example' with '$(name)' in file.txt."


# Target to append routes based on the name variable
resource: temp.txt import_controller.txt
	@awk '/\/\/generated routes/ {print; system("cat temp.txt"); next} 1' router/router.go > temp_router.go && mv temp_router.go router/router.go
	@rm temp.txt
	@awk '/\/\/generated import/ {print; system("cat import_controller.txt"); next} 1' router/router.go > temp_router.go && mv temp_router.go router/router.go
	@rm import_controller.txt
	@echo "Routes for $(name) created successfully."
	cp -r builder-maker/controller/example controller/$(name)
	mv controller/$(name)/example.go controller/$(name)/$(name).go
	$(MAKE) replace_controller_example
	@echo "Controller for $(name) created successfully."

# Help target to display usage information
help:
	@echo "Usage:"
	@echo ""
	@echo " - Initialize project folders: make folder_init"
	@echo " - Run the application: make run_local"
	@echo " - Create a new route: make resource name=\"[resourcename]\""
