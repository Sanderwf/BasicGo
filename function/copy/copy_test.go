package copy

import (
	"fmt"
	"testing"
)

func getBook() []Book {
	books := make([]Book, 1)
	author := AuthorInfo{"David", 38, 1156}
	price := make(map[string]string)
	price["Europe"] = "$56"
	books[0] = Book{"Tutorial", author, 2020, []string{"math", "art"}, price}
	return books
}

func TestDeepCopyByGob(t *testing.T) {
	books := getBook()
	fmt.Println(books)
	booksCpy := make([]Book, 1)
	err := DeepCopyByGob(&booksCpy, books)
	if err != nil {
		t.Log("\nerr: ", err)
	}
	t.Logf("\n类型： %T: \n 值：%v", booksCpy, booksCpy)
}

func TestDeepCopyByJson(t *testing.T) {
	books := getBook()
	fmt.Println(books)
	for _,book:= range books{
		book.Title = "!!!!!!!!!!!!!!!!"
	}
	fmt.Println(books)
	dst := make([]Book, 1)
	err := DeepCopyByJson(books, &dst)
	if err != nil {
		t.Log("\nerr: ", err)
	}
	t.Logf("\n类型： %T: \n 值：%v", dst, dst)
}

func TestDeepCopyByCustom(t *testing.T) {
	books := getBook()
	dst := DeepCopyByCustom(books)
	t.Log(dst)
}
