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

package fileOpening

import (
	"testing"
	"path/filepath"
	"strings"
	"fmt"
)

/**
 *	test has existing file as parameter
 */
func Test_InputFromFileToString_1(t *testing.T) {

	path, err := filepath.Abs("")
	if err != nil {
		t.Error("Error during searching absolute path.")
	}
	
	if strings.HasSuffix(path, "/src/fileOpening") {
        path = path[:len(path)-len("/src/fileOpening")]
    }
	fmt.Println(path)

	result, err := InputFromFileToString(path + "/data/testFile.sh")

	if result != "# comment\nNAME=Slavo\nfor i = 5\nEMAIL=johndoe@mail.com" && err != nil {
		t.Error("inputFromFile did not work properly,\nerror has occured when parameter was regular text file")
	} else {
		t.Log("inputFromFile works properly with regular text file as parameter.\n")
	}
}

/**
 *	test has non-existing file as parameter
 */
func Test_InputFromFileToString_2(t *testing.T) {
	result, err := InputFromFileToString("testFile2.sh")

	if result != "" && err != nil {
		t.Error("InputFromFileToString did not work properly,\nerror has occured when parameter was file that does not exist")
	} else {
		t.Log("InputFromFileToString works properly when we try to open file that does not exist.\n")
	}
}