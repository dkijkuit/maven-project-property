package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Project struct {
	XMLName     xml.Name `xml:"project"`
	Name        string   `xml:"name"`
	GroupId     string   `xml:"groupId"`
	ArtifactId  string   `xml:"artifactId"`
	Description string   `xml:"description"`
	Version     string   `xml:"version"`
}

func main() {
	var pomFile string
	var property string = "version"
	var cutoffChar string

	flag.StringVar(&pomFile, "f", "pom.xml", "File location of pom.xml")
	flag.StringVar(&cutoffChar, "c", "-", "Cutoff character for removing trailing text")
	enumFlag(&property, "p", []string{"groupId", "version", "artifactId", "name", "description"}, "Project property to retreive (default \"version\")")

	flag.Parse()

	project, err := getProject(pomFile)
	if err != nil {
		log.Fatalln(err)
		os.Exit(2)
	}

	value := getProperty(project, property, cutoffChar)
	fmt.Println(value)
}

func getProject(pomFile string) (Project, error) {
	xmlFile, err := os.Open(pomFile)

	if err != nil {
		return Project{}, err
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	project := Project{}
	xml.Unmarshal(byteValue, &project)

	return project, err
}

func getProperty(project Project, property string, cutoffChar string) string {
	var value string
	switch property {
	case "groupId":
		value = project.GroupId
	case "artifactId":
		value = project.ArtifactId
	case "version":
		value = project.Version
	case "name":
		value = project.Name
	case "description":
		value = project.Description
	}

	if strings.Contains(value, cutoffChar) {
		value = strings.Split(value, cutoffChar)[0]
	}

	return value
}

func enumFlag(target *string, name string, safelist []string, usage string) {
	flag.Func(name, usage, func(flagValue string) error {
		for _, allowedValue := range safelist {
			if flagValue == allowedValue {
				*target = flagValue
				return nil
			}
		}

		return fmt.Errorf("must be one of %v", safelist)
	})
}
