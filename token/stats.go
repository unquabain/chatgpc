package token

import (
	"io"
	"math/rand"
	"strings"
)

type Stats map[string]map[string]int

func NewStats() Stats {
	return make(Stats)
}

func (ss Stats) Record(ring *Ring) {
	segments := ring.Segments()
	last := len(segments) - 1
	prefix := segments[:last]
	head := segments[last]

	for {
		last_stats, ok := ss[strings.Join(prefix, ``)]
		if !ok {
			last_stats = make(map[string]int)
		}
		last_stats[head]++
		ss[strings.Join(prefix, ``)] = last_stats
		if len(prefix) == 0 {
			break
		}
		prefix = prefix[1:]
	}
}

func (ss Stats) Read(n int, source io.Reader) error {
	ring := NewRing(n)
	reader := NewReader(source)
	for reader.Scan() {
		ring.Push(reader.Text())
		ss.Record(ring)
	}
	return reader.Err()
}

func (ss Stats) Next(ring *Ring) string {
	var (
		last_stats map[string]int
		ok         bool
		winner     string
		sum        int
		prefix     []string
	)
	prefix = ring.Segments()
	for {
		last_stats, ok = ss[strings.Join(prefix, ``)]
		if ok || len(prefix) == 0 {
			break
		}
		prefix = prefix[1:]
	}

	for candidate, weight := range last_stats {
		sum += weight
		if rand.Intn(sum) < weight {
			winner = candidate
		}
	}
	return winner
}

func (ss Stats) Iterator(n int) *StatsIterator {
	return &StatsIterator{stats: ss, ring: NewRing(n - 1)}
}

type StatsIterator struct {
	stats Stats
	ring  *Ring
	text  string
	err   error
}

func (si *StatsIterator) Scan() bool {
	if si.err != nil {
		return false
	}

	si.text = si.stats.Next(si.ring)
	if si.text == `` {
		return false
	}
	si.ring.Push(si.text)
	return true
}

func (si *StatsIterator) Text() string {
	return si.text
}

func (si *StatsIterator) Err() error {
	return si.err
}
