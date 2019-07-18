package lengthconv

import "fmt"

type Meter float64
type Foot float64

func (m Meter) String() string {
	return fmt.Sprintf("%g m", m)
}

func (f Foot) String() string {
	return fmt.Sprintf("%g ft.", f)
}
