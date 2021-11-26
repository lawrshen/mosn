/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package transcoder

import (
	"encoding/json"

	"mosn.io/api"
)

type config struct {
	Type  string                 `json:"type,omitempty"`
	Trans map[string]interface{} `json:"trans,omitempty"`
}

func (c *config) GetPhase(key string) api.ReceiverFilterPhase {
	phase, ok := c.Trans[key].(float64)
	if !ok {
		return api.AfterRoute
	}
	if api.ReceiverFilterPhase(phase) <= api.AfterChooseHost && api.ReceiverFilterPhase(phase) >= api.BeforeRoute {
		return api.ReceiverFilterPhase(phase)
	}
	// If receiver_phase does not exist in the configuration, set the default api.AfterRoute value
	return api.AfterRoute
}

func parseConfig(cfg interface{}) (*config, error) {
	filterConfig := &config{}
	data, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, filterConfig); err != nil {
		return nil, err
	}
	return filterConfig, nil
}
