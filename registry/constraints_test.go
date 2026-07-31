package registry

import (
	"testing"

	. "github.com/xregistry/server/common"
)

// TestConstraintsEqualsWildcardModelError verifies that a constraint's
// "equals" can't reference "*" (the wildcard attribute name) since it's
// not a real, named attribute. This is a pure model-verification error -
// no DB/Registry needed, so it's tested directly against a parsed Model
// rather than through a live *registry.Registry.
func TestConstraintsEqualsWildcardModelError(t *testing.T) {
	// Model error: equals references "*" which is a wildcard, not a named attr
	modelSrc := `{
	  "groups": { "dirs": {
	    "singular": "dir",
	    "attributes": {
	      "*": { "type": "string" }
	    },
	    "constraints": {
	      "files.mystr": { "equals": "someattr" }
	    },
	    "resources": {"files": {"singular": "file", "hasdocument": false,
	      "attributes": { "mystr": { "type": "string" } } } } } } }`

	m, xErr := ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)

	XCheckErr(t, m.Verify(),
		`^(?s)^.*model_error.*equals.*someattr.*can not be found`)
}
