package commonids

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = CloudServicesIPConfigurationId{}

// CloudServicesIPConfigurationId is a struct representing the Resource ID for a Cloud Services I P Configuration
type CloudServicesIPConfigurationId struct {
	SubscriptionId       string
	ResourceGroup        string
	CloudServiceName     string
	RoleInstanceName     string
	NetworkInterfaceName string
	IpConfigurationName  string
}

// NewCloudServicesIPConfigurationID returns a new CloudServicesIPConfigurationId struct
func NewCloudServicesIPConfigurationID(subscriptionId string, resourceGroup string, cloudServiceName string, roleInstanceName string, networkInterfaceName string, ipConfigurationName string) CloudServicesIPConfigurationId {
	return CloudServicesIPConfigurationId{
		SubscriptionId:       subscriptionId,
		ResourceGroup:        resourceGroup,
		CloudServiceName:     cloudServiceName,
		RoleInstanceName:     roleInstanceName,
		NetworkInterfaceName: networkInterfaceName,
		IpConfigurationName:  ipConfigurationName,
	}
}

// ParseCloudServicesIPConfigurationID parses 'input' into a CloudServicesIPConfigurationId
func ParseCloudServicesIPConfigurationID(input string) (*CloudServicesIPConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(CloudServicesIPConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CloudServicesIPConfigurationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroup, ok = parsed.Parsed["resourceGroup"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroup' was not found in the resource id %q", input)
	}

	if id.CloudServiceName, ok = parsed.Parsed["cloudServiceName"]; !ok {
		return nil, fmt.Errorf("the segment 'cloudServiceName' was not found in the resource id %q", input)
	}

	if id.RoleInstanceName, ok = parsed.Parsed["roleInstanceName"]; !ok {
		return nil, fmt.Errorf("the segment 'roleInstanceName' was not found in the resource id %q", input)
	}

	if id.NetworkInterfaceName, ok = parsed.Parsed["networkInterfaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'networkInterfaceName' was not found in the resource id %q", input)
	}

	if id.IpConfigurationName, ok = parsed.Parsed["ipConfigurationName"]; !ok {
		return nil, fmt.Errorf("the segment 'ipConfigurationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseCloudServicesIPConfigurationIDInsensitively parses 'input' case-insensitively into a CloudServicesIPConfigurationId
// note: this method should only be used for API response data and not user input
func ParseCloudServicesIPConfigurationIDInsensitively(input string) (*CloudServicesIPConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(CloudServicesIPConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CloudServicesIPConfigurationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroup, ok = parsed.Parsed["resourceGroup"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroup' was not found in the resource id %q", input)
	}

	if id.CloudServiceName, ok = parsed.Parsed["cloudServiceName"]; !ok {
		return nil, fmt.Errorf("the segment 'cloudServiceName' was not found in the resource id %q", input)
	}

	if id.RoleInstanceName, ok = parsed.Parsed["roleInstanceName"]; !ok {
		return nil, fmt.Errorf("the segment 'roleInstanceName' was not found in the resource id %q", input)
	}

	if id.NetworkInterfaceName, ok = parsed.Parsed["networkInterfaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'networkInterfaceName' was not found in the resource id %q", input)
	}

	if id.IpConfigurationName, ok = parsed.Parsed["ipConfigurationName"]; !ok {
		return nil, fmt.Errorf("the segment 'ipConfigurationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateCloudServicesIPConfigurationID checks that 'input' can be parsed as a Cloud Services I P Configuration ID
func ValidateCloudServicesIPConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCloudServicesIPConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Cloud Services I P Configuration ID
func (id CloudServicesIPConfigurationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/cloudServices/%s/roleInstances/%s/networkInterfaces/%s/ipConfigurations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.CloudServiceName, id.RoleInstanceName, id.NetworkInterfaceName, id.IpConfigurationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Cloud Services I P Configuration ID
func (id CloudServicesIPConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("subscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("resourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroup", "example-resource-group"),
		resourceids.StaticSegment("providers", "providers", "providers"),
		resourceids.ResourceProviderSegment("resourceProvider", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("cloudServices", "cloudServices", "cloudServices"),
		resourceids.UserSpecifiedSegment("cloudServiceName", "cloudServiceValue"),
		resourceids.StaticSegment("roleInstances", "roleInstances", "roleInstances"),
		resourceids.UserSpecifiedSegment("roleInstanceName", "roleInstanceValue"),
		resourceids.StaticSegment("networkInterfaces", "networkInterfaces", "networkInterfaces"),
		resourceids.UserSpecifiedSegment("networkInterfaceName", "networkInterfaceValue"),
		resourceids.StaticSegment("ipConfigurations", "ipConfigurations", "ipConfigurations"),
		resourceids.UserSpecifiedSegment("ipConfigurationName", "ipConfigurationValue"),
	}
}

// String returns a human-readable description of this Cloud Services I P Configuration ID
func (id CloudServicesIPConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group: %q", id.ResourceGroup),
		fmt.Sprintf("Cloud Service Name: %q", id.CloudServiceName),
		fmt.Sprintf("Role Instance Name: %q", id.RoleInstanceName),
		fmt.Sprintf("Network Interface Name: %q", id.NetworkInterfaceName),
		fmt.Sprintf("Ip Configuration Name: %q", id.IpConfigurationName),
	}
	return fmt.Sprintf("Cloud Services I P Configuration (%s)", strings.Join(components, "\n"))
}
