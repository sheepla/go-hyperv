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
			"----------------------\n%v %v [%v]\n"+
				"\tProcessors:\t%v\n"+
				"\tMemory:\t\tusage=%v available=%v buffer=%v\n"+
				"\tUptime:\t\t%v\n"+
				"\tvSwitch:\t%v\n",
			vm.ID, vm.Name, vm.State,
			vm.NumberOfProcessors,
			vm.MemoryUsage, vm.MemoryAvailable, vm.AvailableMemoryBuffer,
			vm.UpTime,
			vm.VirtualSwitchNames,
		)
	}
}
