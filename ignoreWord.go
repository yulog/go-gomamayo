package gomamayo

import (
	"fmt"
	"io"
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
		return nil
	}

	// 削除
	err = db.Delete(c.NewQuery(collection).Where(c.Field("surface").Eq(word)))
	if err != nil {
		return err
	}

	fmt.Println("Remove ignore word:", word)
	return nil
}

// RemoveAllIgnoreWords はすべての除外ワードを削除する
func RemoveAllIgnoreWords() error {
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
		return nil
	}

	// すべて削除
	err = db.DropCollection(collection)
	if err != nil {
		return err
	}

	fmt.Println("Remove all ignore word")
	return nil
}

// ListIgnoreWords は除外ワードを一覧する
func ListIgnoreWords(w io.Writer) error {
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
		return nil
	}

	docs, err := db.FindAll(c.NewQuery(collection))
	if err != nil {
		return err
	}

	for _, doc := range docs {
		fmt.Fprintln(w, doc.Get("surface"))
	}
	return nil
}

// ImportIgnoreWords は除外ワードをインポートする
func ImportIgnoreWords(path string) error {
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

	if collectionExists {
		return fmt.Errorf("collection already exist. could not import")
	}

	if path == "" {
		path = collection + ".json"
	}

	// インポート
	err = db.ImportCollection(collection, path)
	if err != nil {
		return err
	}

	fmt.Println("Import ignore word")
	return nil
}

// ExportIgnoreWords は除外ワードをエクスポートする
func ExportIgnoreWords(path string) error {
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
		return fmt.Errorf("collection does not exist. could not export")
	}

	if path == "" {
		path = collection + ".json"
	}

	// エクスポート
	err = db.ExportCollection(collection, path)
	if err != nil {
		return err
	}

	fmt.Println("Export ignore word")
	return nil
}
