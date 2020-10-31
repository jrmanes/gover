package main

import (
    "fmt"
    "log"
    "bytes"
    "strings"
    "strconv"
    "os/exec"
    "io/ioutil"

    "gopkg.in/yaml.v2"
)

type Config struct {
        Version string `yaml:"version"`
        Language string `yaml:"language"`
}

type SemanticVersion struct {
	Major int 
	Minor int
	Patch int
}

func (v *Config)getConf() *Config {
	defaultFileName := "gover.yml"
        yamlFile, err := ioutil.ReadFile(defaultFileName)
        if err != nil {
                log.Println("Error:", err)
        }
	err = yaml.Unmarshal(yamlFile, v)
	if err != nil {
		log.Fatalf("Cannot Unmarshal file: %v", err)
	}
	return v
}

func (v *Config)GetVersion() SemanticVersion {
	var major, minor, patch int
	s := strings.Split(v.Version, ".")

//	fmt.Println(s)
	fmt.Printf("Current version: %s\n", v.Version)
	
	major, _ = strconv.Atoi(s[0])
	minor, _ = strconv.Atoi(s[1])
	patch, _ = strconv.Atoi(s[2])
	
	patch += 1

	sm := SemanticVersion{
		major,
		minor,
		patch,
	}
	
	return sm
}

func ExecGitCommand() {
    var out bytes.Buffer
	cmd := exec.Command("git", "log -n 1 --pretty=format:%s $hash")

	err := cmd.Run()
	if err != nil {
	    // something went wrong
	}

	cmd.Stdout = &out
	fmt.Printf("%s\n", out)

}

func main() {
	var v Config
	var sm SemanticVersion

	v.getConf()
//	fmt.Printf("Version: %s - Language: %s", v.Version, v.Language)

	sm = v.GetVersion()
	fmt.Println("Next Version: ", sm.Major)
	fmt.Println("Next Version: ", sm.Minor)
	fmt.Println("Next Version: ", sm.Patch)
	ExecGitCommand()
}
