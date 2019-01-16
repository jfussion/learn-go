package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Joseph", "Manila"},
			[]string{"Joseph", "Manila"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Joseph", 25},
			[]string{"Joseph"},
		},
		{
			"Nested fields",
			&Person{
				"Joseph",
				Profile{25, "Manila"},
			},
			[]string{"Joseph", "Manila"},
		},
		{
			"Slices",
			[]Profile{
				{33, "Manila"},
				{20, "Cebu"},
			},
			[]string{"Manila", "Cebu"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "Manila"},
				{20, "Cebu"},
			},
			[]string{"Manila", "Cebu"},
		},
		{
			"Maps",
			map[string]string{
				"foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	var contains bool
	for _, x := range haystack {
		if needle == x {
			contains = true
			break
		}
	}
	if !contains {
		t.Errorf("expected %v to contain '%s' but it didnt", haystack, needle)
	}
}
