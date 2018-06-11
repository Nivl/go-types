// Package octets contains consts and methods to simplify dealing with
// file sizes
package octets

// List of all different supported type
const (
	Byte int64 = 1.0 << (10 * iota)
	KiloByte
	MegaByte
	GigaByte
	TeraByte
	PetaByte
	ExaByte

	B  int64 = Byte
	KB int64 = KiloByte
	MB int64 = MegaByte
	GB int64 = GigaByte
	TB int64 = TeraByte
	PB int64 = PetaByte
	EB int64 = ExaByte

	Octet     int64 = Byte
	KiloOctet int64 = KiloByte
	MegaOctet int64 = MegaByte
	GigaOctet int64 = GigaByte
	TeraOctet int64 = GigaByte
	PetaOctet int64 = PetaByte
	ExaOctet  int64 = ExaByte

	O  int64 = Octet
	Ko int64 = KiloOctet
	Mo int64 = MegaOctet
	Go int64 = GigaOctet
	To int64 = TeraOctet
	Po int64 = PetaOctet
	Eo int64 = ExaOctet
)
