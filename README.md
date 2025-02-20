# kbgen

**kbgen** is a tool that generates Go structs for Kubebuilder from KCL schemas. This approach ensures that schemas are compatible and ready for use—even when dealing with complex schema structures. By transforming KCL schema definitions into Go structs, kbgen simplifies integration with Kubebuilder-based (Kubernetes Operator) projects.

## Installation

### 1. Download and Install Pre-built Binary

Install kbgen using our dynamic installer. This script automatically detects the latest version and downloads the correct binary for your operating system and architecture. Make sure you have both curl and jq installed.

Run the following commands in your terminal:

```bash
export KBGEN_VERSION=$(curl -s "https://api.github.com/repos/eminaktas/kbgen/tags" | jq -r '.[0].name')
curl -sSL https://raw.githubusercontent.com/eminaktas/kbgen/$KBGEN_VERSION/install.sh | sh
```

### 2. Use `go install`

If you have a Go environment set up, you can install **kbgen** with:

```bash
go install github.com/eminaktas/kbgen/cmd/kbgen@latest
```

This will compile and install kbgen into your `$GOPATH/bin`.

### 3. Build it from Source Code

Clone the repository and build the binary:

```bash
git clone https://github.com/eminaktas/kbgen.git
cd kbgen
go build -o kbgen ./cmd/kbgen
```

Then, move the binary to your `$PATH`:

```bash
sudo mv kbgen /usr/local/bin/
chmod +x /usr/local/bin/kbgen
```

## Usage

After installation, you can generate Go structs by running the following command:

```bash
kbgen gen --outputDir=example/example-go/pkg --packageName=models --programPath=example/kubeconf --directory=example/kubeconf/models --configPath=example/conf.yaml
```

### Command Flags

- --outputDir: Output directory for the generated Go structs.
- --packageName: Package name for the generated Go code.
- --programPath: The directory where the KCL module (mod file) exists.
- --directory: Directory containing the .k files used to generate the schemas.
- --configPath: Path to the YAML configuration file.

### Configuration

The generator uses a YAML configuration file to fine-tune its behavior. Below is an example configuration (example/conf.yaml):

```yaml
# List of schema names for which a DeepCopyInto function should be generated.
deepCopyNominee:
  - Server

# Optional configuration for handling multiple types in KCL.
# Uncomment and adjust if you need to override the default type (*apiextensionsv1.JSON).
# customAnyType:
#   type: '*apiextensionsv1.JSON'
#   import: 'apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"'

# mapTypeSchemas:
# Workaround for index signature schemas.
# Due to a known issue in kcl-go (see: https://github.com/kcl-lang/kcl-go/issues/411),
# schemas using index signatures might not be detected correctly.
# Use this configuration to explicitly define them.
mapTypeSchemas:
  - name: EnvMap
    pkgPath: "konfig.models.frontend.container.env"
    item:
      schemaName: Env
      # Optional: Documentation for the schema.
      schemaDoc: "Env represents an environment variable present in a Container."
```

### Workaround for Index Signature Schemas

Due to an issue in the underlying kcl-go library (see [Issue #411](https://github.com/kcl-lang/kcl-go/issues/411)), schemas that use index signatures may not be automatically detected.
Workaround:
Use the mapTypeSchemas configuration in your YAML file to explicitly define these schemas. For example, if your schema uses an index signature and is named EnvMap, add an entry under mapTypeSchemas as shown in the configuration above.

### Kubebuilder Example

Below is an example of how the generated structs can be integrated into a Kubebuilder project:

```go
package v1alpha1

import "github.com/eminaktas/kbgen/example/example-go/pkg/models"

// KubeConfSpec defines the desired state of KubeConf.
type KubeConfSpec struct {
   Server *models.Server `json:"server,omitempty"`
}
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with your suggestions or bug fixes.

## License

This project is licensed under the [Apache-2.0 License](LICENSE).
