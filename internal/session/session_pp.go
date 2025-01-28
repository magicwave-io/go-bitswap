package session

import "time"

func (s *Session) FetchQueueCount() int {
	return len(s.sw.toFetch.elems)
}
func (s *Session) LiveWantsCount() int {
	return len(s.sw.liveWants)
}
func (s *Session) LiveWantsOldestAge() int {
	var t int = 0
	now := time.Now()
	for _, time := range s.sw.liveWants {
		v := int(time.Sub(now).Seconds())
		if v < t {
			t = v
		}
	}
	return t
}
func (s *Session) LiveWantsFirstAge() int {
	var t int = 0
	now := time.Now()
	for _, time := range s.sw.liveWants {
		v := int(time.Sub(now).Seconds())
		if v > t {
			t = v
		}
	}
	return t
}
