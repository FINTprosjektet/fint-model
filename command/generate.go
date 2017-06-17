package command

import (
	"github.com/codegangsta/cli"
	"os"
	"fmt"
	"github.com/FINTprosjektet/fint-model/package"
	"strings"
	"github.com/FINTprosjektet/fint-model/github"
	"github.com/FINTprosjektet/fint-model/document"
	"github.com/FINTprosjektet/fint-model/generator"
	"io/ioutil"
)

const basePath = "java/src/main/java/no/fint/model"

func CmdGenerate(c *cli.Context) {

	var tag string
	if c.GlobalString("tag") == "latest" {
		tag = github.GetLatest()
	} else {
		tag = c.GlobalString("tag")
	}

	if c.String("lang") == "JAVA" {
		generateJavaCode(tag)
	}

	if c.String("lang") == "NET" {
		generateNetCode(tag)
	}

}

func generateJavaCode(tag string) {

	//document.GetFile(tag)
	document.Get(tag)
	fmt.Println("Generating Java code:")
	setupJavaDirStructure(tag)
	classes := generator.GetClasses(tag)
	for _, c := range classes {
		fmt.Printf("  > Creating class: %s.java\n", c.Name)
		var class string

		if len(c.Extends) > 0 {
			class = generator.GetExtendedJavaClass(c)
		} else if c.Abstract {
			class = generator.GetAbstractJavaClass(c)
		} else {
			class = generator.GetJavaClass(c)
		}

		path := fmt.Sprintf("%s/%s/%s.java", basePath, strings.Replace(c.Package, ".", "/", -1), c.Name)
		err := ioutil.WriteFile(path, []byte(class), 0777)
		if err != nil {
			fmt.Printf("Unable to write file: %s", err)
		}
	}

	fmt.Println("Finish generating Java code!")
}

func generateNetCode(tag string) {
	fmt.Println("Not yet implemented")
}

func setupJavaDirStructure(tag string) {
	fmt.Println("  > Setup directory structure.")
	os.RemoveAll("java")
	err := os.MkdirAll(basePath, 0777)
	if err != nil {
		fmt.Println("Unable to create base structure")
		fmt.Println(err)
	}
	for _, pkg := range packages.DistinctPackageList(tag) {
		path := fmt.Sprintf("%s/%s", basePath, strings.Replace(pkg, ".", "/", -1))
		err := os.MkdirAll(path, 0777)
		if err != nil {
			fmt.Println("Unable to create package structure")
			fmt.Println(err)
		}

	}
}
