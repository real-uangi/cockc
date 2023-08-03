// Package bar
// @author uangi
// @date 2023/8/1 15:06
package bar

import (
	"fmt"
	"sync"
	"time"
	"unicode/utf8"
)

type ProgressBar struct {
	percentage float64
	cur        float64
	total      float64
	graph      string
	stick      string
	start      time.Time
	mu         sync.Mutex
}

func NewProgressBar(total float64) *ProgressBar {
	return &ProgressBar{
		percentage: 0.00,
		cur:        0,
		total:      total,
		graph:      "█",
		stick:      "",
		start:      time.Now(),
	}
}

func (b *ProgressBar) refresh() {
	cost := time.Now().Sub(b.start).Seconds()
	eta := cost*100.0/b.percentage - cost
	fmt.Printf("\r %s [%-50s] %.2f%%  ETA%-8s  %.2f/%.2f", formatTime(int(cost)), b.stick, b.percentage, formatTime(int(eta)), b.cur, b.total)
}

func formatTime(sec int) string {
	return fmt.Sprintf("%02d:%02d:%02d", sec/3600, sec/60%60, sec%60)
}

func (b *ProgressBar) Add(step float64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	last := b.percentage
	b.cur += step
	b.percentage = b.cur / b.total * 100.0

	length := int(b.percentage / 2.0)
	for length > utf8.RuneCountInString(b.stick) {
		b.stick += b.graph
	}
	if last != b.percentage {
		b.refresh()
	}
	if b.IsFinish() {
		b.Finish()
	}
}

func (b *ProgressBar) Finish() {
	fmt.Println()
}

func (b *ProgressBar) IsFinish() bool {
	return b.total <= b.cur
}
