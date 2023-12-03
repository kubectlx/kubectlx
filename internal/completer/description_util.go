package completer

import "fmt"

func typeDescription(format, typeName string) string {
	return fmt.Sprintf(format, typeName)
}
