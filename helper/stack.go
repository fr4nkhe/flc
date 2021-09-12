package helper

type CountItem struct {
	Line int
}

type Stack struct {
	data []CountItem
}

func (s *Stack) Push(item CountItem) bool {
	s.data = append(s.data, item)
	return true
}

func (s *Stack) Pop() (CountItem, bool) {
	length := s.Length()
	var item CountItem
	if length == 0 {
		return item, false
	} else if length == 1 {
		item = s.data[0]
		s.data = []CountItem{}
		return item, true
	} else {
		item = s.data[0]
		s.data = s.data[1:]
		return item, true
	}
}

func (s *Stack) ElementAt(index int) (CountItem, bool) {
	length := len(s.data)
	if index >= length {
		return CountItem{}, false
	}
	return s.data[index], true
}

func (s *Stack) Length() int {
	return len(s.data)
}
