package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v2"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Version:               "1.0.0",
	Use:                   fmt.Sprintf("%s [FLAGS] TEMPLATE [TEMPLATE ...]", filepath.Base(os.Args[0])),
	Short:                 "A CLI for the golang template engine",
	DisableFlagsInUseLine: true,
	SilenceUsage:          true,
	RunE:                  runE,
}

var flagYamlData []string

func init() {
	cmd.PersistentFlags().StringArrayVarP(&flagYamlData, "yaml", "y", nil, "Paths to yaml data files")
}

func runE(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		cmd.Help()
		return nil
	}
	var data interface{}
	for _, f := range flagYamlData {
		b, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}
		if err := yaml.Unmarshal(b, &data); err != nil {
			return err
		}
	}
	t := template.New("") // .Funcs(tmpl.Funcs())
	for _, filename := range args {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		if _, err := t.New(filename).Parse(string(b)); err != nil {
			return err
		}
	}
	if err := t.ExecuteTemplate(os.Stdout, args[0], data); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
