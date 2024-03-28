package speedcontrol

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"testing"

	"github.com/fujiwara/shapeio"
)

func TestSpeedCalculator(t *testing.T) {
	src := bytes.NewReader(bytes.Repeat([]byte{0}, 2*32*1024)) // 32KB
	f, _ := os.Create("./foo")
	writer := shapeio.NewWriter(f)
	writer.SetRateLimit(1024 * 10)
	speedCalculator := New(writer)
	io.Copy(speedCalculator, src)
	f.Close()
	speed := speedCalculator.Speed()
	t.Log(fmt.Sprintf("%v", speed))
}
func TestSpeedPercentCalculator(t *testing.T) {
	f, _ := os.Create("./foo")
	writer := shapeio.NewWriter(f)
	writer.SetRateLimit(1024 * 10)
	speedCalculator := New(writer)
	percentage := NewSpeedControl(speedCalculator)
	mainSpeedCalculator := New(percentage)
	for i := 0; i < 16; i++ {
		src := bytes.NewReader(bytes.Repeat([]byte{0}, 5*1024)) // 32KB
		io.Copy(mainSpeedCalculator, src)
	}
	speedBefore := mainSpeedCalculator.Speed()
	t.Log(fmt.Sprintf("speedBefore  %v", speedBefore))
	percentage.SetPercentage(50)
	mainSpeedCalculator.Reset()
	for i := 0; i < 32; i++ {
		src := bytes.NewReader(bytes.Repeat([]byte{0}, 1024)) // 32KB
		io.Copy(mainSpeedCalculator, src)
	}

	speedAfter := mainSpeedCalculator.Speed()
	t.Log(fmt.Sprintf("speedAfter %v", speedAfter))
	f.Close()
	k := speedBefore / speedAfter
	if math.Abs(k-2) > 0.1 {
		t.Error("Speed Limitter not valid")
	}
}
