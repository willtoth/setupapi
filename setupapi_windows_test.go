// +build windows

package setupapi

import (
	"fmt"
	"testing"
)

func TestDeviceInfo(t *testing.T) {
	guids, err := SetupDiClassGuidsFromNameEx("Processor", "")
	if err != nil {
		t.Errorf("Error retrieving class guid: %s", err.Error())
	}

	di, err := SetupDiGetClassDevsEx(guids[0], "", 0, Present, 0, "", 0)
	if err != nil {
		t.Errorf("Error get class devs ex: %s", err.Error())
	}

	devPath, err := di.DevicePath(guids[0])
	if err != nil {
		fmt.Printf("Error device path: %s", err.Error())
		return
	}

	did, err := di.EnumDeviceInfo(0)
	if err != nil {
		t.Errorf("Error enum device info: %s", err.Error())
	}

	fmt.Println(devPath)
	fmt.Println(did.InstanceID())
	fmt.Println(did.ClassGuid)
}
