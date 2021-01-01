package cks

import "github.com/mig4/gogiggle/cli-experiments/cobra-k8s-selectors/pkg/swapi"

type Port interface {
	List(options *ListOptions) ([]swapi.Character, error)
}
