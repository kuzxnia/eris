package resource

import "errors"

// type ResourceSelectorRegistry interface {
// ResolveResourceSelector(selector *Selector) ([]*Resource, error) 
// }

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

func (r *ResourceSelectorRegistry) ResolveResourceSelector(selector *Selector) ([]*Resource, error) {
	if resourceSelector, ok := r.resourceSelectors[selector.Type]; ok {
    return resourceSelector.Select(selector)
	} else {
		return nil, errors.New("Resource selector not found for type " + selector.Type)
	}
}
