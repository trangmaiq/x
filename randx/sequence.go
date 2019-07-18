package randx

import (
	"crypto/rand"
	"math/big"
)

var rander = rand.Reader

func RuneSequence(l int, allowedRunes []rune) (seq []rune, err error) {
	c := big.NewInt(int64(len(allowedRunes)))
	seq = make([]rune, l)

	for i := 0; i < l; i++ {
		r, err := rand.Int(rander, c)
		if err != nil {
			return seq, err
		}

		rn := allowedRunes[r.Uint64()]
		seq[i] = rn
	}

	return seq, nil
}
