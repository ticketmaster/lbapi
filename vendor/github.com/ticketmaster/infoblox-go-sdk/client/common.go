package client

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/tidwall/pretty"
)

// MarshalInterface converts a generic interface to a pointer.
func MarshalInterface(in interface{}, out interface{}) (err error) {
	inB, err := json.Marshal(in)
	if err != nil {
		log.Print(err)
		return
	}
	return json.Unmarshal(inB, &out)
}

// OrderUnique sorts unique keys within a string array.
func OrderUnique(in []string) ([]string, error) {
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

// FromJSONFile converts a file object to a pointer.
func FromJSONFile(path string, ob interface{}) (err error) {
	toReturn := ob
	openedFile, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}
	byteValue, err := ioutil.ReadAll(openedFile)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(byteValue, &toReturn)
	if err != nil {
		log.Println(err)
		return
	}
	defer openedFile.Close()
	return
}

// FromFile returns the contents of a file.
func FromFile(path string) (r []byte, err error) {
	openedFile, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}
	defer openedFile.Close()
	r, err = ioutil.ReadAll(openedFile)
	if err != nil {
		log.Println(err)
	}
	return
}

// ToJSON converts interface to string.
func ToJSON(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		log.Println(err.Error())
	}
	return string(bytes)
}

// ToPrettyJSON converts interface to bytes.
func ToPrettyJSON(p interface{}) []byte {
	bytes, err := json.Marshal(p)
	if err != nil {
		log.Println(err.Error())
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
