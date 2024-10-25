package speedcontrol

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/fujiwara/shapeio"
)

const avgCount = 50

type SpeedDetector struct {
	w           io.Writer
	locker      sync.Mutex
	dutations   []time.Duration
	byteLens    []int
	index       int
	saveTime    bool
	dataCount   int
	sumDuration time.Duration
}

func New(w io.Writer) *SpeedDetector {
	return &SpeedDetector{
		w:         w,
		dutations: make([]time.Duration, avgCount),
		byteLens:  make([]int, avgCount),
		locker:    sync.Mutex{},
		saveTime:  true,
	}
}
func (s *SpeedDetector) Speed() float64 {
	s.locker.Lock()
	defer s.locker.Unlock()
	s.saveTime = false
	sum := 0
	var timeSum time.Duration = 0
	for i := 0; i < avgCount; i++ {
		sum += s.byteLens[i]
		timeSum += s.dutations[i]
	}
	seconds := timeSum.Seconds()
	s.saveTime = true
	return float64(sum) / seconds
}
func (s *SpeedDetector) Reset() {
	s.locker.Lock()
	defer s.locker.Unlock()
	s.dutations = make([]time.Duration, avgCount)
	s.byteLens = make([]int, avgCount)
	s.index = 0
	s.dataCount = 0
	s.sumDuration = 0
}

func (s *SpeedDetector) Write(p []byte) (n int, err error) {
	if s.saveTime {
		s.locker.Lock()
		defer s.locker.Unlock()
		now := time.Now()
		//n, err = bufio.NewWriter(s.w).Write(p)
		n, err = s.w.Write(p)
		if n != len(p) {
			fmt.Println("")
		}
		d := time.Since(now)
		s.byteLens[s.index] = n
		s.sumDuration += s.dutations[s.index]
		s.dutations[s.index] = d
		s.index = (s.index + 1) % avgCount
		s.dataCount += 1
		return n, err
	} else {
		n, err = s.w.Write(p)
		return n, err
	}
}

//PERDENTAGE
type SpeedControlWriter struct {
	limiter       *shapeio.Writer
	speedDetector *SpeedDetector
	hasLimit      bool
}

func NewSpeedControl(w io.Writer) *SpeedControlWriter {
	speedDetector := New(w)
	limiter := shapeio.NewWriter(speedDetector)
	return &SpeedControlWriter{
		speedDetector: speedDetector,
		limiter:       limiter,
		hasLimit:      false,
	}
}
func (s *SpeedControlWriter) SetPercentage(percent int) {
	speed := s.speedDetector.Speed()
	s.hasLimit = percent != 100
	if speed != 0 {
		speedLimit := speed * (float64(percent) / float64(100))
		s.limiter.SetRateLimit(speedLimit)
	}
}

func (s *SpeedControlWriter) Write(p []byte) (n int, err error) {
	if s.hasLimit {
		return s.limiter.Write(p)
	} else {
		return s.speedDetector.Write(p)
	}
}
