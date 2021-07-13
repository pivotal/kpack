package cnb

import (
	"sort"

	"github.com/Masterminds/semver/v3"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/pkg/errors"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
)

type StoreBuildpackRepository struct {
	Keychain authn.Keychain

	ClusterStore *buildapi.ClusterStore
}

func (s *StoreBuildpackRepository) FindByIdAndVersion(id, version string) (RemoteBuildpackInfo, error) {
	storeBuildpack, err := s.findBuildpack(id, version)
	if err != nil {
		return RemoteBuildpackInfo{}, err
	}

	layer, err := layerFromStoreBuildpack(s.Keychain, storeBuildpack)
	if err != nil {
		return RemoteBuildpackInfo{}, err
	}

	layers, err := s.layersForOrder(storeBuildpack.Order)
	if err != nil {
		return RemoteBuildpackInfo{}, err
	}

	info := DescriptiveBuildpackInfo{
		BuildpackInfo: buildapi.BuildpackInfo{
			Id:      storeBuildpack.Id,
			Version: storeBuildpack.Version,
		},
		Homepage: storeBuildpack.Homepage,
	}

	return RemoteBuildpackInfo{
		BuildpackInfo: info,
		Layers: append(layers, buildpackLayer{
			v1Layer:       layer,
			BuildpackInfo: info,
			BuildpackLayerInfo: BuildpackLayerInfo{
				LayerDiffID: storeBuildpack.DiffId,
				Order:       storeBuildpack.Order,
				API:         storeBuildpack.API,
				Stacks:      storeBuildpack.Stacks,
				Homepage:    storeBuildpack.Homepage,
			},
		}),
	}, nil
}

func (s *StoreBuildpackRepository) findBuildpack(id, version string) (buildapi.StoreBuildpack, error) {
	var matchingBuildpacks []buildapi.StoreBuildpack
	for _, buildpack := range s.ClusterStore.Status.Buildpacks {
		if buildpack.Id == id {
			matchingBuildpacks = append(matchingBuildpacks, buildpack)
		}
	}

	if len(matchingBuildpacks) == 0 {
		return buildapi.StoreBuildpack{}, errors.Errorf("could not find buildpack with id '%s'", id)
	}

	if version == "" {
		return highestVersion(matchingBuildpacks)
	}

	for _, buildpack := range matchingBuildpacks {
		if buildpack.Version == version {
			return buildpack, nil
		}
	}

	return buildapi.StoreBuildpack{}, errors.Errorf("could not find buildpack with id '%s' and version '%s'", id, version)
}

// TODO: ensure there are no cycles in the buildpack graph
func (s *StoreBuildpackRepository) layersForOrder(order buildapi.Order) ([]buildpackLayer, error) {
	var buildpackLayers []buildpackLayer
	for _, orderEntry := range order {
		for _, buildpackRef := range orderEntry.Group {
			buildpackInfo, err := s.FindByIdAndVersion(buildpackRef.Id, buildpackRef.Version)
			if err != nil {
				return nil, err
			}

			buildpackLayers = append(buildpackLayers, buildpackInfo.Layers...)
		}

	}
	return buildpackLayers, nil
}

func highestVersion(matchingBuildpacks []buildapi.StoreBuildpack) (buildapi.StoreBuildpack, error) {
	for _, bp := range matchingBuildpacks {
		if _, err := semver.NewVersion(bp.Version); err != nil {
			return buildapi.StoreBuildpack{}, errors.Errorf("cannot find buildpack '%s' with latest version due to invalid semver '%s'", bp.Id, bp.Version)
		}
	}
	sort.Sort(byBuildpackVersion(matchingBuildpacks))
	return matchingBuildpacks[len(matchingBuildpacks)-1], nil
}

type byBuildpackVersion []buildapi.StoreBuildpack

func (b byBuildpackVersion) Len() int {
	return len(b)
}

func (b byBuildpackVersion) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byBuildpackVersion) Less(i, j int) bool {
	return semver.MustParse(b[i].Version).LessThan(semver.MustParse(b[j].Version))
}
