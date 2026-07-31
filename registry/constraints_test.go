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

// TestConstraintsMatchVersionsModelErrors verifies the "matchversions"
// model-definition rules: only allowed on scalar attributes, and never
// on attributes nested under an array, a map, an "ifvalues" clause, or a
// "*" wildcard extension. All pure model-verification errors - no
// DB/Registry needed.
func TestConstraintsMatchVersionsModelErrors(t *testing.T) {
	type test struct {
		name string
		src  string
		err  string
	}

	tests := []test{
		{"scalar - ok", `{
      "groups": { "dirs": { "singular": "dir", "resources": { "files": {
              "singular": "file",
              "attributes": {
                "mystr": {
                  "type": "string",
                  "matchversions": true
                } } } } } } }`, ""},

		{"object - bad", `{
      "groups": { "dirs": { "singular": "dir", "resources": { "files": {
              "singular": "file",
              "attributes": {
                "myobj": {
                  "type": "object",
                  "matchversions": true
                } } } } } } }`, `{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: \"groups.dirs.resources.files.myobj\" is not allowed to have \"matchversions\" set to \"true\" due to it not being a scalar attribute.",
  "subject": "/model",
  "args": {
    "error_detail": "\"groups.dirs.resources.files.myobj\" is not allowed to have \"matchversions\" set to \"true\" due to it not being a scalar attribute"
  },
  "source": "xxx"
}`},

		{"map->obj - bad", `{
      "groups": { "dirs": { "singular": "dir", "resources": { "files": {
              "singular": "file",
              "attributes": {
                "mymap": {
                  "type": "map",
                  "item": { "type": "object" },
                  "matchversions": true
                } } } } } } }`, `{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: \"groups.dirs.resources.files.mymap\" is not allowed to have \"matchversions\" set to \"true\" due to it not being a scalar attribute.",
  "subject": "/model",
  "args": {
    "error_detail": "\"groups.dirs.resources.files.mymap\" is not allowed to have \"matchversions\" set to \"true\" due to it not being a scalar attribute"
  },
  "source": "xxx"
}`},

		{"map->int - bad", `{
      "groups": { "dirs": { "singular": "dir", "resources": { "files": {
              "singular": "file",
              "attributes": {
                "mymap": {
                  "type": "map",
                  "item": { "type": "integer" },
                  "matchversions": true
                } } } } } } }`, `{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: \"groups.dirs.resources.files.mymap\" is not allowed to have \"matchversions\" set to \"true\" due to it not being a scalar attribute.",
  "subject": "/model",
  "args": {
    "error_detail": "\"groups.dirs.resources.files.mymap\" is not allowed to have \"matchversions\" set to \"true\" due to it not being a scalar attribute"
  },
  "source": "xxx"
}`},

		{"array - bad", `{
      "groups": { "dirs": { "singular": "dir", "resources": { "files": {
              "singular": "file",
              "attributes": {
                "myarray": {
                  "type": "array",
                  "item": { "type": "integer" },
                  "matchversions": true
                } } } } } } }`, `{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: \"groups.dirs.resources.files.myarray\" is not allowed to have \"matchversions\" set to \"true\" due to it not being a scalar attribute.",
  "subject": "/model",
  "args": {
    "error_detail": "\"groups.dirs.resources.files.myarray\" is not allowed to have \"matchversions\" set to \"true\" due to it not being a scalar attribute"
  },
  "source": "xxx"
}`},

		{"any - bad", `{
      "groups": { "dirs": { "singular": "dir", "resources": { "files": {
              "singular": "file",
              "attributes": {
                "myany": {
                  "type": "any",
                  "matchversions": true
                } } } } } } }`, `{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: \"groups.dirs.resources.files.myany\" is not allowed to have \"matchversions\" set to \"true\" due to it not being a scalar attribute.",
  "subject": "/model",
  "args": {
    "error_detail": "\"groups.dirs.resources.files.myany\" is not allowed to have \"matchversions\" set to \"true\" due to it not being a scalar attribute"
  },
  "source": "xxx"
}`},

		{"under array item - bad", `{
      "groups": { "dirs": { "singular": "dir", "resources": { "files": {
              "singular": "file",
              "attributes": {
                "myarray": {
                  "type": "array",
                  "item": { "type": "object",
                            "attributes": { "myint": {
                              "type": "integer", "matchversions": true }}}
                } } } } } } }`, `{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: \"groups.dirs.resources.files.myarray.item.myint\" is not allowed to have \"matchversions\" set to \"true\" due to it being in an array.",
  "subject": "/model",
  "args": {
    "error_detail": "\"groups.dirs.resources.files.myarray.item.myint\" is not allowed to have \"matchversions\" set to \"true\" due to it being in an array"
  },
  "source": "xxx"
}`},

		{"under map item - bad", `{
      "groups": { "dirs": { "singular": "dir", "resources": { "files": {
              "singular": "file",
              "attributes": {
                "mymap": {
                  "type": "map",
                  "item": { "type": "object",
                            "attributes": { "myint": {
                              "type": "integer", "matchversions": true }}}
                } } } } } } }`, `{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: \"groups.dirs.resources.files.mymap.item.myint\" is not allowed to have \"matchversions\" set to \"true\" due to it being in a map.",
  "subject": "/model",
  "args": {
    "error_detail": "\"groups.dirs.resources.files.mymap.item.myint\" is not allowed to have \"matchversions\" set to \"true\" due to it being in a map"
  },
  "source": "xxx"
}`},

		{"under ifvalues - bad", `{
      "groups": { "dirs": { "singular": "dir", "resources": { "files": {
              "singular": "file",
              "attributes": {
                "myint": {
                  "type": "integer",
                  "ifvalues": { "5": {
                      "siblingattributes": {
                        "mystr": { "type": "string", "matchversions": true }
                      } } } }
                } } } } } }`, `{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: \"groups.dirs.resources.files.myint.ifvalues.5.mystr\" is not allowed to have \"matchversions\" set to \"true\" due to it being in an \"ifvalues\" clause.",
  "subject": "/model",
  "args": {
    "error_detail": "\"groups.dirs.resources.files.myint.ifvalues.5.mystr\" is not allowed to have \"matchversions\" set to \"true\" due to it being in an \"ifvalues\" clause"
  },
  "source": "xxx"
}`},

		{"under wildcard - bad", `{
      "groups": { "dirs": { "singular": "dir", "resources": { "files": {
              "singular": "file",
              "attributes": {
                "*": { "type": "integer", "matchversions": true }
                } } } } } }`, `{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: \"groups.dirs.resources.files.*\" is not allowed to have \"matchversions\" set to \"true\" due to it being in a \"*\" extension.",
  "subject": "/model",
  "args": {
    "error_detail": "\"groups.dirs.resources.files.*\" is not allowed to have \"matchversions\" set to \"true\" due to it being in a \"*\" extension"
  },
  "source": "xxx"
}`},
	}

	for _, test := range tests {
		t.Logf("test: %s", test.name)
		m, xErr := ParseModel([]byte(test.src), nil)
		XNoErr(t, xErr)

		got := m.Verify()
		XCheckErr(t, got, test.err)
	}
}
