package shared

import (
	"crypto/md5"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tidwall/pretty"
)

func Check(e error) (msg string) {
	if e != nil {
		msg = e.Error()
	}
	return
}

// SetInt32 - int to *int32.
func SetInt32(i int) *int32 {
	r := int32(i)
	return &r
}

// SetString - string to *string.
func SetString(s string) *string {
	return &s
}

// SetBool - bool to *bool.
func SetBool(s bool) *bool {
	return &s
}

// SetStringFromPointer - interface{} to string.
func SetStringFromPointer(s interface{}) string {
	if s != nil {
		return fmt.Sprintf("%v", s)
	}
	return ""
}

// StringCompare - strings to lower then compare.
func StringCompare(a string, b string) bool {
	if strings.ToLower(a) == strings.ToLower(b) {
		return true
	}
	return false
}

// MarshalInterface converts a generic interface to a marshalled interface.
func MarshalInterface(in interface{}, out interface{}) (err error) {
	inB, err := json.Marshal(in)
	if err != nil {
		return
	}
	return json.Unmarshal(inB, &out)
}

// SortUniqueKeys sorts unique keys within a string array.
func SortUniqueKeys(in []string) ([]string, error) {
	var err error
	var resp []string
	respMap := make(map[string]string)
	for _, v := range in {
		respMap[v] = v
	}
	for _, v := range respMap {
		resp = append(resp, v)
	}
	sort.Strings(resp)
	return resp, err
}

// FormatAviRef parses a url path and returns the AVI UUID.
// Simplifies quering AVI routes using AVI SDK.
func FormatAviRef(in string) string {
	uriArr := strings.SplitAfter(in, "/")
	return uriArr[len(uriArr)-1]
}

// DecodeIO decodes io.Reader objects and maps the result to a pointer.
func DecodeIO(io io.Reader, model interface{}) (err error) {
	decoder := json.NewDecoder(io)
	err = decoder.Decode(&model)
	return
}

// GetMD5Hash returns the MD5 Hash of a string.
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// FromJSONFile parses a JSON file and stores its value in an interface.
func FromJSONFile(path string, ob interface{}) (err error) {
	toReturn := ob
	openedFile, err := os.Open(path)
	if err != nil {
		return
	}
	byteValue, err := ioutil.ReadAll(openedFile)
	if err != nil {
		return
	}
	err = json.Unmarshal(byteValue, &toReturn)
	if err != nil {
		return
	}
	defer openedFile.Close()
	return
}

// FromFile opens a file and returns is contents in bytes.
func FromFile(path string) (r []byte, err error) {
	openedFile, err := os.Open(path)
	if err != nil {
		return
	}
	defer openedFile.Close()
	r, err = ioutil.ReadAll(openedFile)
	if err != nil {
		return
	}
	return
}

// ToJSON converts interface to string.
func ToJSON(p interface{}) (r string) {
	bytes, err := json.Marshal(p)
	if err != nil {
		logrus.Warn(err)
	}
	return string(bytes)
}

// ToPrettyJSON converts interface to bytes.
func ToPrettyJSON(p interface{}) (r []byte) {
	bytes, err := json.Marshal(p)
	if err != nil {
		logrus.Warn(err)
	}
	return pretty.Pretty(bytes)
}

// MapInterface creates a map of an interface's top level keys.
// Values are represented as type interface and can be converted.
// to string using fmt.Sprint.
func MapInterface(i interface{}) (r map[string]interface{}, err error) {
	inrec, _ := json.Marshal(i)
	err = json.Unmarshal(inrec, &r)
	return
}

// RandStringBytesMaskImpr generates a random string. This function is a derivative pulled from: https://github.com/kpbird/golang_random_string/blob/master/main.go
func RandStringBytesMaskImpr(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// EncodePorts ...
func EncodePorts(in interface{}) (r string, err error) {
	var p []Port
	var ports []string

	err = MarshalInterface(in, &p)
	if err != nil {
		return
	}

	for _, v := range p {
		ports = append(ports, strconv.Itoa(v.Port))
	}
	portsSorted, _ := SortUniqueKeys(ports)
	portsStr := strings.Join(portsSorted, ",")
	portsEnc := b64.StdEncoding.EncodeToString([]byte(portsStr))
	return portsEnc, nil
}

func CheckName(name string) bool {
	regEx := regexp.MustCompile(`^prd[0-9]*-.+?-[a-zA-Z][a-zA-Z][a-zA-Z]$`)
	return regEx.MatchString(name)
}

func FetchPrdCode(name string) (r int) {
	regEx := regexp.MustCompile(`^prd[0-9]*`)
	matches := regEx.FindStringSubmatch(name)
	if len(matches) == 0 {
		return 1234
	}
	match := matches[0]
	r, err := strconv.Atoi(strings.TrimSpace(strings.ReplaceAll(match, "prd", "")))
	if err != nil {
		return 1234
	}
	return r
}

func SetName(code int, name string) (r string) {
	////////////////////////////////////////////////////////////////////////////
	fetched := FetchPrdCode(name)
	////////////////////////////////////////////////////////////////////////////
	if fetched != 1234 {
		if CheckName(name) {
			return name
		}
		r = fmt.Sprintf("%s-%v", strings.ToLower(name), RandStringBytesMaskImpr(3))
		return r
	}
	////////////////////////////////////////////////////////////////////////////
	r = fmt.Sprintf("prd%v-%s-%v", code, strings.ToLower(name), RandStringBytesMaskImpr(3))
	return r
}
