package copy

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

type AuthorInfo struct {
	Name    string `json:name`
	Age     int    `json:age`
	Country int    `json:country`
}

type Book struct {
	Title  string     `json:title`
	Author AuthorInfo `json:author`
	Year   int        `json:year`
	Category []string          `json:category`
	Price    map[string]string `json:price`
}

func DeepCopyByGob(dst, src interface{}) error {
	var buffer bytes.Buffer
	err := gob.NewEncoder(&buffer).Encode(src)
	if err != nil {
		return err
	}
	err = gob.NewDecoder(&buffer).Decode(dst)
	if err != nil {
		return err
	}
	return nil
}

func DeepCopyByJson(src []Book,dst *[]Book) error{
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, dst)
	if err != nil {
		return err
	}
	return err
}

func DeepCopyByCustom(src []Book)[]Book {
	dst := make([]Book, len(src))
	for i, book := range src {
		tmpbook := Book{}
		tmpbook.Title = book.Title
		tmpbook.Year = book.Year
		tmpbook.Author = AuthorInfo{}
		tmpbook.Author.Name = book.Author.Name
		tmpbook.Author.Age = book.Author.Age
		tmpbook.Author.Country = book.Author.Country
		tmpbook.Category = make([]string, len(book.Category))
		for index, category := range book.Category {
			tmpbook.Category[index] = category
		}
		tmpbook.Price = make(map[string]string)
		for k, v := range book.Price {
			tmpbook.Price[k] = v
		}
		dst[i] = tmpbook
	}
	return dst
}
