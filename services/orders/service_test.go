package orders_test

import (
	"testing"

	"github.com/Gealber/re-partners-challenge/services/orders"
)

var (
	defaultDimensions = []int{5000, 2000, 1000, 500, 250}
)


type testCase struct {
	name string
	itemsToTest []int
	dimensions []int
	expectedPacks []map[int]int
}

func Test_PackOrder(t *testing.T) {
	srv := orders.New()

	for _, tc := range genTcs() {
		t.Run(tc.name, func(t *testing.T) {
			for i, items := range tc.itemsToTest {
				packing := srv.PackOrder(items, tc.dimensions)
				expectedPack := tc.expectedPacks[i] 
				for k, v := range expectedPack {
					packingValue, ok := packing[k]
					if !ok {
						t.Fail()
					}

					if v != packingValue {
						t.Fail()
					}
				}
			}
		})
	}
}

func genTcs() []testCase {
	return []testCase{
		{
			name: "simple", 
			itemsToTest: []int{1, 250, 251, 501, 12001},
			dimensions: defaultDimensions,
			expectedPacks: []map[int]int{
				{250: 1}, // 1
				{250: 1}, // 250
				{500: 1}, // 251
				{500: 1, 250: 1}, // 501
				{5000: 2, 2000: 1, 250: 1}, // 12001
			},
		},
	}
}
