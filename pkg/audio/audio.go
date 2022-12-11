package audio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cryptix/wav"
)

func AudioToWav(audio []byte) ([]int16, error) {
	r, err := wav.NewReader(bytes.NewReader(audio), int64(binary.Size(audio)))
	if err != nil {
		return nil, fmt.Errorf("creating new reader failed: %w", err)
	}
	var d []int16
	for {
		s, err := r.ReadSample()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("reading sample failed: %w", err)
		}
		d = append(d, int16(s))
	}
	return d, nil
}
