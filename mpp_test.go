package main

import (
	"testing"
)

var project Project = Project{
	Name:        "Test name",
	GroupId:     "test.testdata.net.org.com.test",
	ArtifactId:  "testartifact",
	Version:     "3.1.42",
	Description: "Test description",
}

var projectCutOff Project = Project{
	Name:        "Test name",
	GroupId:     "test.testdata.net.org.com.test",
	ArtifactId:  "testsplit-artifact",
	Version:     "3.1.42-SNAPSHOT",
	Description: "Test description",
}

func TestGetPomFile(t *testing.T) {
	_, err := getProject("./test/pom.xml")
	if err != nil {
		t.Fatal("Unable to read pom.xml file:", err)
	}
}

func TestPomFileNotFound(t *testing.T) {
	_, err := getProject("./test/non_existing_pom.xml")
	if err == nil {
		t.Fatal("Expected not to find the file but no error was thrown")
	}
}

func TestGetAllProperties(t *testing.T) {
	testProperty(t, project, "name", project.Name)
	testProperty(t, project, "groupId", project.GroupId)
	testProperty(t, project, "artifactId", project.ArtifactId)
	testProperty(t, project, "version", project.Version)
	testProperty(t, project, "description", project.Description)
}

func TestGetAllPropertiesCutOff(t *testing.T) {
	testProperty(t, projectCutOff, "name", projectCutOff.Name)
	testProperty(t, projectCutOff, "groupId", projectCutOff.GroupId)
	testProperty(t, projectCutOff, "artifactId", "testsplit")
	testProperty(t, projectCutOff, "version", "3.1.42")
	testProperty(t, projectCutOff, "description", projectCutOff.Description)
}

func testProperty(t *testing.T, projectData Project, property string, expectedValue string) {
	value := getProperty(projectData, property, "-")
	if value == "" || value != expectedValue {
		t.Fatal("Project", property, "mismatch, expected:", expectedValue, "instead we got:", value)
	}
}
