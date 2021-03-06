// Copyright 2012 Darren Elwood <darren@textnode.com>
//
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

// This Go code acts as a fallback for processor architectures other than x86.
//

package timestampCounter

func Counter(result *uint64)

func counter() {
    panic("Unsupported on this architecture")
}
