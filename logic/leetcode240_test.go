package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {

	type testData struct {
		s   string
		t   string
		res bool
	}

	datas := []*testData{
		{s: "anagram", t: "nagaram", res: true},
		{s: "rat", t: "car", res: false},
	}
	for _, data := range datas {
		res := IsAnagram(data.s, data.t)
		assert.Equal(t, data.res, res)
	}

}
