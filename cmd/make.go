package cmd

import (
	"fmt"
	"strings"

	"github.com/mohsenMj/go-starter-kit/app/scafolds"
	"github.com/mohsenMj/go-starter-kit/utils"
)

func Make(args []string) {
	fmt.Println("================= Start Making =================")

	model := args[2]

	utils.CreateFile("app/models/"+strings.ToLower(model)+".model.go", scafolds.Apply("model.scafold.txt", model))
	utils.CreateFile("app/requests/"+strings.ToLower(model)+".request.go", scafolds.Apply("request.scafold.txt", model))
	utils.CreateFile("app/responses/"+strings.ToLower(model)+".response.go", scafolds.Apply("response.scafold.txt", model))
	utils.CreateFile("app/repositories/"+strings.ToLower(model)+".repositorie.go", scafolds.Apply("repositorie.scafold.txt", model))
	utils.CreateFile("app/services/"+strings.ToLower(model)+".service.go", scafolds.Apply("service.scafold.txt", model))
	utils.CreateFile("app/controllers/"+strings.ToLower(model)+".controller.go", scafolds.Apply("controller.scafold.txt", model))
	// utils.CreateFile("app/models/"+strings.ToLower(model)+".model.go", utils.ReadFile("app/scafolds/model.scafold.txt", "$model", model))
	// utils.CreateFile("app/requests/"+strings.ToLower(model)+".model.go", utils.ReadFile("app/scafolds/request.scafold.txt", "$model", model))
	// utils.CreateFile("app/responses/"+strings.ToLower(model)+".model.go", utils.ReadFile("app/scafolds/response.scafold.txt", "$model", model))
	// utils.CreateFile("app/repositories/"+strings.ToLower(model)+".model.go", utils.ReadFile("app/scafolds/repository.scafold.txt", "$model", model))
	// utils.CreateFile("app/services/"+strings.ToLower(model)+".model.go", utils.ReadFile("app/scafolds/service.scafold.txt", "$model", model))
	// utils.CreateFile("app/controllers/"+strings.ToLower(model)+".model.go", utils.ReadFile("app/scafolds/controller.scafold.txt", "$model", model))

	fmt.Println("================= Making Done =================")
}
