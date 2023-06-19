package stage

import "github.com/quasilyte/sinecord/synthdb"

type instrument struct {
	fx    string
	oldFx string

	compiledFx func(x float64) float64

	instrumentIndex int
	patchNumber     int32

	period       float64
	enabled      bool
	mappedVolume int32
	volume       float64
	kind         synthdb.InstrumentKind
}

func (inst *instrument) SetFx(fx string) {
	inst.oldFx = inst.fx
	inst.fx = fx
}
