package radix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRadixTreeInsert(t *testing.T) {
	var testCases = []struct {
		name  string
		pairs map[string]int
	}{
		{
			name: "Simple Insertion and Search",
			pairs: map[string]int{
				"apple":  1,
				"banana": 2,
			},
		},
		{
			name: "Prefix Sharing",
			pairs: map[string]int{
				"app":   1,
				"apple": 2,
				"apps":  3,
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			rt := NewRadixTree[int]()
			for k, v := range tt.pairs {
				rt.Insert(k, v)
			}
			for k, expectV := range tt.pairs {
				realV, ok := rt.Search(k)
				assert.True(t, ok, "Key %s not found", k)
				assert.Equal(t, expectV, realV, "Value for key %s mismatch", k)
			}
		})
	}
}



func TestRadixTreeDelete(t *testing.T) {
	var testCases = []struct {
		name  string
		deletetions []string
		insertions map[string]int
	}{
      {
      	name:        "Simple deleteion",
        insertions:  map[string]int{ "rust": 1, "go": 2, "java": 3, "c": 4 },
      	deletetions: []string{"java"},
      },
      {
      	name:        "Deletion with shared prefix",
        insertions:  map[string]int{ "bro": 1, "brother": 2, "brotherhood": 3, "brotherbrotherbrother": 4},
      	deletetions: []string{"brother"},
      },
      {
      	name:        "Delete root",
        insertions:  map[string]int{"he": 1, "hell": 2, "hello": 3, "helloworld": 4},
      	deletetions: []string{"he"},
      },
      {
      	name:        "Delete by not existing key",
        insertions:  map[string]int{ "rust": 1, "go": 2, "java": 3, "c": 4 },
      	deletetions: []string{"zig"},
      },
    }
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			rt := NewRadixTree[int]()
			for k, v := range tt.insertions {
				rt.Insert(k, v)
			}
			for _, delK := range tt.deletetions {
                rt.Delete(delK)
			}
            for _, d := range tt.deletetions {
                delete(tt.insertions, d) 
            }

			for k, expectV := range tt.insertions {
				realV, ok := rt.Search(k)
				assert.True(t, ok, "Key %s not found", k)
				assert.Equal(t, expectV, realV, "Value for key %s mismatch", k)
			}
            for _, d := range tt.deletetions {
				_, ok := rt.Search(d)
				assert.False(t, ok, "Key %s was't deleted", d)
            }
		})
	}
}

// Test cases for Radix Tree
var testCases = []struct {
	name    string
	actions []func(*RadixTree[int])
	expect  map[string]int
}{
	{
		name: "Simple Insertion and Search",
		actions: []func(*RadixTree[int]){
			func(rt *RadixTree[int]) { rt.Insert("apple", 1) },
			func(rt *RadixTree[int]) { rt.Insert("banana", 2) },
		},
		expect: map[string]int{
			"apple":  1,
			"banana": 2,
		},
	},
	{
		name: "Prefix Sharing",
		actions: []func(*RadixTree[int]){
			func(rt *RadixTree[int]) { rt.Insert("app", 1) },
			func(rt *RadixTree[int]) { rt.Insert("apple", 2) },
			func(rt *RadixTree[int]) { rt.Insert("apps", 3) },
		},
		expect: map[string]int{
			"app":   1,
			"apple": 2,
			"apps":  3,
		},
	},
	{
	},
}

