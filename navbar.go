package uree_core_package

import (
	"fmt"
	"os/exec"

	uree_package "github.com/akakou-hobby/uree-package"
	"github.com/sqweek/dialog"
)

type CommpileCPackage struct{}

func (pkg CommpileCPackage) Run(req uree_package.Request) uree_package.Response {
	resp := uree_package.Response{
		req.Body,
	}

	outputPath := fmt.Sprintf("%s.out", req.Path)

	fmt.Printf("gcc %s -o %s", req.Path, outputPath)

	out, err := exec.Command("gcc", req.Path, "-o", outputPath).CombinedOutput()
	if err != nil {
		dialog.Message("Error", string(out)).Title("Compile error").Info()
		return resp
	}

	err = exec.Command("chmod", "+x", outputPath).Run()
	if err != nil {
		dialog.Message("Error", err).Title("chmod Error").Info()
		return resp
	}

	out, err = exec.Command(outputPath).Output()
	dialog.Message("Output", string(out)).Title("Output").Info()

	return resp
}

func (pkg CommpileCPackage) SetUpOptional() string {
	return ""
}

func (pkg CommpileCPackage) GetName() string {
	return "Run C"
}
