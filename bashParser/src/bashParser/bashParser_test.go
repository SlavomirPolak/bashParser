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

import "testing"

/**
 *	test of keys
 */
func Test_useShlex_1(t *testing.T) {
	keyValueMap := useShlex("/home/wolfik/gocode/src/bashParser/data/testFile.sh")
	
	expectedKeys := []string{"NAME", "EMAIL"}
	keys := make([]string, 0, len(keyValueMap))
	
	for k := range keyValueMap {
		keys = append(keys, k)
	}
	
	for i := 0; i < len(keyValueMap)-1; i++ {
		if expectedKeys[i] != keys[i] {
			t.Error("useShlex did not work properly, there are unexpected variables names in map")
		} else {
			t.Log("useShlex works properly")
		}
	}
}

/**
 *	test of values of keys "NAME" and "EMAIL"
 */
func Test_useShlex_2(t *testing.T) {
	keyValueMap := useShlex("/home/wolfik/gocode/src/bashParser/data/testFile.sh")
	
	if keyValueMap["NAME"] != "Slavo" {
		t.Error("useShlex did not work properly, value of key \"NAME\" is unexpected")
	} else {
		t.Log("useShlex works properly")
	}
	
	if keyValueMap["EMAIL"] != "johndoe@mail.com" {
		t.Error("useShlex did not work properly, value of key \"EMAIL\" is unexpected")
	} else {
		t.Log("useShlex works properly")
	}
}