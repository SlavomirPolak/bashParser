/*
Copyright 2012 Google Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package bashParser

/*
Read file using inputFromFileToString, then use go-shlex for parsing
shell commands from string variable and return map of variables(key)
with its content (value of key).
*/

import (
	"strings"
	"io/ioutil"
	
	"github.com/flynn-archive/go-shlex"
)

func inputFromFileToString(path string) (string, error) {
	dat, err := ioutil.ReadFile(path)
    if err != nil {
		return "", err
	}
    return string(dat), nil
}

func UseShlex(path string) (map[string]string, error) {	
	bashString, err := inputFromFileToString(path)
	if err != nil {
		return nil, err
	}

	keyValueArray, err := shlex.Split(bashString)
	if err != nil {
		return nil, err
	}
	
	var keyValueMap map[string] string
	keyValueMap = make(map[string]string)
	
	for i := 0; i< len(keyValueArray); i++ {
		if strings.Contains(keyValueArray[i], "=") {
			keyValue := strings.SplitAfterN(keyValueArray[i], "=", 2)
			keyValue[0] = keyValue[0][:len(keyValue[0])-1]
			if keyValue[0] == strings.ToUpper(keyValue[0]) && keyValue[0] != "" {
				keyValueMap[keyValue[0]] = keyValue[1]
			}
		}
	}
	
	return keyValueMap, nil
}