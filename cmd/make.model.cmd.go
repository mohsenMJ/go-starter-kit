package cmd

import (
	"fmt"
	"strings"

	"github.com/mohsenMj/go-starter-kit/app/scafolds"
	"github.com/mohsenMj/go-starter-kit/utils"
)

func CmdMakeModelHandler(args []string) {
	fmt.Println("================= Start Making =================")

	model := args[2]

	utils.CreateFile("app/models/"+strings.ToLower(model)+".model.go", scafolds.Apply("model.scafold.txt", model))
	utils.CreateFile("app/requests/"+strings.ToLower(model)+".request.go", scafolds.Apply("request.scafold.txt", model))
	utils.CreateFile("app/responses/"+strings.ToLower(model)+".response.go", scafolds.Apply("response.scafold.txt", model))
	utils.CreateFile("app/repositories/"+strings.ToLower(model)+".repositorie.go", scafolds.Apply("repository.scafold.txt", model))
	utils.CreateFile("app/services/"+strings.ToLower(model)+".service.go", scafolds.Apply("service.scafold.txt", model))
	utils.CreateFile("app/controllers/"+strings.ToLower(model)+".controller.go", scafolds.Apply("controller.scafold.txt", model))

	fmt.Println("================= Making Done =================")
}
