/// go-bindata -o=etc/asset.go -pkg=etc etc/config.yaml
package etc

import "fmt"

func Asset(name string) ([]byte, error) {
	return nil, fmt.Errorf("Asset %s not found", name)
}
