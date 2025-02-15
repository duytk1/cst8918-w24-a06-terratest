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

	// Ensure NetworkInterfaces is not nil
	require.NotNil(t, vm.NetworkProfile.NetworkInterfaces)
	require.Greater(t, len(*vm.NetworkProfile.NetworkInterfaces), 0, "No network interfaces found for VM")

	// Correctly dereference the network interface slice
	nicID := (*vm.NetworkProfile.NetworkInterfaces)[0].ID
	require.NotNil(t, nicID, "NIC ID is nil")

	// Fetch Network Interface
	nic, err := azure.GetNetworkInterfaceE(subscriptionID, resourceGroupName, *nicID)
	require.NoError(t, err)
	assert.NotNil(t, nic)

	// Confirm the VM is running the correct Ubuntu version
	expectedOS := "Ubuntu"
	require.NotNil(t, vm.StorageProfile.ImageReference.Offer)
	assert.Contains(t, *vm.StorageProfile.ImageReference.Offer, expectedOS, "VM is not running the expected OS")
}
