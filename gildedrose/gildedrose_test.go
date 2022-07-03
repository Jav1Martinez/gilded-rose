package gildedrose_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Gildedrose(t *testing.T) {

	t.Run("Update quality execution should decrease items SellIn value", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"foo", 10, 10},
		}
		gildedrose.UpdateQuality(items)
		assert.Equal(t, 9, items[0].SellIn)
	})
}
