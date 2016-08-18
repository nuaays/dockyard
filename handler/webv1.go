/*
Copyright 2015 The ContainerOps Authors All rights reserved.

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

package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/macaron.v1"
)

//
func IndexV1Handler(ctx *macaron.Context) (int, []byte) {
	discovery := ctx.Query("ac-discovery")

	if len(discovery) > 0 {

	}

	result, _ := json.Marshal(map[string]string{"message": "Dockyard Backend REST API Service"})
	return http.StatusOK, result
}
