//go:generate stringer -type=Pill

package painkiller

// Pill is pill
type Pill int

// Pill Name
const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
