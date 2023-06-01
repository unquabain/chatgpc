package token

import "github.com/Unquabain/chatgpc/generic"

type Ring struct {
	max  int
	data *generic.Ring[string]
}

func NewRing(n int) *Ring {
	return &Ring{
		max:  n,
		data: nil,
	}
}

func (r *Ring) Push(val string) {
	if r.data == nil {
		r.data = generic.NewRing[string](1)
	} else if r.data.Len() < r.max {
		r.data.Link(generic.NewRing[string](1))
	}
	r.data = r.data.Next()
	r.data.Set(val)
}

func (r *Ring) Segments() []string {
	segments := make([]string, 0, r.data.Len())
	r.data.Next().Do(func(seg string) {
		segments = append(segments, seg)
	})
	return segments
}
