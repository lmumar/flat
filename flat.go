package flat

import (
	"encoding/json"
	"strings"
)

// ResultMap map type returned/consumed by the library
type ResultMap map[string]interface{}

// FlattenJSON generates a new JSON string with keys flattened.
//
// Given the following JSON string:
//   { name: {
//      first_name: 'John',
//      last_name: 'Doe'
//    }}
//
// It generates the following JSON string:
//  {'name.first_name': 'John', 'name.last_name': 'Doe'}
func FlattenJSON(str string) (string, error) {
	var inmap ResultMap
	err := json.Unmarshal([]byte(str), &inmap)
	if err != nil {
		return "", err
	}
	outmap, err := FlattenMap(inmap)
	if err != nil {
		return "", err
	}
	b, err := json.Marshal(outmap)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// FlattenMap generates a new map with keys flattened.
//
// Given the following JSON string:
//   { name: {
//      first_name: 'John',
//      last_name: 'Doe',
//      experiences: [
//        {
//				year: {start: 2000, end: 2003},
//              position: 'programmer',
//              company: 'acme'
//		  }
//      ]
//    }}
//
// It generates the following map:
//  map[string]interface{} {
//    "name.first_name": "John",
//    "name.last_name": "Doe",
//    "experiences": []map[string]interface{} {
//	    map[string]interface{} {
// 	      "year.start": 2000,
//        "year.end": 2003,
//        "position": "programmer",
//        "company": "acme",
//      },
//    }
// }
func FlattenMap(inmap ResultMap) (ResultMap, error) {
	outmap := make(ResultMap)
	flatten("", inmap, outmap)
	return outmap, nil
}

func flatten(prefix string, inmap, outmap ResultMap) {
	for k, v := range inmap {
		nk := newkey(prefix, k)
		switch v.(type) {
		case map[string]interface{}:
			flatten(nk, v.(map[string]interface{}), outmap)
		default:
			outmap[nk] = v
		}
	}
}

func newkey(prefix, key string) string {
	if prefix == "" {
		return key
	}
	if key == "" {
		return prefix
	}
	return strings.Join([]string{prefix, key}, ".")
}

// UnflattenJSON unflattens a flat JSON string and returns it
// to the caller as a new flat JSON string.
func UnflattenJSON(j string) (string, error) {
	var inmap ResultMap
	err := json.Unmarshal([]byte(j), &inmap)
	if err != nil {
		return "", err
	}
	outmap := UnflattenMap(inmap)
	b, err := json.Marshal(outmap)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// UnflattenMap unflattens a flat map.
func UnflattenMap(inmap ResultMap) ResultMap {
	outmap := make(ResultMap)
	for k, v := range inmap {
		kparts := strings.Split(k, ".")
		unflatten(kparts, v, outmap)
	}
	return outmap
}

func unflatten(kparts []string, value interface{}, outmap ResultMap) {
	key := kparts[0]
	if len(kparts) > 1 {
		vmap, exists := outmap[key]
		if !exists {
			vmap = make(ResultMap)
		}
		unflatten(kparts[1:], value, vmap.(ResultMap))
		outmap[key] = vmap
	} else {
		outmap[key] = value
	}
}
