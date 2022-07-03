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

	t.Run("Update quality execution should decrease items Quality value", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"foo", 10, 10},
		}
		gildedrose.UpdateQuality(items)
		assert.Equal(t, 9, items[0].Quality)
	})

	t.Run("Update quality execution should decrease items Quality value 2x faster after sell date has passed", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"foo", 0, 10},
		}
		gildedrose.UpdateQuality(items)
		assert.Equal(t, 8, items[0].Quality)
	})

	t.Run("Update quality execution should not decrease items Quality if the current value is 0", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"foo", 0, 0},
		}
		gildedrose.UpdateQuality(items)
		assert.Equal(t, 0, items[0].Quality)
	})
}
