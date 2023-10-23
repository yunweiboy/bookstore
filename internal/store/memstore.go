package store

import (
	mystore "bookstore/store"
	"bookstore/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}

func (ms *MemStore) Create(book *mystore.Book) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.books[book.Id]; ok {
		return mystore.ErrExist
	}

	nBook := *book
	ms.books[book.Id] = &nBook

	return nil
}

func (ms *MemStore) Update(book *mystore.Book) error {
	ms.Lock()
	defer ms.Unlock()

	oldBook, ok := ms.books[book.Id]
	if !ok {
		return mystore.ErrNotFound
	}

	nBook := *oldBook
	if book.Name != "" {
		nBook.Name = book.Name
	}

	if book.Authors != nil {
		nBook.Authors = book.Authors
	}

	if book.Press != "" {
		nBook.Press = book.Press
	}

	ms.books[book.Id] = &nBook

	return nil
}

func (ms *MemStore) Get(id string) (mystore.Book, error) {
	ms.RLock() // 读锁
	defer ms.RUnlock()

	t, ok := ms.books[id]
	if ok {
		return *t, nil // 获取到某个图书条目
	}
	return mystore.Book{}, mystore.ErrNotFound
}

func (ms *MemStore) Delete(id string) error {
	ms.Lock() // 写锁
	defer ms.Unlock()

	if _, ok := ms.books[id]; ok {
		return mystore.ErrNotFound
	}

	delete(ms.books, id)
	return nil
}

func (ms *MemStore) GetAll() ([]mystore.Book, error) {
	ms.RLock()
	defer ms.RUnlock()

	allBooks := make([]mystore.Book, 0, len(ms.books))
	for _, book := range ms.books {
		allBooks = append(allBooks, *book)
	}

	return allBooks, nil
}
