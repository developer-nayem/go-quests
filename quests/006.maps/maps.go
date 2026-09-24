package maps

type Cache struct {
	data map[string]int
}

func NewCache() *Cache {
	// TODO: initialize the cache
	// Read README.md for the instructions

	return &Cache{
		data: make(map[string]int),
	}
}

func (c *Cache) Set(key string, value int) {
	// TODO: implement
	// Read README.md for the instructions
	c.data[key] = value
}

func (c *Cache) Get(key string) (int, bool) {
	// TODO: implement
	// Read README.md for the instructions
	item, exists := c.data[key]

	if exists {
		return item,exists
	}

	return 0, false
}

func (c *Cache) Delete(key string) {
	// TODO: implement
	// Read README.md for the instructions
	delete(c.data,key)
}

func (c *Cache) Count() int {
	// TODO: implement
	// Read README.md for the instructions
	return len(c.data)
}

func (c *Cache) AllKeys() []string {
	// TODO: implement
	// Read README.md for the instructions
	allKeys := make([]string, 0, len(c.data))
	for key := range c.data{
		allKeys = append(allKeys, key)
	}
	return allKeys
}

func (c *Cache) RemoveBelow(limit int) {
	// TODO: implement
	// Read README.md for the instructions
	for key := range c.data{
		if c.data[key] < limit {
			delete(c.data, key)
		}
	}
}
