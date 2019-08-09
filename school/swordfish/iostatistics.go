// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swordfish

// IOStatistics is used to represent the IO statistics of the requested
// object.
type IOStatistics struct {
	// NonIORequestTime shall be an ISO 8601 conformant duration describing the
	// time that the resource is busy processing non IO requests.
	NonIORequestTime string
	// NonIORequests shall represent the total count from the time of last reset
	// or wrap of non IO requests.
	NonIORequests int64
	// ReadHitIORequests shall represent the total count from the time of last
	// reset or wrap of read IO requests satisfied from memory.
	ReadHitIORequests int64
	// ReadIOKiBytes shall represent the total number of kibibytes read from the
	// time of last reset or wrap.
	ReadIOKiBytes int64
	// ReadIORequestTime shall be an ISO 8601 conformant duration describing the
	// time that the resource is busy processing read requests.
	ReadIORequestTime string
	// ReadIORequests shall represent the total count from the time of last
	// reset or wrap of read IO requests satisfied from either media or memory
	// (i.e. from a storage device or from a cache).
	ReadIORequests int64
	// WriteHitIORequests shall represent the total count from the time of last
	// reset or wrap of write IO requests coallesced into memory.
	WriteHitIORequests int64
	// WriteIOKiBytes shall represent the total number of kibibytes written from
	// the time of last reset or wrap.
	WriteIOKiBytes int64
	// WriteIORequestTime shall be an ISO 8601 conformant duration describing
	// the time that the resource is busy processing write requests.
	WriteIORequestTime string
	// WriteIORequests shall represent the total count from the time of last
	// reset or wrap of write IO requests.
	WriteIORequests int64
}
