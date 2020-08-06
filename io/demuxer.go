package io

type Demuxer struct {
	inputs map[int]chan (PositionDTO)
}

const total = 10

func NewDemuxer() Demuxer {
	inputs := make(map[int]chan (PositionDTO))
	for i := 0; i < total; i++ {
		c := make(chan (PositionDTO))
		inputs[i] = c
	}

	return Demuxer{inputs: inputs}
}

func (d Demuxer) Demux(dto PositionDTO) {
	chanPosition := dto.RideID % 10

	d.inputs[chanPosition] <- dto
}

func (d Demuxer) close() {
	for _, c := range d.inputs {
		close(c)
	}
}
