package test_data

import (
	"github.com/containerd/containerd/images"
	"github.com/docker/distribution"
	"github.com/docker/distribution/manifest"
	"github.com/docker/distribution/manifest/schema2"
)

var (
	ManifestHelloWorld = schema2.Manifest{
		Versioned: manifest.Versioned{
			SchemaVersion: 2,
			MediaType:     images.MediaTypeDockerSchema2Manifest,
		},
		Config: distribution.Descriptor{
			MediaType: images.MediaTypeDockerSchema2Config,
			Size:      1729,
			Digest:    "sha256:a1cec1bba3abba8c203e81c8c1d315e932e26bee578725a3a376704ce4203f90",
		},
		Layers: []distribution.Descriptor{
			{
				MediaType: images.MediaTypeDockerSchema2LayerGzip,
				Size:      709433,
				Digest:    "sha256:9f5a74d688881b9b1c2e97b70e9480f89d26e62f221c4130f5a9f28c1a96bf99",
			},
		},
	}
	ManifestHelloWorldDigest = "sha256:c7c17e128826cacab26313fa54f3137c5a692611fcbf2e4279ca2226f8f406b9"
)

//func init() {
//	bytes, err := json.Marshal(ManifestHelloWorld)
//	awesome_error.CheckFatal(err)
//	ManifestHelloWorldDigest = digest.FromBytes(bytes)
//}
