package yaf

import (
	"io"

	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func Filter(dst io.Writer, src io.Reader, res map[string]map[string]struct{}, exclude bool) error {
	r := &kio.ByteReader{
		OmitReaderAnnotations: true,
		AnchorsAweigh:         true,
		Reader:                src,
	}

	nodes, err := r.Read()
	if err != nil {
		return err
	}

	filtered := make([]*yaml.RNode, 0, len(nodes))
	for _, node := range nodes {
		_, matched := res[node.GetApiVersion()][node.GetKind()]
		if (matched && exclude) || (!matched && !exclude) {
			continue
		}
		filtered = append(filtered, node)
	}

	return (&kio.ByteWriter{Writer: dst}).Write(filtered)
}
