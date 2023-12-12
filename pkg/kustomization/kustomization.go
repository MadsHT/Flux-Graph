package kustomization

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Structs for handling the information that comes from the kustomization api's
type Kustomization struct {
	Metadata Metadata `yaml:"metadata"`
	Spec     Spec     `yaml:"spec"`
	HRSlice  []string
}

type Metadata struct {
	Name string `yaml:"name"`
}

type Spec struct {
	DependsOn []DependsOn `yaml:"dependsOn"`
	Path      string      `yaml:"path"`
}

type DependsOn struct {
	Name string `yaml:"name"`
}

var Kustomizations []Kustomization

// Unmarshals a .yaml file into the app struct
func (ks *Kustomization) GetValuesFromYamlFile(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(file, &ks)
}

func (ks *Kustomization) HasDependsOn() bool {
	return ks.Spec.DependsOn != nil
}

// Enclosing name in double quotes since the graph language wants them to be in quotes
func (ks *Kustomization) Name() string {
	return "\"" + ks.Metadata.Name + "\""
}

func (ks *Kustomization) GetDependencies() []string {
	dependencies := []string{}

	for _, v := range ks.Spec.DependsOn {

		// Enclosing name in double quotes since the graph language wants them to be in quotes
		dependencies = append(dependencies, "\""+v.Name+"\"")
	}
	return dependencies
}
