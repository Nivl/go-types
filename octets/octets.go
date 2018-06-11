// Package octets contains consts and methods to simplify dealing with
// file sizes
package octets

// List of all different supported type
const (
	Byte = 1.0 << (10 * iota)
	KiloByte
	MegaByte
	GigaByte
	TeraByte
	PetaByte
	ExaByte

	B  = Byte
	KB = KiloByte
	MB = MegaByte
	GB = GigaByte
	TB = TeraByte
	PB = PetaByte
	EB = ExaByte

	Octet     = Byte
	KiloOctet = KiloByte
	MegaOctet = MegaByte
	GigaOctet = GigaByte
	TeraOctet = GigaByte
	PetaOctet = PetaByte
	ExaOctet  = ExaByte

	O  = Octet
	Ko = KiloOctet
	Mo = MegaOctet
	Go = GigaOctet
	To = TeraOctet
	Po = PetaOctet
	Eo = ExaOctet
)
