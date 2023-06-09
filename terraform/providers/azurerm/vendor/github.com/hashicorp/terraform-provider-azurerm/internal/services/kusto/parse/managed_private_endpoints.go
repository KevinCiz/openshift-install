package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type ManagedPrivateEndpointsId struct {
	SubscriptionId             string
	ResourceGroup              string
	ClusterName                string
	ManagedPrivateEndpointName string
}

func NewManagedPrivateEndpointsID(subscriptionId, resourceGroup, clusterName, managedPrivateEndpointName string) ManagedPrivateEndpointsId {
	return ManagedPrivateEndpointsId{
		SubscriptionId:             subscriptionId,
		ResourceGroup:              resourceGroup,
		ClusterName:                clusterName,
		ManagedPrivateEndpointName: managedPrivateEndpointName,
	}
}

func (id ManagedPrivateEndpointsId) String() string {
	segments := []string{
		fmt.Sprintf("Managed Private Endpoint Name %q", id.ManagedPrivateEndpointName),
		fmt.Sprintf("Cluster Name %q", id.ClusterName),
		fmt.Sprintf("Resource Group %q", id.ResourceGroup),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Managed Private Endpoints", segmentsStr)
}

func (id ManagedPrivateEndpointsId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Kusto/Clusters/%s/ManagedPrivateEndpoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.ClusterName, id.ManagedPrivateEndpointName)
}

// ManagedPrivateEndpointsID parses a ManagedPrivateEndpoints ID into an ManagedPrivateEndpointsId struct
func ManagedPrivateEndpointsID(input string) (*ManagedPrivateEndpointsId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := ManagedPrivateEndpointsId{
		SubscriptionId: id.SubscriptionID,
		ResourceGroup:  id.ResourceGroup,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.ResourceGroup == "" {
		return nil, fmt.Errorf("ID was missing the 'resourceGroups' element")
	}

	if resourceId.ClusterName, err = id.PopSegment("Clusters"); err != nil {
		return nil, err
	}
	if resourceId.ManagedPrivateEndpointName, err = id.PopSegment("ManagedPrivateEndpoints"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
