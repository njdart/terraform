package plugins

import (
	"github.com/hashicorp/terraform/internal/addrs"
	"github.com/hashicorp/terraform/internal/depsfile"
	"github.com/hashicorp/terraform/internal/getproviders"
	"github.com/hashicorp/terraform/internal/providers"
)

// WithProviderRequirements returns a Finder which has the same settings
// as the reciever except that its set of provider requirements (if any)
// is replaced with the given set.
//
// Finder guarantees that only providers included in the requirements set
// will be successfully launchable, with those outside of the set always
// returning an error on launch even if they happen to be available on disk
// to run.
//
// An exception is the set of providers included in either the dev overrides
// or unmanaged providers sets. In either case, Finder will still only allow
// a provider if it's included in the given requirements, but it will skip
// checking version constraints because a provider under development doesn't
// yet have a version number to check against.
//
// The initial result of NewFinder has an empty set of provider requirements,
// so some component typically must call WithProviderRequirements prior to
// creating provider factories, or else there will be no providers available
// for use.
func (l Finder) WithProviderRequirements(reqs getproviders.Requirements) Finder {
	l.providerRequirements = reqs
	return l
}

// WithDependencyLocks returns a Finder which has the same settings as the
// reciever except that its set of dependency locks (if any) is replaced with
// the given set.
//
// Finder guarantees that only providers included in the locks will be
// successfully launchable, except providers included in either the set of
// development overrides or the set of unmanaged providers.
//
// The initial result of NewFinder has an empty set of dependency locks,
// so some component typically must call WithDependencyLocks prior to creating
// provider factories, or else there will be no valid providers to use.
//
// We currently use dependency locks only for provider plugins, but the scope
// of the dependency lock mechanism might grow in future.
func (l Finder) WithDependencyLocks(locks *depsfile.Locks) Finder {
	l.dependencyLocks = locks
	return l
}

// WithForcedProviderChecksums returns a Finder which has the same settings
// as the reciever except for forcing exact checksums for all provider plugins.
//
// The initial result of NewFinder doesn't impose this additional requirement,
// and a caller would typically impose this requirement only when applying
// a saved plan file, in order to force using only exactly the same plugin
// executables as what generated the plan.
//
// Calling this method with an empty map means that no providers are allowed
// at all. To un-constrain forced checksums completely, call
// WithoutForcedProviderChecksums. The result of NewFinder is as if you had
// already called WithoutForcedProviderChecksums.
func (l Finder) WithForcedProviderChecksums(checksums map[addrs.Provider]getproviders.Hash) Finder {
	if checksums == nil {
		// we need a non-nil map so we can recognize the difference between
		// WithForcedProviderSHA256s on an empty map vs.
		// WithoutForcedProviderSHA256s.
		checksums = make(map[addrs.Provider]getproviders.Hash)
	}
	l.providerForceChecksums = checksums
	return l
}

// WithoutForcedProviderChecksums returns a Finder which has the same settings
// as the receiver except for removing the effect of a previous call to
// WithForcedProviderChecksums.
//
// Having no forced checksums is the default for the result from NewFinder,
// so this method should not be necessary in most cases.
func (l Finder) WithoutForcedProviderChecksums() Finder {
	l.providerForceChecksums = nil
	return l
}

// ProviderFactories builds a map of provider factory functions based on the
// settings of the receiving Finder.
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
func (l Finder) ProviderFactories() map[addrs.Provider]providers.Factory {
	ret := make(map[addrs.Provider]providers.Factory, len(l.providerRequirements))
	// TODO: Implement
	return ret
}
