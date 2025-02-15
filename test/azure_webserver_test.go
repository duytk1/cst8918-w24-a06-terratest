package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAzureLinuxVMCreation(t *testing.T) {
	t.Parallel()

	subscriptionID := "6799aad2-f70b-451b-981c-b85924560ee6"
	resourceGroupName := "cst8918-rg"
	vmName := "VM1"

	// Verify VM exists
	vm, err := azure.GetVirtualMachineE(subscriptionID, resourceGroupName, vmName)
	require.NoError(t, err)
	assert.NotNil(t, vm)

	// Confirm NIC exists and is connected to the VM
	nicName := vm.NetworkProfile.NetworkInterfaces[0].ID
	nic, err := azure.GetNetworkInterfaceE(subscriptionID, resourceGroupName, nicName)
	require.NoError(t, err)
	assert.NotNil(t, nic)

	// Confirm the VM is running the correct Ubuntu version
	expectedOS := "Ubuntu"
	assert.Contains(t, *vm.StorageProfile.ImageReference.Offer, expectedOS, "VM is not running the expected OS")
}
