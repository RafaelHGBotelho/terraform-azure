package frontdoor

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// TODO: move this to the network one

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "FrontDoor"
}

// WebsiteCategories returns a list of categories which can be used for the sidebar
func (r Registration) WebsiteCategories() []string {
	return []string{
		"Network",
	}
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"azurerm_frontdoor":                            resourceFrontDoor(),
		"azurerm_frontdoor_firewall_policy":            resourceFrontDoorFirewallPolicy(),
		"azurerm_frontdoor_custom_https_configuration": resourceFrontDoorCustomHttpsConfiguration(),
	}
}
