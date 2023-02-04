package hyperv

import (
	"fmt"
	"time"

	"github.com/bi-zone/wmi"
)

const hypervNamespace = `root\virtualization\v2`

type VM struct {
	ID                              string `wmi:"Name"`
	Name                            string `wmi:"ElementName"`
	InstanceID                      string
	AllocatedGPU                    string
	Shielded                        bool
	CreationTime                    time.Time
	State                           State `wmi:"EnabledState"`
	OtherEnabledState               string
	GuestOperatingSystem            string
	HealthState                     uint16
	Heartbeat                       uint16
	MemoryUsage                     uint64
	MemoryAvailable                 int
	AvailableMemoryBuffer           int
	SwapFilesInUse                  bool
	Notes                           string
	Version                         string
	NumberOfProcessors              uint16
	OperationalStatus               []uint16
	ProcessorLoad                   uint16
	ProcessorLoadHistory            []uint16
	StatusDescriptions              []string
	ThumbnailImage                  []uint8
	ThumbnailImageHeight            uint16
	ThumbnailImageWidth             uint16
	UpTime                          uint64
	ReplicationState                uint16
	ReplicationStateEx              []uint16
	ReplicationHealth               uint16
	ReplicationHealthEx             []uint16
	ReplicationMode                 uint16
	ApplicationHealth               uint16
	IntegrationServicesVersionState uint16
	MemorySpansPhysicalNumaNodes    bool
	ReplicationProviderId           []string
	EnhancedSessionModeState        uint16
	VirtualSwitchNames              []string
	VirtualSystemSubType            string
	HostComputerSystemName          string
	// AsynchronousTasks       CIM_ConcreteJob
	// TestReplicaSystem        CIM_ComputerSystem
	// Snapshots        CIM_VirtualSystemSettingData
}

type State uint16

// https://learn.microsoft.com/en-us/previous-versions//cc136818(v=vs.85)
const (
	StateUnknown State = iota
	StateOther
	StateRunning
	StateOff
	StateShuttingDown
	StateNotApplicable
	StateEnabledButOffline
	StateInTest
	StateDeferred
	StateQuiesce
	StateStarting
)

func (s State) String() string {
	switch s {
	case StateUnknown:
		return "unknown"
	case StateOther:
		return "other"
	case StateRunning:
		return "running"
	case StateOff:
		return "off"
	case StateShuttingDown:
		return "shutting down"
	case StateNotApplicable:
		return "not applicable"
	case StateEnabledButOffline:
		return "enabled but offline"
	case StateInTest:
		return "in test"
	case StateDeferred:
		return "deferred"
	case StateQuiesce:
		return "quiesce"
	case StateStarting:
		return "starting"
	default:
		if 11 <= s && 32767 <= s {
			return "DMTF reserved"
		}

		if 32768 <= s && 65535 <= s {
			return "vendor reserved"
		}

		return ""
	}
}

// c.f. https://learn.microsoft.com/en-us/windows/win32/hyperv_v2/msvm-computersystem
func GetVMList() (*[]VM, error) {
	var vms []VM

	q := wmi.CreateQueryFrom(&vms, "Msvm_SummaryInformation", "")
	if err := wmi.QueryNamespace(q, &vms, hypervNamespace); err != nil {
		return nil, fmt.Errorf("failed to get VM list: %w\n\t- query:\t%s\n\t- namespace:\t%s", err, q, hypervNamespace)
	}

	return &vms, nil
}
