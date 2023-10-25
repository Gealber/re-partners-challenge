package dimensions

import (
	"strconv"
	"strings"
	"sort"

	badger "github.com/dgraph-io/badger/v4"
)

const (
	defaultKey = "dimensions"
)

type repo struct {
	db *badger.DB
}

func New(db *badger.DB) *repo {
	return &repo{db: db}
}

func (r *repo) Get() ([]int, error) {
	dimensions := make([]int, 0)
	err := r.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(defaultKey))
		if err != nil {
			return err
		}

		valCopy, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}

		for _, valStr := range strings.Split(string(valCopy), ",") {
			val, _ := strconv.Atoi(valStr)
			dimensions = append(dimensions, val)
		}

		return nil
	})

	// sorting descending order
	sort.Slice(dimensions, func(i, j int) bool {
		return dimensions[i] > dimensions[j]
	})

	return dimensions, err
}

func (r *repo) Update(dimensions []int) error {
	dataStr := ""
	for _, val := range dimensions {
		dataStr += strconv.Itoa(val) + ","
	}

	dataStr = dataStr[:len(dataStr)-1]

	return r.db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(defaultKey), []byte(dataStr))
		err := txn.SetEntry(e)
		return err
	})
}
