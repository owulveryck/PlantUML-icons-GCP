package main

import (
	"fmt"
	"image"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/sunshineplan/imgconv"
)

func main() {
	err := filepath.WalkDir(".", createPlantuml)
	if err != nil {
		log.Fatal(err)
	}
}

func createPlantuml(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if filepath.Ext(path) != ".png" {
		return nil
	}
	outputFile := strings.TrimSuffix(path, filepath.Ext(path)) + ".puml"
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()
	name, err := createSprite(path, output)
	if err != nil {
		log.Fatal(err)
	}
	err = createEntity(name, output)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func createSprite(filename string, w io.Writer) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	basename := filepath.Base(filename)
	name := strings.TrimSuffix(basename, filepath.Ext(basename))

	image, _, err := image.Decode(f)
	if err != nil {
		return "", err
	}
	dst := imgconv.ToGray(imgconv.Resize(image, imgconv.ResizeOption{Width: 64}))
	bounds := dst.Bounds()
	fmt.Fprintf(w, "sprite $%v [%vx%v/16] {\n", name, bounds.Max.Y, bounds.Max.X)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c, _, _, _ := dst.At(x, y).RGBA()
			fmt.Fprintf(w, "%X", (c)>>12)
		}
		fmt.Fprintf(w, "\n")
	}
	fmt.Fprintln(w, "}")
	return name, nil
}

func createEntity(name string, w io.Writer) error {
	tmpl, err := template.New("test").Parse(`
EntityColoring({{.}})
!define {{.}}(e_alias, e_label, e_techn) Entity(e_alias, e_label, e_techn, #036ffc, {{.}}, {{.}})
!define {{.}}(e_alias, e_label, e_techn, e_descr) Entity(e_alias, e_label, e_techn, e_descr, #036ffc, {{.}}, {{.}})
!define {{.}}Participant(p_alias, p_label, p_techn) Participant(p_alias, p_label, p_techn, #036ffc, {{.}}, {{.}})
!define {{.}}Participant(p_alias, p_label, p_techn, p_descr) Participant(p_alias, p_label, p_techn, p_descr, #036ffc, {{.}}, {{.}})
`)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, name)
}
