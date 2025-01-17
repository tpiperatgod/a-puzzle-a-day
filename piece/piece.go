package piece

import (
	"fmt"
	"strings"

	"github.com/gonejack/a-puzzle-a-day/board"
)

var pieces = []piece{
	{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
	},
	{
		{1, 0, 0},
		{1, 1, 1},
		{0, 0, 1},
	},
	{
		{0, 0, 1, 1},
		{1, 1, 1, 0},
	},
	{
		{1, 1, 1, 1},
		{0, 0, 1, 0},
	},
	{
		{0, 1, 1},
		{1, 1, 1},
	},
	{
		{1, 1, 1},
		{1, 0, 1},
	},
	{
		{1, 1, 1},
		{1, 1, 1},
	},
	{
		{1, 0, 0},
		{1, 0, 0},
		{1, 1, 1},
	},
}

var Pieces [][]piece

func init() {
	for _, p := range pieces {
		pt := p.transforms()
		Pieces = append(Pieces, pt)
	}
}

const n = 4

type piece [n][n]int

func (p piece) Print() {
	for r := range p {
		fmt.Println(p[r])
	}
	fmt.Println(strings.Repeat("-", 10))
}
func (p piece) rotate() (rp piece) {
	for r := range rp {
		for c := range rp[r] {
			rp[r][c] = p[c][n-r-1]
		}
	}
	return
}
func (p piece) flip() (f piece) {
	for r := range f {
		for c := 0; c < n; c++ {
			f[r][c] = p[r][n-c-1]
		}
	}
	return
}
func (p piece) transforms() (ps []piece) {
	ps = append(ps, p)

	for i := 0; i < 3; i++ {
		ps = append(ps, ps[i].rotate())
	}
	for i := 4; i < 8; i++ {
		ps = append(ps, ps[i-4].flip())
	}

	m := make(map[piece]struct{})
	for _, p := range ps {
		m[p.shift()] = struct{}{}
	}

	ps = nil
	for p := range m {
		ps = append(ps, p)
	}

	return
}
func (p piece) shift() (s piece) {
	si := 0
	for _, r := range p {
		if r[0]+r[1]+r[2]+r[3] == 0 {
			continue
		}
		s[si] = r
		si += 1
	}

	for {
		for r := range s {
			if s[r][0] != 0 {
				return
			}
		}

		for r := range s {
			for c := range s[r] {
				if c < 3 {
					s[r][c] = s[r][c+1]
				} else {
					s[r][c] = 0
				}
			}
		}
	}
}
func (p piece) CanPlace(b *board.Board7x7, row, col int) bool {
	return p.put(b, row, col, "", false)
}
func (p piece) Place(b *board.Board7x7, row, col int, text string) bool {
	return p.put(b, row, col, text, true)
}
func (p piece) put(b *board.Board7x7, row int, col int, text string, doWrite bool) (suc bool) {
	// find first not empty block
	dr, dc := 0, 0
	for dr = 0; dr < n; dr++ {
		for dc = 0; dc < n; dc++ {
			if p[dr][dc] == 1 {
				goto place
			}
		}
	}

place:
	for r := dr; r < n; r++ {
		for c := 0; c < n; c++ {
			if p[r][c] == 0 {
				continue
			}
			tr := row + r - dr
			tc := col + c - dc
			ok := b.CanSet(tr, tc)
			if !ok {
				return false
			}
			if doWrite {
				b.Set(text, tr, tc)
			}
		}
	}

	return true
}
