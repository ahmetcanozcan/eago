package lib

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestURL(t *testing.T) {
	_ = qt.New(t)
	t.Run("split", func(t *testing.T) {
		c := qt.New(t)
		urlString := "this/is/a/test"
		partsArr := [][]string{
			splitPath(urlString),
			splitPath("/" + urlString),
			splitPath(urlString + "/"),
			splitPath("/" + urlString + "/"),
		}
		expected := []string{"this", "is", "a", "test"}

		for _, parts := range partsArr {
			c.Assert(parts, qt.CmpEquals(), expected)
		}

	})

	t.Run("urlPart", func(t *testing.T) {
		c := qt.New(t)

		paramPart := newURLPart("_id")

		c.Assert(paramPart.content, qt.Equals, "id")
		c.Assert(paramPart.partType, qt.Equals, parameterURLPart)

		constPart := newURLPart("const")

		c.Assert(constPart.content, qt.Equals, "const")
		c.Assert(constPart.partType, qt.Equals, constantURLPart)

	})

	// Helper function for getting rid of redundancy
	genareteTestURLStrings := func(str string) []string {
		return []string{
			"/" + str,
			str + "/",
			"/" + str + "/",
		}
	}

	t.Run("Check", func(t *testing.T) {
		c := qt.New(t)
		constURLs := []string{
			"/this/is/a/test/",
			"/",
			"",
		}
		for _, constURL := range constURLs {
			constURLPath := NewURLPath(constURL)
			for _, v := range genareteTestURLStrings(constURL) {
				c.Assert(constURLPath.Check(v), qt.IsTrue)
			}
		}

		t.Run("ParamURLPart", func(t *testing.T) {
			c := qt.New(t)

			paramURL := "my/name/is/_name"
			paramURLPath := NewURLPath(paramURL)
			for _, val := range genareteTestURLStrings("my/name/is/ahmet") {
				c.Assert(paramURLPath.Check(val), qt.IsTrue)
			}

		})

	})

	t.Run("GetParam", func(t *testing.T) {
		c := qt.New(t)
		urlPath := NewURLPath("profile/_a/_b/_c")
		params := urlPath.GetURLParams("profile/1/2/3")
		c.Assert(params,
			qt.CmpEquals(),
			map[string]string{
				"a": "1",
				"b": "2",
				"c": "3",
			})
	})

	t.Run("GetQueryParam", func(t *testing.T) {
		c := qt.New(t)

		for _, testCase := range []struct {
			s string
			m map[string]interface{}
		}{
			{
				s: "name=HAL&test=HELLO",
				m: map[string]interface{}{
					"name": "HAL",
					"test": "HELLO",
				},
			},
			{
				s: "name=HAL&test",
				m: map[string]interface{}{
					"name": "HAL",
					"test": true,
				},
			},
		} {
			res := GetQueryValues(testCase.s)
			c.Assert(res, qt.CmpEquals(), testCase.m)
		}
	})
}
