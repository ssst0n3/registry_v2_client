package test_data

import (
	"github.com/containerd/containerd/images"
	"github.com/docker/distribution"
	"github.com/docker/distribution/manifest"
	"github.com/docker/distribution/manifest/schema2"
)

var ManifestHelloWorld = schema2.Manifest{
	Versioned: manifest.Versioned{
		SchemaVersion: 2,
		MediaType:     images.MediaTypeDockerSchema2Manifest,
	},
	Config: distribution.Descriptor{
		MediaType: images.MediaTypeDockerSchema2Config,
		Size:      1510,
		Digest:    "sha256:bf756fb1ae65adf866bd8c456593cd24beb6a0a061dedf42b26a993176745f6b",
	},
	Layers: []distribution.Descriptor{
		{
			MediaType: images.MediaTypeDockerSchema2LayerGzip,
			Size:      2529,
			Digest:    "sha256:0e03bdcc26d7a9a57ef3b6f1bf1a210cff6239bff7c8cac72435984032851689",
		},
	},
}
