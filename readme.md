# lstags

experimental lister of tags in grep-like format

```
internal/sidecar/sidecar_gen.go:17:1:NewSidecarBuilder() *SidecarBuilder
internal/sidecar/sidecar_gen.go:20:1:(SidecarBuilder) Cloud(cloud *cloud.Cloud) *SidecarBuilder
internal/sidecar/sidecar_gen.go:24:1:(SidecarBuilder) Name(name string) *SidecarBuilder
internal/sidecar/sidecar_gen.go:28:1:(SidecarBuilder) PipelinesDir(pipelinesDir string) *SidecarBuilder
internal/sidecar/sidecar_gen.go:32:1:(SidecarBuilder) Slug(slug string) *SidecarBuilder
internal/sidecar/sidecar_gen.go:36:1:(SidecarBuilder) CommandConsumer(commandConsumer cloud.CommandConsumer) *SidecarBuilder
internal/sidecar/sidecar_gen.go:40:1:(SidecarBuilder) OutputConsumer(outputConsumer cloud.OutputConsumer) *SidecarBuilder
internal/sidecar/sidecar_gen.go:44:1:(SidecarBuilder) SshKey(sshKey string) *SidecarBuilder
internal/sidecar/sidecar_gen.go:48:1:(SidecarBuilder) Build() *Sidecar
internal/sshkey/sshkey.go:14:1:GeneratePair(path string, size int) error
internal/sshkey/sshkey.go:52:1:generatePrivateKey(bitSize int) (*rsa.PrivateKey, error)
internal/sshkey/sshkey.go:66:1:marshalPrivateKeyToPEM(private *rsa.PrivateKey) []byte
internal/sshkey/sshkey.go:78:1:generatePublicKey(private *rsa.PublicKey) ([]byte, error)
```

# install

```
go get -v -u github.com/reconquest/lstags/cmd/lstags
```
