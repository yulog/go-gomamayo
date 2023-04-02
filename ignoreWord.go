package gomamayo

import (
	"errors"
	"fmt"
	"strings"

	c "github.com/ostafen/clover/v2"
)

const (
	collection   = "ignoreWords"
	ignoreWordDB = "ignoreWord-db"
)

type ignoreWord struct {
	Surface string `clover:"surface"`
}

// 除外ワードを適用する
func applyIgnoreWordsRemoval(input string) (string, error) {
	db, err := c.Open(ignoreWordDB)
	if err != nil {
		return input, err
	}
	defer db.Close()

	// collection の有無
	collectionExists, _ := db.HasCollection(collection)

	if !collectionExists {
		return input, nil
	}

	ignore := new(ignoreWord)

	db.ForEach(c.NewQuery(collection), func(doc *c.Document) bool {
		doc.Unmarshal(ignore)
		if strings.Contains(input, ignore.Surface) {
			input = strings.ReplaceAll(input, ignore.Surface, "")
			// fmt.Println("Ignore word:", ignore.Surface)
		}
		return true
	})

	return input, nil
}

// AddIgnoreWord は除外ワードを追加する
func AddIgnoreWord(word string) error {
	db, err := c.Open(ignoreWordDB)
	if err != nil {
		return err
	}
	defer db.Close()

	// collection の有無
	collectionExists, err := db.HasCollection(collection)
	if err != nil {
		return err
	}

	if !collectionExists {
		// なければ作る
		db.CreateCollection(collection)
	}

	// document を作る
	doc := c.NewDocument()
	doc.Set("surface", word)

	// collection に document を挿入
	docId, err := db.InsertOne(collection, doc)
	if err != nil {
		return err
	}

	fmt.Println("Add ignore word:", docId, word)
	return nil
}

// RemoveIgnoreWord は除外ワードを削除する
func RemoveIgnoreWord(word string) error {
	db, err := c.Open(ignoreWordDB)
	if err != nil {
		return err
	}
	defer db.Close()

	// collection の有無
	collectionExists, err := db.HasCollection(collection)
	if err != nil {
		return err
	}

	if !collectionExists {
		return errors.New("collection does not exist")
	}

	// 削除
	err = db.Delete(c.NewQuery(collection).Where(c.Field("surface").Eq(word)))
	if err != nil {
		return err
	}

	fmt.Println("Remove ignore word:", word)
	return nil
}
