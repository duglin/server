package tests

import (
	"testing"

	"github.com/duglin/xreg-github/registry"
)

func TestMetaSimple(t *testing.T) {
	reg := NewRegistry("TestMetaSimple")
	defer PassDeleteReg(t, reg)

	gm, _ := reg.Model.AddGroupModel("dirs", "dir")
	rm, _ := gm.AddResourceModel("files", "file", 0, true, true, false) // noDoc
	rm.AddMetaAttr("foo", registry.ANY)

	// Simple no body create PUT
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f1/meta",
		Method: "PUT",
		Code:   201,
		ResHeaders: []string{
			"Location: http://localhost:8181/dirs/d1/files/f1/meta",
		},
		ResBody: `{
  "fileid": "f1",
  "self": "http://localhost:8181/dirs/d1/files/f1/meta",
  "xid": "/dirs/d1/files/f1/meta",
  "epoch": 1,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:01Z",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/1"
}
`,
	})

	xHTTP(t, reg, "GET", "/dirs/d1/files/f1?inline", ``, 200, `{
  "fileid": "f1",
  "versionid": "1",
  "self": "http://localhost:8181/dirs/d1/files/f1",
  "xid": "/dirs/d1/files/f1",
  "epoch": 1,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:01Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f1/meta",
  "meta": {
    "fileid": "f1",
    "self": "http://localhost:8181/dirs/d1/files/f1/meta",
    "xid": "/dirs/d1/files/f1/meta",
    "epoch": 1,
    "createdat": "2024-01-01T12:00:01Z",
    "modifiedat": "2024-01-01T12:00:01Z",

    "defaultversionid": "1",
    "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/1"
  },
  "versionsurl": "http://localhost:8181/dirs/d1/files/f1/versions",
  "versions": {
    "1": {
      "fileid": "f1",
      "versionid": "1",
      "self": "http://localhost:8181/dirs/d1/files/f1/versions/1",
      "xid": "/dirs/d1/files/f1/versions/1",
      "epoch": 1,
      "isdefault": true,
      "createdat": "2024-01-01T12:00:01Z",
      "modifiedat": "2024-01-01T12:00:01Z"
    }
  },
  "versionscount": 1
}
`)

	// Simple create no body POST - error
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f11/meta",
		Method: "POST",
		Code:   405,
		ResBody: `POST not allowed on a 'meta'
`,
	})

	// Simple create no body PUT, URL too long
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f11/meta/xxx",
		Method: "PUT",
		Code:   404,
		ResBody: `URL is too long
`,
	})

	// Simple create no body POST, URL too long
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f11/meta/xxx",
		Method: "POST",
		Code:   404,
		ResBody: `URL is too long
`,
	})

	// Simple create no body PATCH, URL too long
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f11/meta/xxx",
		Method: "PATCH",
		Code:   404,
		ResBody: `URL is too long
`,
	})

	// Simple create no body PATCH
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f12/meta",
		Method: "PUT",
		Code:   201,
		ResHeaders: []string{
			"Location: http://localhost:8181/dirs/d1/files/f12/meta",
		},
		ResBody: `{
  "fileid": "f12",
  "self": "http://localhost:8181/dirs/d1/files/f12/meta",
  "xid": "/dirs/d1/files/f12/meta",
  "epoch": 1,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:01Z",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f12/versions/1"
}
`,
	})

	// Simple body create PUT + ext
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f2/meta",
		Method: "PUT",
		ReqBody: `{
  "foo": "bar"
}
`,
		Code: 201,
		ResHeaders: []string{
			"Location: http://localhost:8181/dirs/d1/files/f2/meta",
		},
		ResBody: `{
  "fileid": "f2",
  "self": "http://localhost:8181/dirs/d1/files/f2/meta",
  "xid": "/dirs/d1/files/f2/meta",
  "epoch": 1,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:01Z",
  "foo": "bar",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f2/versions/1"
}
`,
	})

	// Simple body create PATCH + ext
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f21/meta",
		Method: "PUT",
		ReqBody: `{
  "foo": "bar"
}
`,
		Code: 201,
		ResHeaders: []string{
			"Location: http://localhost:8181/dirs/d1/files/f21/meta",
		},
		ResBody: `{
  "fileid": "f21",
  "self": "http://localhost:8181/dirs/d1/files/f21/meta",
  "xid": "/dirs/d1/files/f21/meta",
  "epoch": 1,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:01Z",
  "foo": "bar",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f21/versions/1"
}
`,
	})

	// Update PUT no body - erases ext
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f2/meta",
		Method: "PUT",
		Code:   200,
		ResHeaders: []string{
			"-Location",
		},
		ResBody: `{
  "fileid": "f2",
  "self": "http://localhost:8181/dirs/d1/files/f2/meta",
  "xid": "/dirs/d1/files/f2/meta",
  "epoch": 2,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f2/versions/1"
}
`,
	})

	xHTTP(t, reg, "GET", "/dirs/d1/files/f2?inline", ``, 200, `{
  "fileid": "f2",
  "versionid": "1",
  "self": "http://localhost:8181/dirs/d1/files/f2",
  "xid": "/dirs/d1/files/f2",
  "epoch": 1,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:01Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f2/meta",
  "meta": {
    "fileid": "f2",
    "self": "http://localhost:8181/dirs/d1/files/f2/meta",
    "xid": "/dirs/d1/files/f2/meta",
    "epoch": 2,
    "createdat": "2024-01-01T12:00:01Z",
    "modifiedat": "2024-01-01T12:00:02Z",

    "defaultversionid": "1",
    "defaultversionurl": "http://localhost:8181/dirs/d1/files/f2/versions/1"
  },
  "versionsurl": "http://localhost:8181/dirs/d1/files/f2/versions",
  "versions": {
    "1": {
      "fileid": "f2",
      "versionid": "1",
      "self": "http://localhost:8181/dirs/d1/files/f2/versions/1",
      "xid": "/dirs/d1/files/f2/versions/1",
      "epoch": 1,
      "isdefault": true,
      "createdat": "2024-01-01T12:00:01Z",
      "modifiedat": "2024-01-01T12:00:01Z"
    }
  },
  "versionscount": 1
}
`)

	// Update PATCH no body
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f21/meta",
		Method: "PATCH",
		Code:   200,
		ResHeaders: []string{
			"-Location",
		},
		ResBody: `{
  "fileid": "f21",
  "self": "http://localhost:8181/dirs/d1/files/f21/meta",
  "xid": "/dirs/d1/files/f21/meta",
  "epoch": 2,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",
  "foo": "bar",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f21/versions/1"
}
`,
	})

	// Update PUT empty body
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1/meta",
		Method:  "PUT",
		ReqBody: `{}`,
		Code:    200,
		ResHeaders: []string{
			"-Location",
		},
		ResBody: `{
  "fileid": "f1",
  "self": "http://localhost:8181/dirs/d1/files/f1/meta",
  "xid": "/dirs/d1/files/f1/meta",
  "epoch": 2,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/1"
}
`,
	})

	// Update PATCH empty body
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f21/meta",
		Method:  "PATCH",
		ReqBody: `{}`,
		Code:    200,
		ResHeaders: []string{
			"-Location",
		},
		ResBody: `{
  "fileid": "f21",
  "self": "http://localhost:8181/dirs/d1/files/f21/meta",
  "xid": "/dirs/d1/files/f21/meta",
  "epoch": 3,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",
  "foo": "bar",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f21/versions/1"
}
`,
	})

	// Update PUT + ext
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1/meta",
		Method:  "PUT",
		ReqBody: `{ "foo": "zzz"}`,
		Code:    200,
		ResHeaders: []string{
			"-Location",
		},
		ResBody: `{
  "fileid": "f1",
  "self": "http://localhost:8181/dirs/d1/files/f1/meta",
  "xid": "/dirs/d1/files/f1/meta",
  "epoch": 3,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",
  "foo": "zzz",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/1"
}
`,
	})

	// Update PATCH empty body
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f21/meta",
		Method:  "PATCH",
		ReqBody: `{"foo":"aaa"}`,
		Code:    200,
		ResHeaders: []string{
			"-Location",
		},
		ResBody: `{
  "fileid": "f21",
  "self": "http://localhost:8181/dirs/d1/files/f21/meta",
  "xid": "/dirs/d1/files/f21/meta",
  "epoch": 4,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",
  "foo": "aaa",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f21/versions/1"
}
`,
	})

	// Update PUT + bad ext
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1/meta",
		Method:  "PUT",
		ReqBody: `{ "fff": "zzz"}`,
		Code:    400,
		ResBody: `Invalid extension(s): fff
`,
	})

	// Update PATCH empty body
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f21/meta",
		Method:  "PATCH",
		ReqBody: `{"fff":"aaa"}`,
		Code:    400,
		ResBody: `Invalid extension(s): fff
`,
	})

	// Update PUT, del ext
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1/meta",
		Method:  "PUT",
		ReqBody: `{}`,
		Code:    200,
		ResBody: `{
  "fileid": "f1",
  "self": "http://localhost:8181/dirs/d1/files/f1/meta",
  "xid": "/dirs/d1/files/f1/meta",
  "epoch": 4,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/1"
}
`,
	})

	// Update PATCH, del ext
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f21/meta",
		Method:  "PATCH",
		ReqBody: `{"foo":null}`,
		Code:    200,
		ResBody: `{
  "fileid": "f21",
  "self": "http://localhost:8181/dirs/d1/files/f21/meta",
  "xid": "/dirs/d1/files/f21/meta",
  "epoch": 5,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f21/versions/1"
}
`,
	})

	// Update PUT, add ext again
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1/meta",
		Method:  "PUT",
		ReqBody: `{ "foo": "zz1"}`,
		Code:    200,
		ResBody: `{
  "fileid": "f1",
  "self": "http://localhost:8181/dirs/d1/files/f1/meta",
  "xid": "/dirs/d1/files/f1/meta",
  "epoch": 5,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",
  "foo": "zz1",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/1"
}
`,
	})

	// Update PUT, del ext again
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1/meta",
		Method:  "PUT",
		ReqBody: `{"foo":null}`,
		Code:    200,
		ResBody: `{
  "fileid": "f1",
  "self": "http://localhost:8181/dirs/d1/files/f1/meta",
  "xid": "/dirs/d1/files/f1/meta",
  "epoch": 6,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/1"
}
`,
	})

	// Make sure DELETE fails
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f1/meta",
		Method: "DELETE",
		Code:   405,
		ResBody: `DELETE is not allowed on a "meta"
`,
	})

}

func TestMetaCombos(t *testing.T) {
	reg := NewRegistry("TestMetaCombos")
	defer PassDeleteReg(t, reg)

	gm, _ := reg.Model.AddGroupModel("dirs", "dir")
	rm, _ := gm.AddResourceModel("files", "file", 0, true, true, false) // noDoc
	rm.AddMetaAttr("foo", registry.ANY)

	// Create Resource and set the versionID
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1",
		Method:  "PUT",
		ReqBody: `{"versionid":"v1.0"}`,
		Code:    201,
		ResHeaders: []string{
			"Location: http://localhost:8181/dirs/d1/files/f1",
		},
		ResBody: `{
  "fileid": "f1",
  "versionid": "v1.0",
  "self": "http://localhost:8181/dirs/d1/files/f1",
  "xid": "/dirs/d1/files/f1",
  "epoch": 1,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:01Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f1/meta",
  "versionsurl": "http://localhost:8181/dirs/d1/files/f1/versions",
  "versionscount": 1
}
`,
	})

	// Verify it's all correct
	xHTTP(t, reg, "GET", "/dirs/d1/files/f1?inline", ``, 200, `{
  "fileid": "f1",
  "versionid": "v1.0",
  "self": "http://localhost:8181/dirs/d1/files/f1",
  "xid": "/dirs/d1/files/f1",
  "epoch": 1,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:01Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f1/meta",
  "meta": {
    "fileid": "f1",
    "self": "http://localhost:8181/dirs/d1/files/f1/meta",
    "xid": "/dirs/d1/files/f1/meta",
    "epoch": 1,
    "createdat": "2024-01-01T12:00:01Z",
    "modifiedat": "2024-01-01T12:00:01Z",

    "defaultversionid": "v1.0",
    "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/v1.0"
  },
  "versionsurl": "http://localhost:8181/dirs/d1/files/f1/versions",
  "versions": {
    "v1.0": {
      "fileid": "f1",
      "versionid": "v1.0",
      "self": "http://localhost:8181/dirs/d1/files/f1/versions/v1.0",
      "xid": "/dirs/d1/files/f1/versions/v1.0",
      "epoch": 1,
      "isdefault": true,
      "createdat": "2024-01-01T12:00:01Z",
      "modifiedat": "2024-01-01T12:00:01Z"
    }
  },
  "versionscount": 1
}
`)

	// PUT again with wrong versionid should fail
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1",
		Method:  "PUT",
		ReqBody: `{"versionid":"v2.0"}`,
		Code:    400,
		ResBody: `When "versionid"(v2.0) is present it must match the "defaultversionid"(v1.0)
`,
	})

	// PUT again with wrong fileid should fail
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1",
		Method:  "PUT",
		ReqBody: `{"fileid":"foo"}`,
		Code:    400,
		ResBody: `The "fileid" attribute must be set to "f1", not "foo"
`,
	})

	// PUT on meta with wrong fileid should fail
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1/meta",
		Method:  "PUT",
		ReqBody: `{"fileid":"foo"}`,
		Code:    400,
		ResBody: `meta.fileid must be "f1", not "foo"
`,
	})

	// Create a version, setting vid
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1",
		Method:  "POST",
		ReqBody: `{"versionid":"v2.0"}`,
		Code:    201,
		ResHeaders: []string{
			"Location: http://localhost:8181/dirs/d1/files/f1/versions/v2.0",
		},
		ResBody: `{
  "fileid": "f1",
  "versionid": "v2.0",
  "self": "http://localhost:8181/dirs/d1/files/f1/versions/v2.0",
  "xid": "/dirs/d1/files/f1/versions/v2.0",
  "epoch": 1,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:01Z"
}
`,
	})

	// Verify
	xHTTP(t, reg, "GET", "/dirs/d1/files/f1?inline", ``, 200, `{
  "fileid": "f1",
  "versionid": "v2.0",
  "self": "http://localhost:8181/dirs/d1/files/f1",
  "xid": "/dirs/d1/files/f1",
  "epoch": 1,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:02Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f1/meta",
  "meta": {
    "fileid": "f1",
    "self": "http://localhost:8181/dirs/d1/files/f1/meta",
    "xid": "/dirs/d1/files/f1/meta",
    "epoch": 2,
    "createdat": "2024-01-01T12:00:01Z",
    "modifiedat": "2024-01-01T12:00:02Z",

    "defaultversionid": "v2.0",
    "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/v2.0"
  },
  "versionsurl": "http://localhost:8181/dirs/d1/files/f1/versions",
  "versions": {
    "v1.0": {
      "fileid": "f1",
      "versionid": "v1.0",
      "self": "http://localhost:8181/dirs/d1/files/f1/versions/v1.0",
      "xid": "/dirs/d1/files/f1/versions/v1.0",
      "epoch": 1,
      "createdat": "2024-01-01T12:00:01Z",
      "modifiedat": "2024-01-01T12:00:01Z"
    },
    "v2.0": {
      "fileid": "f1",
      "versionid": "v2.0",
      "self": "http://localhost:8181/dirs/d1/files/f1/versions/v2.0",
      "xid": "/dirs/d1/files/f1/versions/v2.0",
      "epoch": 1,
      "isdefault": true,
      "createdat": "2024-01-01T12:00:02Z",
      "modifiedat": "2024-01-01T12:00:02Z"
    }
  },
  "versionscount": 2
}
`)

	// Update/PUT w/o body should just bump epoch/modifiedat
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f1",
		Method: "PUT",
		Code:   200,
		ResBody: `{
  "fileid": "f1",
  "versionid": "v2.0",
  "self": "http://localhost:8181/dirs/d1/files/f1",
  "xid": "/dirs/d1/files/f1",
  "epoch": 2,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f1/meta",
  "versionsurl": "http://localhost:8181/dirs/d1/files/f1/versions",
  "versionscount": 2
}
`,
	})

	// Make sure resource's epoch didn't change
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f1?inline=meta",
		Method: "GET",
		Code:   200,
		ResBody: `{
  "fileid": "f1",
  "versionid": "v2.0",
  "self": "http://localhost:8181/dirs/d1/files/f1",
  "xid": "/dirs/d1/files/f1",
  "epoch": 2,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f1/meta",
  "meta": {
    "fileid": "f1",
    "self": "http://localhost:8181/dirs/d1/files/f1/meta",
    "xid": "/dirs/d1/files/f1/meta",
    "epoch": 2,
    "createdat": "2024-01-01T12:00:03Z",
    "modifiedat": "2024-01-01T12:00:01Z",

    "defaultversionid": "v2.0",
    "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/v2.0"
  },
  "versionsurl": "http://localhost:8181/dirs/d1/files/f1/versions",
  "versionscount": 2
}
`,
	})

	// Update/PUT - valid vid
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1",
		Method:  "PUT",
		ReqBody: `{"versionid": "v2.0"}`,
		Code:    200,
		ResBody: `{
  "fileid": "f1",
  "versionid": "v2.0",
  "self": "http://localhost:8181/dirs/d1/files/f1",
  "xid": "/dirs/d1/files/f1",
  "epoch": 3,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f1/meta",
  "versionsurl": "http://localhost:8181/dirs/d1/files/f1/versions",
  "versionscount": 2
}
`,
	})

	// Make sure just version's epoch/timestamp changed
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f1?inline=meta",
		Method: "GET",
		Code:   200,
		ResBody: `{
  "fileid": "f1",
  "versionid": "v2.0",
  "self": "http://localhost:8181/dirs/d1/files/f1",
  "xid": "/dirs/d1/files/f1",
  "epoch": 3,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f1/meta",
  "meta": {
    "fileid": "f1",
    "self": "http://localhost:8181/dirs/d1/files/f1/meta",
    "xid": "/dirs/d1/files/f1/meta",
    "epoch": 2,
    "createdat": "2024-01-01T12:00:03Z",
    "modifiedat": "2024-01-01T12:00:01Z",

    "defaultversionid": "v2.0",
    "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/v2.0"
  },
  "versionsurl": "http://localhost:8181/dirs/d1/files/f1/versions",
  "versionscount": 2
}
`,
	})

	// Update/PUT - make defaultversionid sticky
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1",
		Method:  "PUT",
		ReqBody: `{"meta":{"defaultversionsticky":true}}`,
		Code:    200,
		ResBody: `{
  "fileid": "f1",
  "versionid": "v2.0",
  "self": "http://localhost:8181/dirs/d1/files/f1",
  "xid": "/dirs/d1/files/f1",
  "epoch": 4,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f1/meta",
  "versionsurl": "http://localhost:8181/dirs/d1/files/f1/versions",
  "versionscount": 2
}
`,
	})

	// defversticky changed, but so did the default version's epoch/timestamp
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f1?inline=meta",
		Method: "GET",
		Code:   200,
		ResBody: `{
  "fileid": "f1",
  "versionid": "v2.0",
  "self": "http://localhost:8181/dirs/d1/files/f1",
  "xid": "/dirs/d1/files/f1",
  "epoch": 4,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f1/meta",
  "meta": {
    "fileid": "f1",
    "self": "http://localhost:8181/dirs/d1/files/f1/meta",
    "xid": "/dirs/d1/files/f1/meta",
    "epoch": 3,
    "createdat": "2024-01-01T12:00:03Z",
    "modifiedat": "2024-01-01T12:00:02Z",

    "defaultversionid": "v2.0",
    "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/v2.0",
    "defaultversionsticky": true
  },
  "versionsurl": "http://localhost:8181/dirs/d1/files/f1/versions",
  "versionscount": 2
}
`,
	})

	// Make sure just version's epoch/timestamp changed
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f1?inline=meta",
		Method: "GET",
		Code:   200,
		ResBody: `{
  "fileid": "f1",
  "versionid": "v2.0",
  "self": "http://localhost:8181/dirs/d1/files/f1",
  "xid": "/dirs/d1/files/f1",
  "epoch": 4,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f1/meta",
  "meta": {
    "fileid": "f1",
    "self": "http://localhost:8181/dirs/d1/files/f1/meta",
    "xid": "/dirs/d1/files/f1/meta",
    "epoch": 3,
    "createdat": "2024-01-01T12:00:03Z",
    "modifiedat": "2024-01-01T12:00:02Z",

    "defaultversionid": "v2.0",
    "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/v2.0",
    "defaultversionsticky": true
  },
  "versionsurl": "http://localhost:8181/dirs/d1/files/f1/versions",
  "versionscount": 2
}
`,
	})

	// Update/PUT - def attrs at the wrong spot
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1",
		Method:  "PUT",
		ReqBody: `{"defaultversionid": "v1.0","defaultversionsticky":true}`,
		Code:    400,
		ResBody: `Invalid extension(s): defaultversionid,defaultversionsticky
`,
	})

	// Update/PUT - stick it again via meta this time
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1/meta",
		Method:  "PUT",
		ReqBody: `{"defaultversionsticky":true}`,
		Code:    200,
		ResBody: `{
  "fileid": "f1",
  "self": "http://localhost:8181/dirs/d1/files/f1/meta",
  "xid": "/dirs/d1/files/f1/meta",
  "epoch": 4,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "defaultversionid": "v2.0",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/v2.0",
  "defaultversionsticky": true
}
`,
	})

	// meta's epoch changed but the defver didn't
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f1?inline=meta",
		Method: "GET",
		Code:   200,
		ResBody: `{
  "fileid": "f1",
  "versionid": "v2.0",
  "self": "http://localhost:8181/dirs/d1/files/f1",
  "xid": "/dirs/d1/files/f1",
  "epoch": 4,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f1/meta",
  "meta": {
    "fileid": "f1",
    "self": "http://localhost:8181/dirs/d1/files/f1/meta",
    "xid": "/dirs/d1/files/f1/meta",
    "epoch": 4,
    "createdat": "2024-01-01T12:00:03Z",
    "modifiedat": "2024-01-01T12:00:04Z",

    "defaultversionid": "v2.0",
    "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/v2.0",
    "defaultversionsticky": true
  },
  "versionsurl": "http://localhost:8181/dirs/d1/files/f1/versions",
  "versionscount": 2
}
`,
	})

	// Create new version, defverid should not change, nor meta epoch/ts.
	// New vid should be generated - ie '1'
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f1",
		Method: "POST",
		Code:   201,
		ResBody: `{
  "fileid": "f1",
  "versionid": "1",
  "self": "http://localhost:8181/dirs/d1/files/f1/versions/1",
  "xid": "/dirs/d1/files/f1/versions/1",
  "epoch": 1,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:01Z"
}
`,
	})

	// defver and meta should be unchanged
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f1?inline=meta",
		Method: "GET",
		Code:   200,
		ResBody: `{
  "fileid": "f1",
  "versionid": "v2.0",
  "self": "http://localhost:8181/dirs/d1/files/f1",
  "xid": "/dirs/d1/files/f1",
  "epoch": 4,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f1/meta",
  "meta": {
    "fileid": "f1",
    "self": "http://localhost:8181/dirs/d1/files/f1/meta",
    "xid": "/dirs/d1/files/f1/meta",
    "epoch": 5,
    "createdat": "2024-01-01T12:00:03Z",
    "modifiedat": "2024-01-01T12:00:04Z",

    "defaultversionid": "v2.0",
    "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/v2.0",
    "defaultversionsticky": true
  },
  "versionsurl": "http://localhost:8181/dirs/d1/files/f1/versions",
  "versionscount": 3
}
`,
	})

	// Update/PUT - unstick it, '1' should be the def now
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1/meta",
		Method:  "PUT",
		ReqBody: `{"defaultversionsticky":false}`,
		Code:    200,
		ResBody: `{
  "fileid": "f1",
  "self": "http://localhost:8181/dirs/d1/files/f1/meta",
  "xid": "/dirs/d1/files/f1/meta",
  "epoch": 6,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/1"
}
`,
	})

	// meta's epoch changed but the defver didn't
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f1?inline=meta",
		Method: "GET",
		Code:   200,
		ResBody: `{
  "fileid": "f1",
  "versionid": "1",
  "self": "http://localhost:8181/dirs/d1/files/f1",
  "xid": "/dirs/d1/files/f1",
  "epoch": 1,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:01Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f1/meta",
  "meta": {
    "fileid": "f1",
    "self": "http://localhost:8181/dirs/d1/files/f1/meta",
    "xid": "/dirs/d1/files/f1/meta",
    "epoch": 6,
    "createdat": "2024-01-01T12:00:03Z",
    "modifiedat": "2024-01-01T12:00:02Z",

    "defaultversionid": "1",
    "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/1"
  },
  "versionsurl": "http://localhost:8181/dirs/d1/files/f1/versions",
  "versionscount": 3
}
`,
	})

	// Update/PUT - update it via 'defverid' - should err since not sticky
	// except bump epoch
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1/meta",
		Method:  "PUT",
		ReqBody: `{"defaultversionid":"v1.0"}`,
		Code:    400,
		ResBody: `Attribute "defaultversionid" must be "1" since "defaultversionsticky" is "false"
`,
	})

	// Update/PUT - stick it via 'defverid' AND sticky
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1/meta",
		Method:  "PUT",
		ReqBody: `{"defaultversionid":"v1.0","defaultversionsticky":true}`,
		Code:    200,
		ResBody: `{
  "fileid": "f1",
  "self": "http://localhost:8181/dirs/d1/files/f1/meta",
  "xid": "/dirs/d1/files/f1/meta",
  "epoch": 7,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "defaultversionid": "v1.0",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/v1.0",
  "defaultversionsticky": true
}
`,
	})

	// meta's epoch changed but the defver didn't
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f1?inline=meta",
		Method: "GET",
		Code:   200,
		ResBody: `{
  "fileid": "f1",
  "versionid": "v1.0",
  "self": "http://localhost:8181/dirs/d1/files/f1",
  "xid": "/dirs/d1/files/f1",
  "epoch": 1,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:01Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f1/meta",
  "meta": {
    "fileid": "f1",
    "self": "http://localhost:8181/dirs/d1/files/f1/meta",
    "xid": "/dirs/d1/files/f1/meta",
    "epoch": 7,
    "createdat": "2024-01-01T12:00:01Z",
    "modifiedat": "2024-01-01T12:00:02Z",

    "defaultversionid": "v1.0",
    "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/v1.0",
    "defaultversionsticky": true
  },
  "versionsurl": "http://localhost:8181/dirs/d1/files/f1/versions",
  "versionscount": 3
}
`,
	})

	// Update/PUT - change defverid/unstick - error
	// Include extension
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f1/meta",
		Method: "PUT",
		ReqBody: `{
		  "defaultversionid":"v2.0",
		  "defaultversionsticky":null,
		  "foo":"bar"}`,
		Code: 400,
		ResBody: `Attribute "defaultversionid" must be "1" since "defaultversionsticky" is "false"
`,
	})

	// Update/PUT - unstick
	// Include extension
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1/meta",
		Method:  "PUT",
		ReqBody: `{ "defaultversionsticky":null, "foo":"bar" }`,
		Code:    200,
		ResBody: `{
  "fileid": "f1",
  "self": "http://localhost:8181/dirs/d1/files/f1/meta",
  "xid": "/dirs/d1/files/f1/meta",
  "epoch": 8,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",
  "foo": "bar",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/1"
}
`,
	})

	// meta's epoch changed but the defver didn't
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f1?inline=meta",
		Method: "GET",
		Code:   200,
		ResBody: `{
  "fileid": "f1",
  "versionid": "1",
  "self": "http://localhost:8181/dirs/d1/files/f1",
  "xid": "/dirs/d1/files/f1",
  "epoch": 1,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:01Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f1/meta",
  "meta": {
    "fileid": "f1",
    "self": "http://localhost:8181/dirs/d1/files/f1/meta",
    "xid": "/dirs/d1/files/f1/meta",
    "epoch": 8,
    "createdat": "2024-01-01T12:00:03Z",
    "modifiedat": "2024-01-01T12:00:02Z",
    "foo": "bar",

    "defaultversionid": "1",
    "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/1"
  },
  "versionsurl": "http://localhost:8181/dirs/d1/files/f1/versions",
  "versionscount": 3
}
`,
	})

	// Update/PATCH - change defverid+sticky.
	// Ext should remain
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1/meta",
		Method:  "PATCH",
		ReqBody: `{"defaultversionid":"v1.0","defaultversionsticky":true}`,
		Code:    200,
		ResBody: `{
  "fileid": "f1",
  "self": "http://localhost:8181/dirs/d1/files/f1/meta",
  "xid": "/dirs/d1/files/f1/meta",
  "epoch": 9,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",
  "foo": "bar",

  "defaultversionid": "v1.0",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/v1.0",
  "defaultversionsticky": true
}
`,
	})

	// meta's epoch changed but the defver didn't
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f1?inline=meta",
		Method: "GET",
		Code:   200,
		ResBody: `{
  "fileid": "f1",
  "versionid": "v1.0",
  "self": "http://localhost:8181/dirs/d1/files/f1",
  "xid": "/dirs/d1/files/f1",
  "epoch": 1,
  "isdefault": true,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:01Z",

  "metaurl": "http://localhost:8181/dirs/d1/files/f1/meta",
  "meta": {
    "fileid": "f1",
    "self": "http://localhost:8181/dirs/d1/files/f1/meta",
    "xid": "/dirs/d1/files/f1/meta",
    "epoch": 9,
    "createdat": "2024-01-01T12:00:01Z",
    "modifiedat": "2024-01-01T12:00:02Z",
    "foo": "bar",

    "defaultversionid": "v1.0",
    "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/v1.0",
    "defaultversionsticky": true
  },
  "versionsurl": "http://localhost:8181/dirs/d1/files/f1/versions",
  "versionscount": 3
}
`,
	})

	// Update/PATCH - unstick
	// Ext should remain
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1/meta",
		Method:  "PATCH",
		ReqBody: `{"defaultversionsticky":null}`,
		Code:    200,
		ResBody: `{
  "fileid": "f1",
  "self": "http://localhost:8181/dirs/d1/files/f1/meta",
  "xid": "/dirs/d1/files/f1/meta",
  "epoch": 10,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",
  "foo": "bar",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/1"
}
`,
	})

	// Update/PATCH - stick
	// Ext should remain
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1/meta",
		Method:  "PATCH",
		ReqBody: `{"defaultversionsticky":true}`,
		Code:    200,
		ResBody: `{
  "fileid": "f1",
  "self": "http://localhost:8181/dirs/d1/files/f1/meta",
  "xid": "/dirs/d1/files/f1/meta",
  "epoch": 11,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",
  "foo": "bar",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/1",
  "defaultversionsticky": true
}
`,
	})

	// Update/PUT - empty - should erase ext and unstick it
	xCheckHTTP(t, reg, &HTTPTest{
		URL:    "/dirs/d1/files/f1/meta",
		Method: "PUT",
		Code:   200,
		ResBody: `{
  "fileid": "f1",
  "self": "http://localhost:8181/dirs/d1/files/f1/meta",
  "xid": "/dirs/d1/files/f1/meta",
  "epoch": 12,
  "createdat": "2024-01-01T12:00:01Z",
  "modifiedat": "2024-01-01T12:00:02Z",

  "defaultversionid": "1",
  "defaultversionurl": "http://localhost:8181/dirs/d1/files/f1/versions/1"
}
`,
	})

	// Update/PUT meta - bad epoch
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1/meta",
		Method:  "PUT",
		ReqBody: `{"epoch": 1}`,
		Code:    400,
		ResBody: `Attribute "epoch"(1) doesn't match existing value (12)
`,
	})

	// Update/PUT resource - bad epoch
	xCheckHTTP(t, reg, &HTTPTest{
		URL:     "/dirs/d1/files/f1",
		Method:  "PUT",
		ReqBody: `{"meta":{"epoch": 1}}`,
		Code:    400,
		ResBody: `Attribute "epoch"(1) doesn't match existing value (12)
`,
	})

}
