package util

func TestFunc(param string) (aCount, length int) {
	_aCount := 0
	for _, c := range param {
		if c == 'a' {
			_aCount = _aCount + 1
		}
	}
	return len(param), _aCount
}
