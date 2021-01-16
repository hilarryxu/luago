package binchunk

import "github.com/hilarryxu/golua/types"

func Undump(data []byte) *types.Prototype {
	reader := &reader{data}
	reader.checkHeader()
	return nil
}
