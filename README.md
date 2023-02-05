# go-hyperv

A Go module to get Microsoft Hyper-V VM informations and status via [WMI: Windows Management Instrumentation](https://learn.microsoft.com/en-us/windows/win32/wmisdk/about-wmi)

## Usage

Get VM list:

```go
package main

import (
	"fmt"
	"os"

	"github.com/sheepla/hyperv"
)

func main() {
	vms, err := hyperv.GetVMList()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	for _, vm := range *vms {
		fmt.Printf(
			"----------------------\n%v %v [%v]\n\tProcessors:\t%v\n\tMemory:\tusage=%v available=%v buffer=%v\n\tUptime:\t%v\n\tvSwitch:\t%v\n",
			vm.ID,
			vm.Name,
			vm.State,
			vm.NumberOfProcessors,
			vm.MemoryUsage, vm.MemoryAvailable, vm.AvailableMemoryBuffer,
			vm.UpTime,
			vm.VirtualSwitchNames,
		)
	}
}
```

<details>

<summary>Output:</summary>

```
----------------------
354054C8-AE69-4ECB-BC42-7A63BA2688A4 Rocky [enabled but offline]
        Processors:     1
        Memory: usage=0 available=2147483647 buffer=2147483647
        Uptime: 0
        vSwitch:        [ExternalSwitch]
----------------------
782EC864-9404-4AFD-B5C7-58AA6EEBBC24 WS2022 [enabled but offline]
        Processors:     1
        Memory: usage=0 available=2147483647 buffer=2147483647
        Uptime: 0
        vSwitch:        [ExternalSwitch]
----------------------
96948A58-D987-4A71-9DCC-4E125BA48A4E Debian [running]
        Processors:     1
        Memory: usage=1024 available=26 buffer=180
        Uptime: 47863518
        vSwitch:        [ExternalSwitch]
----------------------
E25CD86E-9F94-43CD-B182-33B7CC74E957 ArchLinux [enabled but offline]
        Processors:     4
        Memory: usage=0 available=2147483647 buffer=2147483647
        Uptime: 0
        vSwitch:        [ExternalSwitch]
----------------------
EDF2EDF5-61A4-4AC2-8A2B-BDE67DE4FD12 WS2022Desktop [running]
        Processors:     1
        Memory: usage=4096 available=71 buffer=1236
        Uptime: 38991448
        vSwitch:        [ExternalSwitch]
```

<details>

## Installation

```cmd
go get github.com/sheepla/hyperv
```

