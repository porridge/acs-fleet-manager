package gitops

import (
	"encoding/json"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/operator/apis/platform/v1alpha1"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"sigs.k8s.io/yaml"
)

// Service applies GitOps configuration to Central instances.
type Service interface {
	GetCentral(ctx CentralParams) (v1alpha1.Central, error)
}

type service struct {
	configProvider ConfigProvider
}

// NewService returns a new Service.
func NewService(configProvider ConfigProvider) Service {
	return &service{configProvider: configProvider}
}

// GetCentral returns a Central instance with the given parameters.
func (s *service) GetCentral(params CentralParams) (v1alpha1.Central, error) {
	wr := new(strings.Builder)
	if err := defaultTemplate.Execute(wr, params); err != nil {
		return v1alpha1.Central{}, errors.Wrap(err, "failed to render default template")
	}
	var central v1alpha1.Central
	if err := yaml.Unmarshal([]byte(wr.String()), &central); err != nil {
		return v1alpha1.Central{}, errors.Wrap(err, "failed to unmarshal default central")
	}
	cfg, err := s.configProvider.Get()
	if err != nil {
		return v1alpha1.Central{}, errors.Wrap(err, "failed to get GitOps configuration")
	}
	return applyConfigToCentral(cfg, central, params)
}

// CentralParams represents the parameters for a Central instance.
type CentralParams struct {
	// ID is the ID of the Central instance.
	ID string
	// Name is the name of the Central instance.
	Name string
	// Namespace is the namespace of the Central instance.
	Namespace string
	// Region is the region of the Central instance.
	Region string
	// ClusterID is the ID of the cluster of the Central instance.
	ClusterID string
	// CloudProvider is the cloud provider of the Central instance.
	CloudProvider string
	// CloudAccountID	is the cloud account ID of the Central instance.
	CloudAccountID string
	// SubscriptionID is the subscription ID of the Central instance.
	SubscriptionID string
	// Owner is the owner of the Central instance.
	Owner string
	// OwnerAccountID is the owner account ID of the Central instance.
	OwnerAccountID string
	// OwnerUserID is the owner user ID of the Central instance.
	OwnerUserID string
	// Host is the host of the Central instance.
	Host string
	// OrganizationID is the organization ID of the Central instance.
	OrganizationID string
	// OrganizationName is the organization name of the Central instance.
	OrganizationName string
	// InstanceType is the instance type of the Central instance.
	InstanceType string
	// IsInternal is true if the Central instance is internal.
	IsInternal bool
}

// applyConfigToCentral will apply the given GitOps configuration to the given Central instance.
func applyConfigToCentral(config Config, central v1alpha1.Central, ctx CentralParams) (v1alpha1.Central, error) {
	var overrides []CentralOverride
	for _, override := range config.Centrals.Overrides {
		if !shouldApplyCentralOverride(override, ctx) {
			continue
		}
		overrides = append(overrides, override)
	}
	if len(overrides) == 0 {
		return central, nil
	}
	// render override path templates
	for i, override := range overrides {
		var err error
		overrides[i].Patch, err = renderPatchTemplate(override.Patch, ctx)
		if err != nil {
			return v1alpha1.Central{}, err
		}
	}
	centralBytes, err := json.Marshal(central)
	if err != nil {
		return v1alpha1.Central{}, errors.Wrap(err, "failed to marshal Central instance")
	}
	for _, override := range overrides {
		patchBytes := []byte(override.Patch)
		centralBytes, err = applyPatchToCentral(centralBytes, patchBytes)
		if err != nil {
			return v1alpha1.Central{}, err
		}
	}
	var result v1alpha1.Central
	if err := json.Unmarshal(centralBytes, &result); err != nil {
		return v1alpha1.Central{}, errors.Wrap(err, "failed to unmarshal Central instance")
	}
	return result, nil
}

// shouldApplyCentralOverride returns true if the given Central override should be applied to the given Central instance.
func shouldApplyCentralOverride(override CentralOverride, ctx CentralParams) bool {
	for _, d := range override.InstanceIDs {
		if d == "*" {
			return true
		}
		if d == ctx.ID {
			return true
		}
	}
	return false
}

// applyPatchToCentral will apply the given patch to the given Central instance.
func applyPatchToCentral(centralBytes, patch []byte) ([]byte, error) {
	// convert patch from yaml to json
	jsonPath, err := yaml.YAMLToJSON(patch)
	if err != nil {
		return []byte{}, errors.Wrap(err, "failed to convert override patch from yaml to json")
	}
	// apply patch
	patchedBytes, err := strategicpatch.StrategicMergePatch(centralBytes, jsonPath, v1alpha1.Central{})
	if err != nil {
		return []byte{}, errors.Wrap(err, "failed to apply override to Central instance")
	}
	return patchedBytes, nil
}

func renderPatchTemplate(patchTemplate string, ctx CentralParams) (string, error) {
	tpl, err := template.New("patch").Parse(patchTemplate)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse patch template")
	}
	var writer = new(strings.Builder)
	if err := tpl.Execute(writer, ctx); err != nil {
		return "", errors.Wrap(err, "failed to render patch template")
	}
	return writer.String(), nil
}

// defaultTemplate is the default template for Central instances.
var defaultTemplate = template.Must(template.New("default").Parse(string(defaultCentralTemplate)))
