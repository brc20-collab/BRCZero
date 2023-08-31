package types

import (
	"bytes"
	"fmt"
	"io"
	"sync"

	"github.com/tendermint/go-amino"

	"github.com/pkg/errors"

	"github.com/brc20-collab/brczero/libs/tendermint/crypto/merkle"
	"github.com/brc20-collab/brczero/libs/tendermint/libs/bits"
	tmbytes "github.com/brc20-collab/brczero/libs/tendermint/libs/bytes"
	tmmath "github.com/brc20-collab/brczero/libs/tendermint/libs/math"
	tmproto "github.com/brc20-collab/brczero/libs/tendermint/proto/types"
)

var (
	ErrPartSetUnexpectedIndex = errors.New("error part set unexpected index")
	ErrPartSetInvalidProof    = errors.New("error part set invalid proof")
)

type Part struct {
	Index int                `json:"index"`
	Bytes tmbytes.HexBytes   `json:"bytes"`
	Proof merkle.SimpleProof `json:"proof"`
}

func (part *Part) UnmarshalFromAmino(cdc *amino.Codec, data []byte) error {
	var dataLen uint64 = 0
	var subData []byte

	for {
		data = data[dataLen:]

		if len(data) == 0 {
			break
		}

		pos, aminoType, err := amino.ParseProtoPosAndTypeMustOneByte(data[0])
		if err != nil {
			return err
		}
		data = data[1:]

		if aminoType == amino.Typ3_ByteLength {
			var n int
			dataLen, n, err = amino.DecodeUvarint(data)
			if err != nil {
				return err
			}

			data = data[n:]
			if len(data) < int(dataLen) {
				return fmt.Errorf("not enough data for %s, need %d, have %d", aminoType, dataLen, len(data))
			}
			subData = data[:dataLen]
		}

		switch pos {
		case 1:
			uvint, n, err := amino.DecodeUvarint(data)
			if err != nil {
				return err
			}
			part.Index = int(uvint)
			dataLen = uint64(n)
		case 2:
			part.Bytes = make([]byte, dataLen)
			copy(part.Bytes, subData)
		case 3:
			err = part.Proof.UnmarshalFromAmino(cdc, subData)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpect feild num %d", pos)
		}
	}
	return nil
}

// ValidateBasic performs basic validation.
func (part *Part) ValidateBasic() error {
	if part.Index < 0 {
		return errors.New("negative Index")
	}
	if len(part.Bytes) > BlockPartSizeBytes {
		return errors.Errorf("too big: %d bytes, max: %d", len(part.Bytes), BlockPartSizeBytes)
	}
	if err := part.Proof.ValidateBasic(); err != nil {
		return errors.Wrap(err, "wrong Proof")
	}
	return nil
}

func (part *Part) String() string {
	return part.StringIndented("")
}

func (part *Part) StringIndented(indent string) string {
	return fmt.Sprintf(`Part{#%v
%s  Bytes: %X...
%s  Proof: %v
%s}`,
		part.Index,
		indent, tmbytes.Fingerprint(part.Bytes),
		indent, part.Proof.StringIndented(indent+"  "),
		indent)
}

//-------------------------------------

type PartSetHeader struct {
	Total int              `json:"total"`
	Hash  tmbytes.HexBytes `json:"hash"`
}

func (psh PartSetHeader) AminoSize() int {
	var size int
	if psh.Total != 0 {
		size += 1 + amino.UvarintSize(uint64(psh.Total))
	}
	if len(psh.Hash) != 0 {
		size += 1 + amino.UvarintSize(uint64(len(psh.Hash))) + len(psh.Hash)
	}
	return size
}

func (psh *PartSetHeader) UnmarshalFromAmino(_ *amino.Codec, data []byte) error {
	var dataLen uint64 = 0
	var subData []byte

	for {
		data = data[dataLen:]

		if len(data) == 0 {
			break
		}

		pos, aminoType, err := amino.ParseProtoPosAndTypeMustOneByte(data[0])
		if err != nil {
			return err
		}
		data = data[1:]

		if aminoType == amino.Typ3_ByteLength {
			var n int
			dataLen, n, err = amino.DecodeUvarint(data)
			if err != nil {
				return err
			}

			data = data[n:]
			if len(data) < int(dataLen) {
				return fmt.Errorf("not enough data for %s, need %d, have %d", aminoType, dataLen, len(data))
			}
			subData = data[:dataLen]
		}

		switch pos {
		case 1:
			var n int
			var uvint uint64
			uvint, n, err = amino.DecodeUvarint(data)
			if err != nil {
				return err
			}
			psh.Total = int(uvint)
			dataLen = uint64(n)
		case 2:
			psh.Hash = make([]byte, dataLen)
			copy(psh.Hash, subData)
		default:
			return fmt.Errorf("unexpect feild num %d", pos)
		}
	}
	return nil
}

func (psh PartSetHeader) String() string {
	return fmt.Sprintf("%v:%X", psh.Total, tmbytes.Fingerprint(psh.Hash))
}

func (psh PartSetHeader) IsZero() bool {
	return psh.Total == 0 && len(psh.Hash) == 0
}

func (psh PartSetHeader) Equals(other PartSetHeader) bool {
	return psh.Total == other.Total && bytes.Equal(psh.Hash, other.Hash)
}

// ValidateBasic performs basic validation.
func (psh PartSetHeader) ValidateBasic() error {
	if psh.Total < 0 {
		return errors.New("negative Total")
	}
	// Hash can be empty in case of POLBlockID.PartsHeader in Proposal.
	if err := ValidateHash(psh.Hash); err != nil {
		return errors.Wrap(err, "Wrong Hash")
	}
	return nil
}

// ToProto converts BloPartSetHeaderckID to protobuf
func (psh *PartSetHeader) ToProto() tmproto.PartSetHeader {
	if psh == nil {
		return tmproto.PartSetHeader{}
	}

	return tmproto.PartSetHeader{
		Total: int64(psh.Total),
		Hash:  psh.Hash,
	}
}

func (psh *PartSetHeader) ToIBCProto() tmproto.PartSetHeader {
	if psh == nil {
		return tmproto.PartSetHeader{}
	}
	return tmproto.PartSetHeader{
		Total: int64(psh.Total),
		Hash:  psh.Hash,
	}
}

// FromProto sets a protobuf PartSetHeader to the given pointer
func PartSetHeaderFromProto(ppsh *tmproto.PartSetHeader) (*PartSetHeader, error) {
	if ppsh == nil {
		return nil, errors.New("nil PartSetHeader")
	}
	psh := new(PartSetHeader)
	psh.Total = int(ppsh.Total)
	psh.Hash = ppsh.Hash

	return psh, psh.ValidateBasic()
}

//-------------------------------------

type PartSet struct {
	total int
	hash  []byte

	mtx           sync.Mutex
	parts         []*Part
	partsBitArray *bits.BitArray
	count         int
}

// Returns an immutable, full PartSet from the data bytes.
// The data bytes are split into "partSize" chunks, and merkle tree computed.
func NewPartSetFromData(data []byte, partSize int) *PartSet {
	// divide data into 4kb parts.
	total := (len(data) + partSize - 1) / partSize
	parts := make([]*Part, total)
	partsBytes := make([][]byte, total)
	partsBitArray := bits.NewBitArray(total)
	for i := 0; i < total; i++ {
		part := &Part{
			Index: i,
			Bytes: data[i*partSize : tmmath.MinInt(len(data), (i+1)*partSize)],
		}
		parts[i] = part
		partsBytes[i] = part.Bytes
		partsBitArray.SetIndex(i, true)
	}
	// Compute merkle proofs
	root, proofs := merkle.SimpleProofsFromByteSlices(partsBytes)
	for i := 0; i < total; i++ {
		parts[i].Proof = *proofs[i]
	}
	return &PartSet{
		total:         total,
		hash:          root,
		parts:         parts,
		partsBitArray: partsBitArray,
		count:         total,
	}
}

// Returns an empty PartSet ready to be populated.
func NewPartSetFromHeader(header PartSetHeader) *PartSet {
	return &PartSet{
		total:         header.Total,
		hash:          header.Hash,
		parts:         make([]*Part, header.Total),
		partsBitArray: bits.NewBitArray(header.Total),
		count:         0,
	}
}

func (ps *PartSet) Header() PartSetHeader {
	if ps == nil {
		return PartSetHeader{}
	}
	return PartSetHeader{
		Total: ps.total,
		Hash:  ps.hash,
	}
}

func (ps *PartSet) HasHeader(header PartSetHeader) bool {
	if ps == nil {
		return false
	}
	return ps.Header().Equals(header)
}

func (ps *PartSet) BitArray() *bits.BitArray {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return ps.partsBitArray.Copy()
}

func (ps *PartSet) Hash() []byte {
	if ps == nil {
		return nil
	}
	return ps.hash
}

func (ps *PartSet) HashesTo(hash []byte) bool {
	if ps == nil {
		return false
	}
	return bytes.Equal(ps.hash, hash)
}

func (ps *PartSet) Count() int {
	if ps == nil {
		return 0
	}
	return ps.count
}

func (ps *PartSet) Total() int {
	if ps == nil {
		return 0
	}
	return ps.total
}

func (ps *PartSet) AddPart(part *Part) (bool, error) {
	if ps == nil {
		return false, nil
	}
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	// Invalid part index
	if part.Index >= ps.total {
		return false, ErrPartSetUnexpectedIndex
	}

	// If part already exists, return false.
	if ps.parts[part.Index] != nil {
		return false, nil
	}

	// Check hash proof
	if part.Proof.Verify(ps.Hash(), part.Bytes) != nil {
		return false, ErrPartSetInvalidProof
	}

	// Add part
	ps.parts[part.Index] = part
	ps.partsBitArray.SetIndex(part.Index, true)
	ps.count++
	return true, nil
}

func (ps *PartSet) GetPart(index int) *Part {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return ps.parts[index]
}

func (ps *PartSet) IsComplete() bool {
	return ps.count == ps.total
}

func (ps *PartSet) GetReader() io.Reader {
	if !ps.IsComplete() {
		panic("Cannot GetReader() on incomplete PartSet")
	}
	return NewPartSetReader(ps.parts)
}

type PartSetReader struct {
	i      int
	parts  []*Part
	reader *bytes.Reader
}

func NewPartSetReader(parts []*Part) *PartSetReader {
	return &PartSetReader{
		i:      0,
		parts:  parts,
		reader: bytes.NewReader(parts[0].Bytes),
	}
}

func (psr *PartSetReader) Read(p []byte) (n int, err error) {
	readerLen := psr.reader.Len()
	if readerLen >= len(p) {
		return psr.reader.Read(p)
	} else if readerLen > 0 {
		n1, err := psr.Read(p[:readerLen])
		if err != nil {
			return n1, err
		}
		n2, err := psr.Read(p[readerLen:])
		return n1 + n2, err
	}

	psr.i++
	if psr.i >= len(psr.parts) {
		return 0, io.EOF
	}
	psr.reader = bytes.NewReader(psr.parts[psr.i].Bytes)
	return psr.Read(p)
}

func (ps *PartSet) StringShort() string {
	if ps == nil {
		return "nil-PartSet"
	}
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return fmt.Sprintf("(%v of %v)", ps.Count(), ps.Total())
}

func (ps *PartSet) MarshalJSON() ([]byte, error) {
	if ps == nil {
		return []byte("{}"), nil
	}

	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	return cdc.MarshalJSON(struct {
		CountTotal    string         `json:"count/total"`
		PartsBitArray *bits.BitArray `json:"parts_bit_array"`
	}{
		fmt.Sprintf("%d/%d", ps.Count(), ps.Total()),
		ps.partsBitArray,
	})
}