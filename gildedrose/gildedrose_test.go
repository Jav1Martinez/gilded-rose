package gildedrose_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Item_SellIn_Gildedrose(t *testing.T) {

	t.Run("Update quality execution should decrease items SellIn value", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"foo", 10, 10},
		}
		gildedrose.UpdateQuality(items)
		assert.Equal(t, 9, items[0].SellIn)
	})
}

func Test_Item_Quality_Gildedrose(t *testing.T) {

	quality_tests := []struct {
		name     string
		args     []*gildedrose.Item
		expected int
	}{
		{
			name:     "Update quality execution should decrease items Quality value",
			args:     []*gildedrose.Item{{"foo", 10, 10}},
			expected: 9,
		},
		{
			name:     "Update quality execution should decrease items Quality value 2x faster after sell date has passed",
			args:     []*gildedrose.Item{{"foo", 0, 10}},
			expected: 8,
		},
		{
			name:     "Update quality execution should not decrease items Quality if the current value is 0",
			args:     []*gildedrose.Item{{"foo", 0, 0}},
			expected: 0,
		},
		{
			name:     "Update quality execution should increase Aged Brie quality in 1 point by day if SellIn is higher than 0",
			args:     []*gildedrose.Item{{"Aged Brie", 1, 10}},
			expected: 11,
		},
		{
			name:     "Update quality execution should increase Aged Brie quality in 2 point by day if SellIn is 0 or lower",
			args:     []*gildedrose.Item{{"Aged Brie", 0, 10}},
			expected: 12,
		},
		{
			name:     "Update quality execution should not increase Aged Brie item quality value beyond 50",
			args:     []*gildedrose.Item{{"Aged Brie", 10, 50}},
			expected: 50,
		},
		{
			name:     "Update quality execution should not increase Aged Brie item quality value beyond 50 even if sellIn value is 0 or lower",
			args:     []*gildedrose.Item{{"Aged Brie", 0, 50}},
			expected: 50,
		},
		{
			name:     "Update quality execution should decrease Conjurados items x2 faster as other regular item",
			args:     []*gildedrose.Item{{"Conjurados", 10, 10}},
			expected: 8,
		},
	}

	for _, test := range quality_tests {
		t.Run(test.name, func(t *testing.T) {
			gildedrose.UpdateQuality(test.args)
			assert.Equal(t, test.expected, test.args[0].Quality)
		})
	}
}

func Test_Item_SellIn_Quality_Gildadrose(t *testing.T) {

	t.Run("Update quality execution should not modify any value of Sulfuras items", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"Sulfuras, Hand of Ragnaros", 10, 80},
		}
		gildedrose.UpdateQuality(items)
		assert.Equal(t, 10, items[0].SellIn)
		assert.Equal(t, 80, items[0].Quality)
	})

	t.Run("Update quality execution should increase Backstage quality value x2 when sellIn is between 6 and 10", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"Backstage passes to a TAFKAL80ETC concert", 10, 10},
		}
		gildedrose.UpdateQuality(items)
		assert.Equal(t, 9, items[0].SellIn)
		assert.Equal(t, 12, items[0].Quality)
	})

	t.Run("Update quality execution should increase Backstage quality value x3 when sellIn is between 1 and 5", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"Backstage passes to a TAFKAL80ETC concert", 5, 10},
		}
		gildedrose.UpdateQuality(items)
		assert.Equal(t, 4, items[0].SellIn)
		assert.Equal(t, 13, items[0].Quality)
	})

	t.Run("Update quality execution should decrease Backstage quality value to 0 when sellIn is 0 or lower", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"Backstage passes to a TAFKAL80ETC concert", 0, 10},
		}
		gildedrose.UpdateQuality(items)
		assert.Equal(t, -1, items[0].SellIn)
		assert.Equal(t, 0, items[0].Quality)
	})
}
