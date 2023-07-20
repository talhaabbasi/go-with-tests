package maps

type Dictionary map[string]string

func (dictionary Dictionary) Search(keyword string) string {
	return dictionary[keyword]
}
