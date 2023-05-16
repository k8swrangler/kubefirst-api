/*
Copyright (C) 2021-2023, Kubefirst

This program is licensed under MIT.
See the LICENSE file for more details.
*/
package gitShim

import (
	"context"
	"fmt"
	"io"

	"github.com/google/go-github/v52/github"
)

const (
	KubefirstGitHubOrganization    = "kubefirst"
	KubefirstMarketplaceRepository = "marketplace"
	basePath                       = "/"
)

// GetMarketplaceRepo returns an object detailing the Kubefirst marketplace GitHub repository
func (gh *GitHubClient) GetMarketplaceRepo() (*github.Repository, error) {
	repo, _, err := gh.Client.Repositories.Get(
		context.Background(),
		KubefirstGitHubOrganization,
		KubefirstMarketplaceRepository,
	)
	if err != nil {
		return &github.Repository{}, err
	}

	return repo, nil
}

// ReadMarketplaceRepoContents reads the file and directory contents of the Kubefirst marketplace
// GitHub repository
func (gh *GitHubClient) ReadMarketplaceRepoContents() ([]*github.RepositoryContent, error) {
	_, directoryContent, _, err := gh.Client.Repositories.GetContents(
		context.Background(),
		KubefirstGitHubOrganization,
		KubefirstMarketplaceRepository,
		basePath,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return directoryContent, nil
}

// ParseMarketplaceIndex reads the marketplace repository index
func (gh *GitHubClient) ParseMarketplaceIndex(contents []*github.RepositoryContent) ([]byte, error) {
	for _, content := range contents {
		switch *content.Type {
		case "file":
			switch *content.Name {
			case "index.yaml":
				b, err := gh.readFileContents(content)
				if err != nil {
					return b, err
				}
				return b, nil
			}
		}
	}

	return []byte{}, fmt.Errorf("index.yaml not found in marketplace repository")
}

// readFileContents parses the contents of a file in a GitHub repository
func (gh *GitHubClient) readFileContents(content *github.RepositoryContent) ([]byte, error) {
	rc, _, err := gh.Client.Repositories.DownloadContents(
		context.Background(),
		KubefirstGitHubOrganization,
		KubefirstMarketplaceRepository,
		*content.Path,
		nil,
	)
	if err != nil {
		return []byte{}, err
	}
	defer rc.Close()

	b, err := io.ReadAll(rc)
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}
