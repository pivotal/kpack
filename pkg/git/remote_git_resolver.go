package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/storage/memory"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
)

const defaultRemote = "origin"

type remoteGitResolver struct {
}

func (*remoteGitResolver) Resolve(auth transport.AuthMethod, sourceConfig v1alpha1.SourceConfig) (v1alpha1.ResolvedSourceConfig, error) {
	repo := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: defaultRemote,
		URLs: []string{sourceConfig.Git.URL},
	})
	references, err := repo.List(&git.ListOptions{
		Auth: auth,
	})
	if err != nil {
		return v1alpha1.ResolvedSourceConfig{
			Git: &v1alpha1.ResolvedGitSource{
				URL:      sourceConfig.Git.URL,
				Revision: sourceConfig.Git.Revision, // maybe
				Type:     v1alpha1.Unknown,
				SubPath:  sourceConfig.SubPath,
			},
		}, nil
	}

	for _, ref := range references {
		if ref.Name().Short() == sourceConfig.Git.Revision {
			return v1alpha1.ResolvedSourceConfig{
				Git: &v1alpha1.ResolvedGitSource{
					URL:      sourceConfig.Git.URL,
					Revision: ref.Hash().String(),
					Type:     sourceType(ref),
					SubPath:  sourceConfig.SubPath,
				},
			}, nil
		}
	}

	return v1alpha1.ResolvedSourceConfig{
		Git: &v1alpha1.ResolvedGitSource{
			URL:      sourceConfig.Git.URL,
			Revision: sourceConfig.Git.Revision,
			Type:     v1alpha1.Commit,
			SubPath:  sourceConfig.SubPath,
		},
	}, nil
}

func sourceType(reference *plumbing.Reference) v1alpha1.GitSourceKind {
	switch {
	case reference.Name().IsBranch():
		return v1alpha1.Branch
	case reference.Name().IsTag():
		return v1alpha1.Tag
	default:
		return v1alpha1.Unknown
	}
}
