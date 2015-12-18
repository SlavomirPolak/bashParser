package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	
	"github.com/SlavomirPolak/bashParser/src/bashParser"
	//"github.com/coreos/coreos-cloudinit/datasource"
)

const (
	openstackApiVersion = "latest"
)

type configDrive struct {
	root string
	readFile func(filename string) ([]byte, error)
	variablesMap map[string]string
	err error
	
}

func NewDatasource(root string) *configDrive {
	variablesMap, err := bashParser.UseShlex(root + "testFile.sh")
	if err != nil {
		log.Printf("Error during parsing script file.\n")
	}
	return &configDrive{root, ioutil.ReadFile, variablesMap, err}
}

func (cd *configDrive) IsAvailable() bool {
	_, err := os.Stat(cd.root)
	return !os.IsNotExist(err)
}

func (cd *configDrive) AvailabilityChanges() bool {
	return true
}

func (cd *configDrive) ConfigRoot() string {
	return cd.root
}

/*preskakujem openstackVersionRoot a openstackRoot
func (cd *configDrive) openstackRoot() string {
	return cd.root
}*/

func (cd *configDrive) FetchMetadata() ([]byte, error) {
	var metadata struct {
		SSH_KEY []byte
	}
	
	// searching for SSH_PUBLIC_KEY or SSH_KEY or PUBLIC_SSH_KEY
	var val string
	if cd.variablesMap["SSH_PUBLIC_KEY"] != "" {
		val = cd.variablesMap["SSH_PUBLIC_KEY"]
	} else if cd.variablesMap["SSH_KEY"] != "" {
		val = cd.variablesMap["SSH_KEY"]
	} else if cd.variablesMap["PUBLIC_SSH_KEY"] != "" {
		val = cd.variablesMap["PUBLIC_SSH_KEY"]
	}
	if val != "" {
		metadata.SSH_KEY = []byte(val)
	} else {
		log.Printf("Variable USER_DATA isnt in script file.\n")
	}
	return metadata.SSH_KEY, cd.err
}

func Type() string {
	return "cloud-drive"
}

func (cd *configDrive) FetchUserdata() ([]byte, error) {
	if cd.variablesMap["USER_DATA"] == "" {
		log.Printf("Variable USER_DATA isnt in script file.\n")
		return nil, cd.err
	}
	ret := cd.variablesMap["USER_DATA"]
	return []byte(ret), cd.err
}

func NewVariablesMap(fileName string) (map[string]string, error) {
	variablesMap, err := bashParser.UseShlex(fileName)
	if err != nil {
		log.Printf("Error during parsing script file.")
		return nil, err
	}
	
	return variablesMap, nil
}

func main() {
	ds := NewDatasource("/home/wolfik/gocode/src/bashParser/data/")
	userData, err := ds.FetchUserdata()
	
	fmt.Println(userData, err)
}