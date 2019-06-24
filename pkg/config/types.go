/*
Copyright © 2018 inwinSTACK.inc

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

package config

type Operator struct {
	Kubeconfig string
	Interval   int
	Retry      int

	// PA flags
	Host           string
	Username       string
	Password       string
	APIKey         string
	MoveType       int
	MoveRule       string
	Vsys           string
	CommitWaitTime int
	ForceCommit    bool
	SyncCommit     bool
	DaNPartial     bool
	PaOPartial     bool
}
