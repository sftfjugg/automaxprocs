// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package runtime

// CPUQuotaStatus presents the status of how CPU quota is used
type CPUQuotaStatus int

// TotalMemoryStatus presents the status of how total memory quota is used
type TotalMemoryStatus int

const (
	// CPUQuotaUndefined is returned when CPU quota is undefined
	CPUQuotaUndefined CPUQuotaStatus = iota
	// CPUQuotaUsed is returned when a valid CPU quota can be used
	CPUQuotaUsed
	// CPUQuotaMinUsed is return when CPU quota is smaller than the min value
	CPUQuotaMinUsed

	// TotalMemoryUndefined is returned when total memory is undefined
	TotalMemoryUndefined TotalMemoryStatus = iota
	// TotalMemoryUsed is returned when a valid total memory quota can be used
	TotalMemoryUsed
)
