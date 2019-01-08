package flat

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFlattenMap(t *testing.T) {
	cases := []struct {
		test string
		want ResultMap
	}{
		{
			`{
				"number": 1.4567,
				"bool":   true
			}`,
			ResultMap{
				"number": 1.4567,
				"bool":   true,
			},
		},
		{
			`{
				"name": {
					"first_name": "John",
					"last_name": "Doe"
				}
			}`,
			ResultMap{
				"name.first_name": "John",
				"name.last_name":  "Doe",
			},
		},
		{
			`{
				"a": {
					"b": {
						"d": {
							"e": {
								"c": "c"
							}
						}
					},
					"c": "c"
				}
			}`,
			ResultMap{
				"a.b.d.e.c": "c",
				"a.c":       "c",
			},
		},
	}
	for i, test := range cases {
		var inmap ResultMap
		json.Unmarshal([]byte(test.test), &inmap)
		got := FlattenMap(inmap)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: mismatch, got: %v want: %v", i+1, got, test.want)
		}
	}
}

func TestFlattenJSON(t *testing.T) {
	cases := []struct {
		test string
		want string
	}{
		{
			`{
				"number": 1.4567,
				"bool":   true
			}`,
			`{
				"number": 1.4567,
				"bool":   true
			}`,
		},
		{
			`{
				"name": {
					"first_name": "John",
					"last_name": "Doe"
				}
			}`,
			`{
				"name.first_name": "John",
				"name.last_name": "Doe"
			}`,
		},
		{
			`{
				"a": {
					"b": {
						"d": {
							"e": {
								"c": "c"
							}
						}
					},
					"c": "c"
				}
			}`,
			`{
				"a.b.d.e.c": "c",
				"a.c":       "c"
			}`,
		},
	}
	for i, test := range cases {
		got, err := FlattenJSON(test.test)
		if err != nil {
			t.Errorf("%d: failed to unmarshal test: %v", i+1, err)
		}
		var inmap, outmap ResultMap
		json.Unmarshal([]byte(test.want), &inmap)
		json.Unmarshal([]byte(got), &outmap)
		if !reflect.DeepEqual(inmap, outmap) {
			t.Errorf("%d: mismatch, got: %v want: %v", i+1, inmap, outmap)
		}
	}
}

func TestUnflattenMap(t *testing.T) {
	cases := []struct {
		test string
		want ResultMap
	}{
		{
			`{
				"a.b.c": "c",
				"a.c":   "c"
			}`,
			ResultMap{
				"a": ResultMap{
					"b": ResultMap{
						"c": "c",
					},
					"c": "c",
				},
			},
		},
	}
	for i, test := range cases {
		var inmap ResultMap
		json.Unmarshal([]byte(test.test), &inmap)
		outmap := UnflattenMap(inmap)
		if !reflect.DeepEqual(outmap, test.want) {
			t.Errorf("%d: mismatch, got: %v want: %v", i+1, outmap, test.want)
		}
	}
}

func TestUnflattenJSON(t *testing.T) {
	cases := []struct {
		test string
		want string
	}{
		{
			`{
				"a.b.c": "c",
				"a.c":   "c"
			}`,
			`{
				"a": {
					"b": {
						"c": "c"
					},
					"c": "c"
				}
			}`,
		},
	}
	for i, test := range cases {
		got, err := UnflattenJSON(test.test)
		if err != nil {
			t.Errorf("%d: failed to unmarshal test: %v", i+1, err)
		}
		var inmap, outmap ResultMap
		json.Unmarshal([]byte(test.want), &inmap)
		json.Unmarshal([]byte(got), &outmap)
		if !reflect.DeepEqual(inmap, outmap) {
			t.Errorf("%d: mismatch, got: %v want: %v", i+1, outmap, inmap)
		}
	}
}
