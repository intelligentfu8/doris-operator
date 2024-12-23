// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package resource

import "testing"

func Test_GetPort(t *testing.T) {
	config := map[string]interface{}{}
	config["http_port"] = "80030"
	config["rpc_port"] = "test"
	tks := []string{"http_port", "rpc_port"}
	rk := map[string]int32{
		"http_port": 80030,
		"rpc_port":  9020,
	}
	for _, key := range tks {
		res := GetPort(config, key)
		if res != rk[key] {
			t.Errorf("")
		}
	}
}

func Test_GetTerminationGracePeriodSeconds(t *testing.T) {
	tests := []map[string]interface{}{
		{
			"grace_shutdown_wait_seconds": "60",
		},
		{
			"test_shutdown": "10",
		},
	}

	ress := []int64{60, 0}
	for i, test := range tests {
		res := GetTerminationGracePeriodSeconds(test)
		if res != ress[i] {
			t.Errorf("test TerminationGracePeriodSeconds failed, intput not equal expected")
		}
	}
}
