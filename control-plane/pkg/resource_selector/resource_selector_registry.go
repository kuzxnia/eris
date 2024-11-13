package resourceselector

import "errors"

type ResourceSelectorRegistry struct {
	resourceSelectors map[string]ResourceSelector
}

func ProvideResourceSelectorRegistry(resourceSelectors []ResourceSelector) *ResourceSelectorRegistry {
	resourceSelectorByType := map[string]ResourceSelector{}

	for _, resourceSelector := range resourceSelectors {
		if _, isPresent := resourceSelectorByType[resourceSelector.GetSelectorType()]; isPresent {
			panic("resource selector " + resourceSelector.GetSelectorType() + " already registered. Resource selector types should be unique.")
		}
		resourceSelectorByType[resourceSelector.GetSelectorType()] = resourceSelector
	}

	return &ResourceSelectorRegistry{
		resourceSelectors: resourceSelectorByType,
	}
}

func (r *ResourceSelectorRegistry) ResolveResourceSelectorBySelector(selector *Selector) (ResourceSelector, error) {
	if resourceSelector, ok := r.resourceSelectors[selector.Type]; ok {
		return resourceSelector, nil
	} else {
		return nil, errors.New("Resource selector not found for type " + selector.Type)
	}
}