package models

import (
	"errors"
	"strconv"
)

type indexModel string

type Index struct {
  name string
  count int
}

const IndexModel indexModel = "IndexModel"

// Get data from the model
func (i *Index) Name() string {
  return i.name
}

func (i *Index) SetName(name string) {
  i.name = name
}

func (i *Index) AddCount(count int) {
  i.count += count
}

func (i *Index) GetCount() string {
  return strconv.FormatInt(int64(i.count), 10)
}

func (i indexModel) GetIndex(name string) ( Index, error) {
  count := 0
  if name == "" {
    return Index{}, errors.New("NO NAME GIVEN")
  }

  model := Index {
    name,
    count,
  }

  return model, nil
}
