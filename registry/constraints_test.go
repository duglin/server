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

// TestConstraintsGroupTypeErrors verifies Group-type-level "constraints"
// definition errors (bad key format, unknown Resource type, bad path,
// enum/default/equals mismatches, etc). All pure model-verification
// errors - no DB/Registry needed.
func TestConstraintsGroupTypeErrors(t *testing.T) {
	// Let's test some error cases first

	// Missing period in key
	modelSrc := `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "filesmystr": {}
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "mystr": {
                    "type": "string",
                    "enum": [ "abc", "def", "ghi" ],
                    "strict": true
                  }
                } } } } } }`

	m, xErr := ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(),
		`{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"filesmystr\" has an invalid key. It must be of the form \"<RESOURCES>.<PATH>\".",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"filesmystr\" has an invalid key. It must be of the form \"<RESOURCES>.<PATH>\""
  },
  "source": "xxx"
}`)

	// Unknown RESOURCES
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "datas.files": {}
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "mystr": {
                    "type": "string",
                    "enum": [ "abc", "def", "ghi" ],
                    "strict": true
                  }
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(),
		`{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"datas.files\" has an unknown Resource type \"datas\".",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"datas.files\" has an unknown Resource type \"datas\""
  },
  "source": "xxx"
}`)

	// Bad path - no attr
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.": {}
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "mystr": {
                    "type": "string",
                    "enum": [ "abc", "def", "ghi" ],
                    "strict": true
                  }
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(),
		`{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"files.\" has an empty reference to an attribute.",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"files.\" has an empty reference to an attribute"
  },
  "source": "xxx"
}`)

	// Bad path - syntax error
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.fh'213'": {}
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "mystr": {
                    "type": "string",
                    "enum": [ "abc", "def", "ghi" ],
                    "strict": true
                  }
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(),
		`{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"files.fh'213'\" has an invalid path (fh'213'): Unexpected \"'\" in \"fh'213'\" at pos 3.",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"files.fh'213'\" has an invalid path (fh'213'): Unexpected \"'\" in \"fh'213'\" at pos 3"
  },
  "source": "xxx"
}`)

	// Unknown attribute - root
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.foo": {}
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "mystr": {
                    "type": "string",
                    "enum": [ "abc", "def", "ghi" ],
                    "strict": true
                  }
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(),
		`{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"files.foo\" has an invalid path (foo): Attribute \"foo\" can not be found.",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"files.foo\" has an invalid path (foo): Attribute \"foo\" can not be found"
  },
  "source": "xxx"
}`)

	// Unknown attribute - in obj
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.mystr.foo": {}
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "mystr": {
                    "type": "string",
                    "enum": [ "abc", "def", "ghi" ],
                    "strict": true
                  }
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(),
		`{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"files.mystr.foo\" has an invalid path (mystr.foo): Attribute \"mystr\" is scalar, so \"foo\" is invalid.",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"files.mystr.foo\" has an invalid path (mystr.foo): Attribute \"mystr\" is scalar, so \"foo\" is invalid"
  },
  "source": "xxx"
}`)

	// Unknown attr, step into map - even though it's not valid for constraints
	// just test our infra
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.mymap.foo": {}
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "mymap": {
                    "type": "map",
                    "item": {
                      "type": "string"
                    }
                  }
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(),
		`{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"files.mymap.foo\" has an invalid path (mymap.foo): Attribute \"foo\" can not be found in \"mymap\".",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"files.mymap.foo\" has an invalid path (mymap.foo): Attribute \"foo\" can not be found in \"mymap\""
  },
  "source": "xxx"
}`)

	// Valid attr, but Step into map - which is not allowed
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.mymap.foo": {}
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "mymap": {
                    "type": "map",
                    "item": {
                      "type": "object",
                      "attributes": {
                        "foo": {
                          "type": "integer"
                        }
                      }
                    }
                  }
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(),
		`{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"files.mymap.foo\" has a path (mymap.foo) that includes a map (mymap), which is not allowed.",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"files.mymap.foo\" has a path (mymap.foo) that includes a map (mymap), which is not allowed"
  },
  "source": "xxx"
}`)

	// Unknown attr, step into array-even though it's not valid for constraints
	// just test our infra
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.myarray.foo": {}
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "myarray": {
                    "type": "array",
                    "item": {
                      "type": "object"
                    }
                  }
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(),
		`{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"files.myarray.foo\" has an invalid path (myarray.foo): Attribute \"foo\" can not be found in \"myarray\".",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"files.myarray.foo\" has an invalid path (myarray.foo): Attribute \"foo\" can not be found in \"myarray\""
  },
  "source": "xxx"
}`)

	// Step into array - not allowed
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.myarray.foo": {}
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "myarray": {
                    "type": "array",
                    "item": {
                      "type": "object",
                      "attributes": {
                        "foo": { "type": "integer" }
                      }
                    }
                  }
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(),
		`{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"files.myarray.foo\" has a path (myarray.foo) that includes an array (myarray), which is not allowed.",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"files.myarray.foo\" has a path (myarray.foo) that includes an array (myarray), which is not allowed"
  },
  "source": "xxx"
}`)

	// Unknown step into empty object
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.myobj.foo": {}
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "myobj": {
                    "type": "object"
                  }
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(),
		`{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"files.myobj.foo\" has an invalid path (myobj.foo): Attribute \"foo\" can not be found in \"myobj\".",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"files.myobj.foo\" has an invalid path (myobj.foo): Attribute \"foo\" can not be found in \"myobj\""
  },
  "source": "xxx"
}`)

	// Stop on object
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.myobj": {}
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "myobj": {
                    "type": "object"
                  }
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(),
		`{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"files.myobj\" has an invalid path (myobj): \"myobj\" must be a scalar.",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"files.myobj\" has an invalid path (myobj): \"myobj\" must be a scalar"
  },
  "source": "xxx"
}`)

	// Validate Enum list - strict, can't extend
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.mystr": { "enum": [ "abc", "bye" ] }
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "mystr": { "type": "string", "enum": [ "abc", "def" ] }
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(),
		`{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"files.mystr\" has an enum value (bye) that isn't part of the inherited attribute's enum list (abc, def).",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"files.mystr\" has an enum value (bye) that isn't part of the inherited attribute's enum list (abc, def)"
  },
  "source": "xxx"
}`)

	// Validate Enum list - not strict, can extend
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.mystr": { "enum": [ "abc", "bye" ] }
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "mystr": { "type":"string", "enum":["abc"],"strict":false }
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XNoErr(t, m.Verify())

	// Validate Enum list - empty enum is same as missing
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.mystr": { "enum": []  }
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "mystr": { "type":"string", "enum":["abc"]}
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XNoErr(t, m.Verify())

	// Validate Enum list - can add even when no enum on base attr
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.mystr": { "enum": [ "abc" ]  }
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "mystr": { "type":"string" }
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XNoErr(t, m.Verify())

	// Validate Enum list - constraint.enum must include attr.default if no
	// new default is defined
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.mystr": { "enum": [ "abc", "bye" ] }
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "mystr": {"type":"string", "default":"def","required": true }
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(), `{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"files.mystr\" has an enum set (abc, bye) that doesn't include the attribute's default value (def).",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"files.mystr\" has an enum set (abc, bye) that doesn't include the attribute's default value (def)"
  },
  "source": "xxx"
}`)

	// Validate Enum list - constraint.enum must include attr.default if no
	// new default is defined
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.mystr": { "enum": [ "abc" ] }
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "mystr": {"type":"string", "default":"def",
                    "enum": [ "abc", "def" ], "required": true }
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(), `{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"files.mystr\" has an enum set (abc) that doesn't include the attribute's default value (def).",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"files.mystr\" has an enum set (abc) that doesn't include the attribute's default value (def)"
  },
  "source": "xxx"
}`)

	// Validate Enum list - extend attr w/enum + default
	// new default is defined
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.mystr": { "enum": [ "abc", "def" ], "default": "def" }
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "mystr": {"type":"string"}
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XNoErr(t, m.Verify())

	// Validate Enum list - all enum values of the right/same type
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.mystr": { "enum": [ "abc", 123 ] }
              },
              "resources": {"files": {"singular": "file",
                "attributes": {
                  "mystr": {"type":"string"}
                } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(), `{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"files.mystr\" has an enum value (123) that must be of type \"string\".",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"files.mystr\" has an enum value (123) that must be of type \"string\""
  },
  "source": "xxx"
}`)

	// Validate Equals - "" == ignore
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.mystr": { "equals": "" }
              },
              "resources": {"files": {"singular": "file",
                "attributes": { "mystr": {"type":"string"} } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XNoErr(t, m.Verify())

	// Validate Equals - missing attr
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.mystr": { "equals": "foo" }
              },
              "resources": {"files": {"singular": "file",
                "attributes": { "mystr": {"type":"string"} } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(), `{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"files.mystr\" has an \"equals\" reference (foo) that isn't valid: Attribute \"foo\" can not be found.",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"files.mystr\" has an \"equals\" reference (foo) that isn't valid: Attribute \"foo\" can not be found"
  },
  "source": "xxx"
}`)

	// Validate Equals - non-scalar attr
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.mystr": { "equals": "labels" }
              },
              "resources": {"files": {"singular": "file",
                "attributes": { "mystr": {"type":"string"} } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(), `{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"files.mystr\" has an \"equals\" reference (labels) that must be a scalar.",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"files.mystr\" has an \"equals\" reference (labels) that must be a scalar"
  },
  "source": "xxx"
}`)

	// Validate Equals - must be same type
	modelSrc = `{
      "groups": { "dirs": {
              "singular": "dir",
              "constraints": {
                "files.mystr": { "equals": "epoch" }
              },
              "resources": {"files": {"singular": "file",
                "attributes": { "mystr": {"type":"string"} } } } } } }`

	m, xErr = ParseModel([]byte(modelSrc), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(), `{
  "type": "https://github.com/xregistry/spec/blob/main/core/spec.md#model_error",
  "title": "There was an error in the model definition provided: Group Type \"dirs\" constraint \"files.mystr\" references an attribute of type \"string\" but its \"equals\" (epoch) references an attribute of type \"uinteger\". They need to match.",
  "subject": "/model",
  "args": {
    "error_detail": "Group Type \"dirs\" constraint \"files.mystr\" references an attribute of type \"string\" but its \"equals\" (epoch) references an attribute of type \"uinteger\". They need to match"
  },
  "source": "xxx"
}`)

}

// TestConstraintsModelEqualsPathThroughIfvaluesOnlyAttr verifies that an
// "equals" path segment can't traverse through an attribute that's only
// defined via an "ifvalues" siblingattributes clause (not a static attr).
// Pure model-verification error - no DB/Registry needed.
func TestConstraintsModelEqualsPathThroughIfvaluesOnlyAttr(t *testing.T) {
	// Model error: equals path traverses through ifvalues-only attribute
	// (gfoo is defined only via ifvalues siblingattributes, not static attrs)
	modelSrcBad := `{
	  "groups": { "dirs": {
	    "singular": "dir",
	    "attributes": {
	      "gobj": {
	        "type": "object",
	        "ifvalues": {
	          "special": {
	            "siblingattributes": { "gfoo": { "type": "string" } }
	          }
	        }
	      }
	    },
	    "constraints": {
	      "files.mystr": { "equals": "gfoo" }
	    },
	    "resources": {"files": {"singular": "file", "hasdocument": false,
	      "attributes": { "mystr": { "type": "string" } } } } } } }`

	m, xErr := ParseModel([]byte(modelSrcBad), nil)
	XNoErr(t, xErr)
	XCheckErr(t, m.Verify(),
		`^(?s)^.*model_error.*equals.*gfoo.*can not be found`)
}
