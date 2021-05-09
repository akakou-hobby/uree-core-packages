package uree_core_package

import (
	"fmt"
	"os/exec"
	"strings"

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

type FileSidePallet struct {
}

func (pkg FileSidePallet) Run(req uree_package.Request) uree_package.Response {
	divided := strings.Split(req.Path, "/")
	divided = divided[:len(divided)-1]

	parent := strings.Join(divided, "/")
	fmt.Print(parent)

	out, _ := exec.Command("tree", parent).Output()

	filtered_out := strings.Replace(string(out[:10000]), " ", "&ensp;", -1)
	filtered_out = strings.Replace(filtered_out, "\n", "<br>", -1)

	return uree_package.Response{filtered_out}
}

func (pkg FileSidePallet) SetUpOptional() string {
	return ""
}

func (pkg FileSidePallet) GetName() string {
	return ""
}

func (pkg FileSidePallet) GetIconPath() string {
	return "./img/file.png"
}
