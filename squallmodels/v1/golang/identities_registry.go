package squallmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(APICheckIdentity)
	elemental.RegisterIdentity(MapNodeIdentity)
	elemental.RegisterIdentity(DependencyMapSubviewIdentity)
	elemental.RegisterIdentity(APIAuthorizationPolicyIdentity)
	elemental.RegisterIdentity(SyscallAccessIdentity)
	elemental.RegisterIdentity(ComputedPolicyIdentity)
	elemental.RegisterIdentity(TagIdentity)
	elemental.RegisterIdentity(EnforcerIdentity)
	elemental.RegisterIdentity(MapEdgeIdentity)
	elemental.RegisterIdentity(FilePathIdentity)
	elemental.RegisterIdentity(FileAccessIdentity)
	elemental.RegisterIdentity(NamespaceIdentity)
	elemental.RegisterIdentity(IntegrationIdentity)
	elemental.RegisterIdentity(PolicyRuleIdentity)
	elemental.RegisterIdentity(ExternalServiceIdentity)
	elemental.RegisterIdentity(PolicyIdentity)
	elemental.RegisterIdentity(FlowStatisticIdentity)
	elemental.RegisterIdentity(FileAccessPolicyIdentity)
	elemental.RegisterIdentity(SystemCallIdentity)
	elemental.RegisterIdentity(EnforcerProfileIdentity)
	elemental.RegisterIdentity(ComputedDependencyMapViewIdentity)
	elemental.RegisterIdentity(RenderedPolicyIdentity)
	elemental.RegisterIdentity(NamespaceMappingPolicyIdentity)
	elemental.RegisterIdentity(ProcessingUnitIdentity)
	elemental.RegisterIdentity(DependencyMapViewIdentity)
	elemental.RegisterIdentity(DependencyMapIdentity)
	elemental.RegisterIdentity(VulnerabilityIdentity)
	elemental.RegisterIdentity(EnforcerProfileMappingPolicyIdentity)
	elemental.RegisterIdentity(ActivityIdentity)
	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(NetworkAccessPolicyIdentity)
}

// ModelVersion returns the current version of the model
func ModelVersion() float64 { return 1.0 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case APICheckIdentity.Name:
		return NewAPICheck()
	case MapNodeIdentity.Name:
		return NewMapNode()
	case DependencyMapSubviewIdentity.Name:
		return NewDependencyMapSubview()
	case APIAuthorizationPolicyIdentity.Name:
		return NewAPIAuthorizationPolicy()
	case SyscallAccessIdentity.Name:
		return NewSyscallAccess()
	case ComputedPolicyIdentity.Name:
		return NewComputedPolicy()
	case TagIdentity.Name:
		return NewTag()
	case EnforcerIdentity.Name:
		return NewEnforcer()
	case MapEdgeIdentity.Name:
		return NewMapEdge()
	case FilePathIdentity.Name:
		return NewFilePath()
	case FileAccessIdentity.Name:
		return NewFileAccess()
	case NamespaceIdentity.Name:
		return NewNamespace()
	case IntegrationIdentity.Name:
		return NewIntegration()
	case PolicyRuleIdentity.Name:
		return NewPolicyRule()
	case ExternalServiceIdentity.Name:
		return NewExternalService()
	case PolicyIdentity.Name:
		return NewPolicy()
	case FlowStatisticIdentity.Name:
		return NewFlowStatistic()
	case FileAccessPolicyIdentity.Name:
		return NewFileAccessPolicy()
	case SystemCallIdentity.Name:
		return NewSystemCall()
	case EnforcerProfileIdentity.Name:
		return NewEnforcerProfile()
	case ComputedDependencyMapViewIdentity.Name:
		return NewComputedDependencyMapView()
	case RenderedPolicyIdentity.Name:
		return NewRenderedPolicy()
	case NamespaceMappingPolicyIdentity.Name:
		return NewNamespaceMappingPolicy()
	case ProcessingUnitIdentity.Name:
		return NewProcessingUnit()
	case DependencyMapViewIdentity.Name:
		return NewDependencyMapView()
	case DependencyMapIdentity.Name:
		return NewDependencyMap()
	case VulnerabilityIdentity.Name:
		return NewVulnerability()
	case EnforcerProfileMappingPolicyIdentity.Name:
		return NewEnforcerProfileMappingPolicy()
	case ActivityIdentity.Name:
		return NewActivity()
	case RootIdentity.Name:
		return NewRoot()
	case NetworkAccessPolicyIdentity.Name:
		return NewNetworkAccessPolicy()
	default:
		return nil
	}
}

// IdentifiableForCategory returns a new instance of the Identifiable for the given category name.
func IdentifiableForCategory(category string) elemental.Identifiable {

	switch category {
	case APICheckIdentity.Category:
		return NewAPICheck()
	case MapNodeIdentity.Category:
		return NewMapNode()
	case DependencyMapSubviewIdentity.Category:
		return NewDependencyMapSubview()
	case APIAuthorizationPolicyIdentity.Category:
		return NewAPIAuthorizationPolicy()
	case SyscallAccessIdentity.Category:
		return NewSyscallAccess()
	case ComputedPolicyIdentity.Category:
		return NewComputedPolicy()
	case TagIdentity.Category:
		return NewTag()
	case EnforcerIdentity.Category:
		return NewEnforcer()
	case MapEdgeIdentity.Category:
		return NewMapEdge()
	case FilePathIdentity.Category:
		return NewFilePath()
	case FileAccessIdentity.Category:
		return NewFileAccess()
	case NamespaceIdentity.Category:
		return NewNamespace()
	case IntegrationIdentity.Category:
		return NewIntegration()
	case PolicyRuleIdentity.Category:
		return NewPolicyRule()
	case ExternalServiceIdentity.Category:
		return NewExternalService()
	case PolicyIdentity.Category:
		return NewPolicy()
	case FlowStatisticIdentity.Category:
		return NewFlowStatistic()
	case FileAccessPolicyIdentity.Category:
		return NewFileAccessPolicy()
	case SystemCallIdentity.Category:
		return NewSystemCall()
	case EnforcerProfileIdentity.Category:
		return NewEnforcerProfile()
	case ComputedDependencyMapViewIdentity.Category:
		return NewComputedDependencyMapView()
	case RenderedPolicyIdentity.Category:
		return NewRenderedPolicy()
	case NamespaceMappingPolicyIdentity.Category:
		return NewNamespaceMappingPolicy()
	case ProcessingUnitIdentity.Category:
		return NewProcessingUnit()
	case DependencyMapViewIdentity.Category:
		return NewDependencyMapView()
	case DependencyMapIdentity.Category:
		return NewDependencyMap()
	case VulnerabilityIdentity.Category:
		return NewVulnerability()
	case EnforcerProfileMappingPolicyIdentity.Category:
		return NewEnforcerProfileMappingPolicy()
	case ActivityIdentity.Category:
		return NewActivity()
	case RootIdentity.Category:
		return NewRoot()
	case NetworkAccessPolicyIdentity.Category:
		return NewNetworkAccessPolicy()
	default:
		return nil
	}
}

// ContentIdentifiableForIdentity returns a new instance of a ContentIdentifiable for the given identity name.
func ContentIdentifiableForIdentity(identity string) elemental.ContentIdentifiable {

	switch identity {
	case APICheckIdentity.Name:
		return &APIChecksList{}
	case MapNodeIdentity.Name:
		return &MapNodesList{}
	case DependencyMapSubviewIdentity.Name:
		return &DependencyMapSubviewsList{}
	case APIAuthorizationPolicyIdentity.Name:
		return &APIAuthorizationPoliciesList{}
	case SyscallAccessIdentity.Name:
		return &SyscallAccessList{}
	case ComputedPolicyIdentity.Name:
		return &ComputedPoliciesList{}
	case TagIdentity.Name:
		return &TagsList{}
	case EnforcerIdentity.Name:
		return &EnforcersList{}
	case MapEdgeIdentity.Name:
		return &MapEdgesList{}
	case FilePathIdentity.Name:
		return &FilePathsList{}
	case FileAccessIdentity.Name:
		return &FileAccessList{}
	case NamespaceIdentity.Name:
		return &NamespacesList{}
	case IntegrationIdentity.Name:
		return &IntegrationsList{}
	case PolicyRuleIdentity.Name:
		return &PolicyRulesList{}
	case ExternalServiceIdentity.Name:
		return &ExternalServicesList{}
	case PolicyIdentity.Name:
		return &PoliciesList{}
	case FlowStatisticIdentity.Name:
		return &FlowStatisticsList{}
	case FileAccessPolicyIdentity.Name:
		return &FileAccessPoliciesList{}
	case SystemCallIdentity.Name:
		return &SystemCallsList{}
	case EnforcerProfileIdentity.Name:
		return &EnforcerProfilesList{}
	case ComputedDependencyMapViewIdentity.Name:
		return &ComputedDependencyMapViewsList{}
	case RenderedPolicyIdentity.Name:
		return &RenderedPoliciesList{}
	case NamespaceMappingPolicyIdentity.Name:
		return &NamespaceMappingPoliciesList{}
	case ProcessingUnitIdentity.Name:
		return &ProcessingUnitsList{}
	case DependencyMapViewIdentity.Name:
		return &DependencyMapViewsList{}
	case DependencyMapIdentity.Name:
		return &DependencyMapsList{}
	case VulnerabilityIdentity.Name:
		return &VulnerabilitiesList{}
	case EnforcerProfileMappingPolicyIdentity.Name:
		return &EnforcerProfileMappingPoliciesList{}
	case ActivityIdentity.Name:
		return &ActivitiesList{}
	case NetworkAccessPolicyIdentity.Name:
		return &NetworkAccessPoliciesList{}
	default:
		return nil
	}
}

// ContentIdentifiableForCategory returns a new instance of a ContentIdentifiable for the given category name.
func ContentIdentifiableForCategory(category string) elemental.ContentIdentifiable {

	switch category {
	case APICheckIdentity.Category:
		return &APIChecksList{}
	case MapNodeIdentity.Category:
		return &MapNodesList{}
	case DependencyMapSubviewIdentity.Category:
		return &DependencyMapSubviewsList{}
	case APIAuthorizationPolicyIdentity.Category:
		return &APIAuthorizationPoliciesList{}
	case SyscallAccessIdentity.Category:
		return &SyscallAccessList{}
	case ComputedPolicyIdentity.Category:
		return &ComputedPoliciesList{}
	case TagIdentity.Category:
		return &TagsList{}
	case EnforcerIdentity.Category:
		return &EnforcersList{}
	case MapEdgeIdentity.Category:
		return &MapEdgesList{}
	case FilePathIdentity.Category:
		return &FilePathsList{}
	case FileAccessIdentity.Category:
		return &FileAccessList{}
	case NamespaceIdentity.Category:
		return &NamespacesList{}
	case IntegrationIdentity.Category:
		return &IntegrationsList{}
	case PolicyRuleIdentity.Category:
		return &PolicyRulesList{}
	case ExternalServiceIdentity.Category:
		return &ExternalServicesList{}
	case PolicyIdentity.Category:
		return &PoliciesList{}
	case FlowStatisticIdentity.Category:
		return &FlowStatisticsList{}
	case FileAccessPolicyIdentity.Category:
		return &FileAccessPoliciesList{}
	case SystemCallIdentity.Category:
		return &SystemCallsList{}
	case EnforcerProfileIdentity.Category:
		return &EnforcerProfilesList{}
	case ComputedDependencyMapViewIdentity.Category:
		return &ComputedDependencyMapViewsList{}
	case RenderedPolicyIdentity.Category:
		return &RenderedPoliciesList{}
	case NamespaceMappingPolicyIdentity.Category:
		return &NamespaceMappingPoliciesList{}
	case ProcessingUnitIdentity.Category:
		return &ProcessingUnitsList{}
	case DependencyMapViewIdentity.Category:
		return &DependencyMapViewsList{}
	case DependencyMapIdentity.Category:
		return &DependencyMapsList{}
	case VulnerabilityIdentity.Category:
		return &VulnerabilitiesList{}
	case EnforcerProfileMappingPolicyIdentity.Category:
		return &EnforcerProfileMappingPoliciesList{}
	case ActivityIdentity.Category:
		return &ActivitiesList{}
	case NetworkAccessPolicyIdentity.Category:
		return &NetworkAccessPoliciesList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		APICheckIdentity,
		MapNodeIdentity,
		DependencyMapSubviewIdentity,
		APIAuthorizationPolicyIdentity,
		SyscallAccessIdentity,
		ComputedPolicyIdentity,
		TagIdentity,
		EnforcerIdentity,
		MapEdgeIdentity,
		FilePathIdentity,
		FileAccessIdentity,
		NamespaceIdentity,
		IntegrationIdentity,
		PolicyRuleIdentity,
		ExternalServiceIdentity,
		PolicyIdentity,
		FlowStatisticIdentity,
		FileAccessPolicyIdentity,
		SystemCallIdentity,
		EnforcerProfileIdentity,
		ComputedDependencyMapViewIdentity,
		RenderedPolicyIdentity,
		NamespaceMappingPolicyIdentity,
		ProcessingUnitIdentity,
		DependencyMapViewIdentity,
		DependencyMapIdentity,
		VulnerabilityIdentity,
		EnforcerProfileMappingPolicyIdentity,
		ActivityIdentity,
		RootIdentity,
		NetworkAccessPolicyIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{
	"apiauth":    APIAuthorizationPolicyIdentity,
	"apiauths":   APIAuthorizationPolicyIdentity,
	"fps":        FilePathIdentity,
	"fp":         FilePathIdentity,
	"ns":         NamespaceIdentity,
	"extsrv":     ExternalServiceIdentity,
	"extsrvs":    ExternalServiceIdentity,
	"profiles":   EnforcerProfileIdentity,
	"profile":    EnforcerProfileIdentity,
	"rpols":      RenderedPolicyIdentity,
	"rpol":       RenderedPolicyIdentity,
	"nsmaps":     NamespaceMappingPolicyIdentity,
	"nsmap":      NamespaceMappingPolicyIdentity,
	"nspolicies": NamespaceMappingPolicyIdentity,
	"nspolicy":   NamespaceMappingPolicyIdentity,
	"pus":        ProcessingUnitIdentity,
	"pu":         ProcessingUnitIdentity,
	"vulns":      VulnerabilityIdentity,
	"vul":        VulnerabilityIdentity,
	"srvpol":     EnforcerProfileMappingPolicyIdentity,
	"srvpols":    EnforcerProfileMappingPolicyIdentity,
	"netpols":    NetworkAccessPolicyIdentity,
	"netpol":     NetworkAccessPolicyIdentity,
}

// IdentityFromAlias returns the Identity associated to the given alias.
func IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case APICheckIdentity:
		return []string{}
	case MapNodeIdentity:
		return []string{}
	case DependencyMapSubviewIdentity:
		return []string{}
	case APIAuthorizationPolicyIdentity:
		return []string{
			"apiauth",
			"apiauths",
		}
	case SyscallAccessIdentity:
		return []string{}
	case ComputedPolicyIdentity:
		return []string{}
	case TagIdentity:
		return []string{}
	case EnforcerIdentity:
		return []string{}
	case MapEdgeIdentity:
		return []string{}
	case FilePathIdentity:
		return []string{
			"fps",
			"fp",
		}
	case FileAccessIdentity:
		return []string{}
	case NamespaceIdentity:
		return []string{
			"ns",
		}
	case IntegrationIdentity:
		return []string{}
	case PolicyRuleIdentity:
		return []string{}
	case ExternalServiceIdentity:
		return []string{
			"extsrv",
			"extsrvs",
		}
	case PolicyIdentity:
		return []string{}
	case FlowStatisticIdentity:
		return []string{}
	case FileAccessPolicyIdentity:
		return []string{}
	case SystemCallIdentity:
		return []string{}
	case EnforcerProfileIdentity:
		return []string{
			"profiles",
			"profile",
		}
	case ComputedDependencyMapViewIdentity:
		return []string{}
	case RenderedPolicyIdentity:
		return []string{
			"rpols",
			"rpol",
		}
	case NamespaceMappingPolicyIdentity:
		return []string{
			"nsmaps",
			"nsmap",
			"nspolicies",
			"nspolicy",
		}
	case ProcessingUnitIdentity:
		return []string{
			"pus",
			"pu",
		}
	case DependencyMapViewIdentity:
		return []string{}
	case DependencyMapIdentity:
		return []string{}
	case VulnerabilityIdentity:
		return []string{
			"vulns",
			"vul",
		}
	case EnforcerProfileMappingPolicyIdentity:
		return []string{
			"srvpol",
			"srvpols",
		}
	case ActivityIdentity:
		return []string{}
	case RootIdentity:
		return []string{}
	case NetworkAccessPolicyIdentity:
		return []string{
			"netpols",
			"netpol",
		}
	}

	return nil
}
