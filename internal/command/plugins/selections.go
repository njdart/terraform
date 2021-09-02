package plugins

import (
	"github.com/hashicorp/terraform/internal/addrs"
	"github.com/hashicorp/terraform/internal/providers"
	"github.com/hashicorp/terraform/internal/provisioners"
)

// Selections represents a specific set of plugins chosen by a Finder, based
// on its settings for where to look and other constraints that must be
// applied.
type Selections struct {
}

// FindPlugins uses the settings in the receiver to select a specific set of
// provider plugins that will either meet the configured constraints or
// generate errors explaining why they don't.
func (f Finder) FindPlugins() (*Selections, error) {
	// TODO: Implement
	return &Selections{}, nil
}

// ProviderFactories builds a map of provider factory functions based on the
// recieving sections.
//
// Terraform Core (terraform.NewContext, specifically) interacts with
// providers by calling a factory function each time it needs a new provider
// instance, and so the result of this method serves as the interface between
// Terraform Core and a plugin finder for the purpose of launching providers
// in particular. (Other plugin types have similar Finder methods, which
// callers must use separately.)
//
// This function always succeeds itself, because verification of the available
// providers is deferred until Terraform Core attempts to instantiate one.
// Some or all of the factory functions might therefore fail with an error
// in the event that the Finder is misconfigured or that the working directory
// is in an inconsistent state.
//
// The result only contains factories for providers included in the requirements
// set previously passed to WithProviderRequirements.
func (s *Selections) ProviderFactories() map[addrs.Provider]providers.Factory {
	ret := make(map[addrs.Provider]providers.Factory)
	// TODO: Implement
	return ret
}

// ProvisionerFactories builds a map of provisioner factory functions based on
// the recieving selections.
//
// Terraform Core (terraform.NewContext, specifically) interacts with
// provisioners by calling a factory function each time it needs a new
// provisioner instance, and so the result of this method serves as the
// interface between Terraform Core and a plugin finder for the purpose of
// launching provisioners in particular. (Other plugin types have similar
// Finder methods, which callers must use separately.)
//
// This function always succeeds itself, because verification of the available
// provisioners is deferred until Terraform Core attempts to instantiate one.
// Some or all of the factory functions might therefore fail with an error
// in the event that the Finder is misconfigured or that the working directory
// is in an inconsistent state.
func (s *Selections) ProvisionerFactories() map[string]provisioners.Factory {
	ret := make(map[string]provisioners.Factory)
	// TODO: Implement
	return ret
}
