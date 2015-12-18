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
	
}

func NewDatasource(root string) *configDrive {
	return &configDrive{root, ioutil.ReadFile}
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

func (cd *configDrive) FetchUserdata() ([]byte, error) {
	ret, err := fetchVariableFromShellScript(cd.root + "testFile.sh", "USER_DATA")
	return []byte(ret), err
}

func (cd *configDrive) FetchMetadata() ([]byte, error) {
	var metadata struct {
		SSH_KEY []byte
	}
	
	// searching for SSH_PUBLIC_KEY or SSH_KEY or PUBLIC_SSH_KEY
	val, err := fetchVariableFromShellScript(cd.root + "testFile.sh", "SSH_PUBLIC_KEY")
	if val == "" {
		val, err = fetchVariableFromShellScript(cd.root + "testFile.sh", "SSH_KEY")
		if val == "" {
			val, err = fetchVariableFromShellScript(cd.root + "testFile.sh", "PUBLIC_SSH_KEY")
		}
	}
	if val != "" {
		metadata.SSH_KEY = []byte(val)
	}
	return metadata.SSH_KEY, err
}

func Type() string {
	return "cloud-drive"
}

func fetchVariableFromShellScript(fileName string, variableName string) (string, error) {
	log.Printf("Attempting to read " + variableName + " from %q\n", fileName)
	variablesMap, err := bashParser.UseShlex(fileName)
	if os.IsNotExist(err) {
		err = nil
	}
	
	ret := variablesMap[variableName]
	if ret == "" {
		log.Printf("Variable " + variableName + " isnt in script file.\n")		
	}
	
	return ret, nil
}

func main() {
	ds := NewDatasource("/home/wolfik/gocode/src/bashParser/data/")
	k, _ := ds.FetchUserdata()
	g, _ := ds.FetchMetadata()
	m := string(g)
	h := string(k)
	fmt.Println(h, m)
}