package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	gourl "net/url"
	"os"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"testing"

	log "github.com/duglin/dlog"
	"github.com/xregistry/server/registry"
)

func TestMain(m *testing.M) {
	if tmp := os.Getenv("VERBOSE"); tmp != "" {
		if tmpInt, err := strconv.Atoi(tmp); err == nil {
			log.SetVerbose(tmpInt)
		}
	}

	// call flag.Parse() here if TestMain uses flags
	registry.DeleteDB("testreg")
	registry.CreateDB("testreg")
	registry.OpenDB("testreg")

	// DBName := "registry"
	// if !registry.DBExists(DBName) {
	// registry.CreateDB(DBName)
	// }
	// registry.OpenDB(DBName)

	// Start HTTP server

	server := registry.NewServer(8181).Start()

	// Run the tests
	rc := m.Run()

	// Shutdown HTTP server
	server.Close()

	if rc == 0 {
		// registry.DeleteDB("testreg")
	}
	os.Exit(rc)
}

func NewRegistry(name string, opts ...registry.RegOpt) *registry.Registry {
	var err error

	reg, _ := registry.FindRegistry(nil, name)
	if reg != nil {
		reg.Delete()
		reg.SaveAllAndCommit()
	}

	reg, err = registry.NewRegistry(nil, name, opts...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating registry %q: %s\n", name, err)
		ShowStack()
		os.Exit(1)
	}

	reg.SaveAllAndCommit()

	registry.DefaultRegDbSID = reg.DbSID

	/*
		// Now find it again and start a new Tx
		reg, err = registry.FindRegistry(nil, name)
		if err != nil {
			panic(err.Error())
		}
		if reg == nil {
			panic("nil")
		}
	*/

	return reg
}

func PassDeleteReg(t *testing.T, reg *registry.Registry) {
	if !t.Failed() {
		if os.Getenv("NO_DELETE_REGISTRY") == "" {
			// We do this to make sure that we can support more than
			// one registry in the DB at a time
			reg.Delete()

			/*
				rows, err := reg.Query("select * from Props")
				if err != nil || len(rows) != 0 {
					fmt.Printf("Rows: %s", ToJSON(rows))
					panic(fmt.Sprintf("Props left around: %s / %d", err, len(rows)))
				}
			*/
		}
		registry.DefaultRegDbSID = ""
	}
	reg.SaveAllAndCommit() // should this be Rollback() ?
}

func Caller() string {
	_, me, _, _ := runtime.Caller(0)

	for depth := 1; ; depth++ {
		_, file, line, ok := runtime.Caller(depth)
		if !ok {
			break
		}
		if file != me {
			return fmt.Sprintf("%s:%d", path.Base(file), line)
		}

	}
	return "unknownFile"
}

func Fail(t *testing.T, str string, args ...any) {
	t.Helper()
	text := strings.TrimSpace(fmt.Sprintf(str, args...))
	t.Fatalf("%s\n\n", text)
}

func xCheckErr(t *testing.T, err error, errStr string) {
	t.Helper()
	if err == nil {
		if errStr == "" {
			return
		}
		t.Fatalf("\nGot:<no err>\nExp: %s", errStr)
	}

	if errStr == "" {
		t.Fatalf("Test failed: %s", err)
	}

	if err.Error() != errStr {
		t.Fatalf("\nGot: %s\nExp: %s", err.Error(), errStr)
	}
}

func xCheck(t *testing.T, b bool, errStr string, args ...any) {
	t.Helper()
	if !b {
		t.Fatalf(errStr, args...)
	}
}

func ToJSON(obj interface{}) string {
	if obj != nil && reflect.TypeOf(obj).String() == "*errors.errorString" {
		obj = obj.(error).Error()
	}
	buf, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error Marshaling: %s", err)
	}
	return string(buf)
}

func xNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
}

func xCheckGet(t *testing.T, reg *registry.Registry, url string, expected string) {
	t.Helper()
	xNoErr(t, reg.SaveAllAndCommit())

	if len(url) > 0 {
		url = strings.TrimLeft(url, "/")
	}

	res, err := http.Get("http://localhost:8181/" + url)
	xNoErr(t, err)

	body, err := io.ReadAll(res.Body)
	buf := bytes.NewBuffer(body)
	daURL, _ := gourl.Parse(url)

	if daURL.Query().Has("noprops") {
		buf = bytes.NewBuffer(RemoveProps(buf.Bytes()))
		// expected = string(RemoveProps([]byte(expected)))
	}
	if daURL.Query().Has("oneline") {
		buf = bytes.NewBuffer(OneLine(buf.Bytes()))
		expected = string(OneLine([]byte(expected)))
	}

	xCheckEqual(t, "URL: "+url+"\n", buf.String(), expected)
}

func xCheckNotEqual(t *testing.T, extra string, gotAny any, expAny any) {
	t.Helper()

	exp := fmt.Sprintf("%v", expAny)
	got := fmt.Sprintf("%v", gotAny)

	if exp != got {
		return
	}

	t.Fatalf("Should differ, but they're both:\n%s", exp)
}

func xCheckGreater(t *testing.T, extra string, newAny any, oldAny any) {
	t.Helper()

	New := fmt.Sprintf("%v", newAny)
	Old := fmt.Sprintf("%v", oldAny)

	if New > Old {
		return
	}

	t.Fatalf("New not > Old:\nOld:\n%s\n\nNew:\n%s", Old, New)
}

var TSREGEXP = `\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[-+]\d{2}:\d{2})`
var TSMASK = TSREGEXP + `||YYYY-MM-DDTHH:MM:SSZ`

// Mask timestamps, but if (for the same input) the same TS is used, make sure
// the mask result is the same for just those two
func MaskTimestamps(input string) string {
	seenTS := map[string]string{}

	replaceFunc := func(input string) string {
		if val, ok := seenTS[input]; ok {
			return val
		}
		val := fmt.Sprintf("YYYY-MM-DDTHH:MM:%02dZ", len(seenTS)+1)
		seenTS[input] = val
		return val
	}

	re := savedREs[TSREGEXP]
	return re.ReplaceAllStringFunc(input, replaceFunc)
}

func xCheckEqual(t *testing.T, extra string, gotAny any, expAny any) {
	t.Helper()
	pos := 0

	exp := fmt.Sprintf("%v", expAny)
	got := fmt.Sprintf("%v", gotAny)

	// expected output starting with "--" means "skip timestamp masking"
	if len(exp) > 2 && exp[0:2] == "--" {
		exp = exp[2:]
	} else {
		got = MaskTimestamps(got)
		exp = MaskTimestamps(exp)
	}

	for pos < len(got) && pos < len(exp) && got[pos] == exp[pos] {
		pos++
	}
	if pos == len(got) && pos == len(exp) {
		return
	}

	if pos == len(got) {
		t.Fatalf("%s\n"+
			"Expected:\n%s\nGot:\n%s\nGot ended early at(%d)[%02X]:\n%q",
			extra, exp, got, pos, exp[pos], got[pos:])
	}

	if pos == len(exp) {
		t.Fatalf("%s\n"+
			"Expected:\n%s\nGot:\n%s\nExp ended early at(%d)[%02X]:\n%q",
			extra, exp, got, pos, got[pos], got[pos:])
	}

	expMax := pos + 90
	if expMax > len(exp) {
		expMax = len(exp)
	}

	t.Fatalf(extra+
		"\nExpected:\n"+exp+
		"\nGot:\n"+got+
		"\nDiff at(%d)[x%0x/x%0x]:"+
		"\nExp subset:\n"+exp[pos:expMax]+
		"\nGot:\n"+got[pos:],
		pos, exp[pos], got[pos])
}

type HTTPTest struct {
	Name       string
	URL        string
	Method     string
	ReqHeaders []string // name:value
	ReqBody    string

	Code        int
	HeaderMasks []string
	ResHeaders  []string // name:value
	BodyMasks   []string // "PROPNAME" or "SEARCH||REPLACE"
	ResBody     string
}

// http code, body
func xGET(t *testing.T, url string) (int, string) {
	t.Helper()
	url = "http://localhost:8181/" + url
	res, err := http.Get(url)
	if err != nil {
		t.Fatalf("HTTP GET error: %s", err)
	}

	body, _ := io.ReadAll(res.Body)
	/*
		if res.StatusCode != 200 {
			t.Logf("URL: %s", url)
			t.Logf("Code: %d\n%s", res.StatusCode, string(body))
		}
	*/

	return res.StatusCode, string(body)
}

func xHTTP(t *testing.T, reg *registry.Registry, verb, url, reqBody string, code int, resBody string) {
	t.Helper()
	xCheckHTTP(t, reg, &HTTPTest{
		URL:        url,
		Method:     verb,
		ReqBody:    reqBody,
		Code:       code,
		ResBody:    resBody,
		ResHeaders: []string{"*"},
	})
}

func xCheckHTTP(t *testing.T, reg *registry.Registry, test *HTTPTest) {
	t.Helper()
	xNoErr(t, reg.SaveAllAndCommit())

	// t.Logf("Test: %s", test.Name)
	// t.Logf(">> %s %s  (%s)", test.Method, test.URL, registry.GetStack()[1])

	if test.Name != "" {
		test.Name += ": "
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	body := io.Reader(nil)
	if test.ReqBody != "" {
		body = bytes.NewReader([]byte(test.ReqBody))
	}

	if len(test.URL) > 0 {
		test.URL = strings.TrimLeft(test.URL, "/")
	}

	req, err := http.NewRequest(test.Method,
		"http://localhost:8181/"+test.URL, body)
	xNoErr(t, err)

	// Add all request headers to the outbound message
	for _, header := range test.ReqHeaders {
		name, value, _ := strings.Cut(header, ":")
		name = strings.TrimSpace(name)
		value = strings.TrimSpace(value)
		req.Header.Add(name, value)
	}

	resBody := []byte{}
	res, err := client.Do(req)
	if res != nil {
		resBody, _ = io.ReadAll(res.Body)
	}

	xNoErr(t, err)
	xCheck(t, res.StatusCode == test.Code,
		fmt.Sprintf("Expected status %d, got %d\n%s",
			test.Code, res.StatusCode, string(resBody)))

	// t.Logf("%v\n%s", res.Header, string(resBody))
	testHeaders := map[string]string{}

	// This stuff is for masking timestamps. Need to make sure that we
	// process the expected and result timestamps in the same order, so
	// use 2 different "seenTS" maps
	testSeenTS := map[string]string{}
	resSeenTS := map[string]string{}
	replaceFunc := func(input string, seenTS map[string]string) string {
		if val, ok := seenTS[input]; ok {
			return val
		}
		val := fmt.Sprintf("YYYY-MM-DDTHH:MM:%02dZ", len(seenTS)+1)
		seenTS[input] = val
		return val
	}
	testReplaceFunc := func(input string) string {
		return replaceFunc(input, testSeenTS)
	}
	resReplaceFunc := func(input string) string {
		return replaceFunc(input, resSeenTS)
	}
	TSre := savedREs[TSREGEXP]

	// Parse expected headers - split and lowercase the name
	for _, v := range test.ResHeaders {
		name, value, _ := strings.Cut(v, ":")
		name = strings.ToLower(name)
		testHeaders[name] = strings.TrimSpace(value)
	}

	// Extract the response headers - lowercase the name.
	// Save the complete list for error reporting (gotHeaders)
	resHeaders := map[string]string{}
	gotHeaders := ""

	for name, vals := range res.Header {
		value := ""
		if len(vals) > 0 {
			value = vals[0]
		}

		name = strings.ToLower(name)
		resHeaders[name] = strings.TrimSpace(value)
		gotHeaders += fmt.Sprintf("\n%s: %s", name, value)
	}

	// Parse the headerMasks, if any so we can quickly use them later on
	headerMasks := []*regexp.Regexp{}
	headerReplace := []string{}
	for _, mask := range test.HeaderMasks {
		var re *regexp.Regexp
		search, replace, _ := strings.Cut(mask, "||")
		if re = savedREs[search]; re == nil {
			re = regexp.MustCompile(search)
			savedREs[search] = re
		}
		headerMasks = append(headerMasks, re)
		headerReplace = append(headerReplace, replace)
	}

	for name, value := range testHeaders {
		if name == "*" {
			continue
			// see comment in next section
		}

		// Make sure headers that start with '-' are NOT in the response
		if name[0] == '-' {
			if _, ok := resHeaders[name[1:]]; ok {
				t.Errorf("%sHeader '%s: %s' should not be "+
					"present\n\nGot headers:%s",
					test.Name, name[1:], value, gotHeaders)
				t.FailNow()
			}
			continue
		}

		resValue, ok := resHeaders[name]
		if !ok {
			t.Errorf("%s\nMissing header: %s: %s\n\nGot headers:%s",
				test.Name, name, value, gotHeaders)
			t.FailNow()
		}

		// Mask timestamps
		if strings.HasSuffix(name, "at") {
			value = TSre.ReplaceAllStringFunc(value, testReplaceFunc)
			resValue = TSre.ReplaceAllStringFunc(resValue, resReplaceFunc)
		}

		first := true // only mask the expected value once
		for i, re := range headerMasks {
			if first {
				value = re.ReplaceAllString(value, headerReplace[i])
				first = false
			}
			resValue = re.ReplaceAllString(resValue, headerReplace[i])
		}

		xCheckEqual(t, "Header:"+name+"\n", resValue, value)
		// Delete the response header so we'll know if there are any
		// unexpected xregistry- headers left around
		delete(resHeaders, name)
	}

	// Make sure we don't have any extra xReg headers
	// testHeaders with just "*":"" means skip all header checks
	// didn't use len(testHeaders) == 0 to ensure we don't skip by accident
	if len(testHeaders) != 1 || testHeaders["*"] != "" {
		for name, _ := range resHeaders {
			if !strings.HasPrefix(name, "xregistry-") {
				continue
			}
			t.Fatalf("%s\nExtra header(%s)\nGot:%s", test.Name, name, gotHeaders)
		}
	}

	// Only check body if not "*"
	if test.ResBody != "*" {
		testBody := test.ResBody

		for _, mask := range test.BodyMasks {
			var re *regexp.Regexp
			search, replace, found := strings.Cut(mask, "||")
			if !found {
				// Must be just a property name
				search = fmt.Sprintf(`("%s": ")(.*)(")`, search)
				replace = `${1}xxx${3}`
			}

			if re = savedREs[search]; re == nil {
				re = regexp.MustCompile(search)
				savedREs[search] = re
			}

			resBody = re.ReplaceAll(resBody, []byte(replace))
			testBody = re.ReplaceAllString(testBody, replace)
		}

		xCheckEqual(t, "Test: "+test.Name+"\nBody:\n",
			string(resBody), testBody)
		if t.Failed() {
			t.FailNow()
		}
	}
}

var savedREs = map[string]*regexp.Regexp{
	TSREGEXP: regexp.MustCompile(TSREGEXP),
}

func xJSONCheck(t *testing.T, gotObj any, expObj any) {
	t.Helper()
	got := ToJSON(gotObj)
	exp := ToJSON(expObj)
	xCheckEqual(t, "", got, exp)
}

func OneLine(buf []byte) []byte {
	buf = RemoveProps(buf)

	re := regexp.MustCompile(`[\r\n]*`)
	buf = re.ReplaceAll(buf, []byte(""))
	re = regexp.MustCompile(`([^a-zA-Z])\s+([^a-zA-Z])`)
	buf = re.ReplaceAll(buf, []byte(`$1$2`))
	re = regexp.MustCompile(`([^a-zA-Z])\s+([^a-zA-Z])`)
	buf = re.ReplaceAll(buf, []byte(`$1$2`))

	return buf
}

func RemoveProps(buf []byte) []byte {
	re := regexp.MustCompile(`\n[^{}]*\n`)
	buf = re.ReplaceAll(buf, []byte("\n"))

	re = regexp.MustCompile(`\s"labels": {\s*},*`)
	buf = re.ReplaceAll(buf, []byte(""))

	re = regexp.MustCompile(`\n *\n`)
	buf = re.ReplaceAll(buf, []byte("\n"))

	re = regexp.MustCompile(`\n *}\n`)
	buf = re.ReplaceAll(buf, []byte("}\n"))

	re = regexp.MustCompile(`}[\s,]+}`)
	buf = re.ReplaceAll(buf, []byte("}}"))
	buf = re.ReplaceAll(buf, []byte("}}"))

	return buf
}

func HTMLify(r *http.Request, buf []byte) []byte {
	str := fmt.Sprintf(`"(https?://%s[^"\n]*?)"`, r.Host)
	re := regexp.MustCompile(str)
	repl := fmt.Sprintf(`"<a href="$1?%s">$1?%s</a>"`,
		r.URL.RawQuery, r.URL.RawQuery)

	return re.ReplaceAll(buf, []byte(repl))
}

func NotNilString(val any) string {
	b := (val).([]byte)
	return string(b)
}

func NewPP() *registry.PropPath {
	return registry.NewPP()
}

func ShowStack() {
	fmt.Printf("-----\n")
	for i := 1; i < 20; i++ {
		pc, file, line, _ := runtime.Caller(i)
		fmt.Printf("Caller: %s:%d\n", path.Base(runtime.FuncForPC(pc).Name()), line)
		if strings.Contains(file, "main") || strings.Contains(file, "testing") {
			break
		}
	}
}

func IsNil(v any) bool {
	return registry.IsNil(v)
}
