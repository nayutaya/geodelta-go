package encoder

import (
	"strings"
)

type DeltaIds []byte
type DeltaCode string

var world_id_to_code map[byte]string = map[byte]string{
	0: "Z",
	1: "Y",
	2: "X",
	3: "W",
	4: "V",
	5: "T",
	6: "S",
	7: "R",
}

var world_code_to_id map[string]byte = map[string]byte{
	"Z": 0,
	"Y": 1,
	"X": 2,
	"W": 3,
	"V": 4,
	"T": 5,
	"S": 6,
	"R": 7,
}

func EncodeWorldDelta(id byte) string {
	if code, ok := world_id_to_code[id]; ok {
		return code
	} else {
		panic("invalid world delta id")
	}
}

func DecodeWorldDelta(code string) byte {
	if id, ok := world_code_to_id[code]; ok {
		return id
	} else {
		panic("invalid world delta code")
	}
}

var sub_ids2_to_code map[[2]byte]string = map[[2]byte]string{
	[2]byte{0, 0}: "2",
	[2]byte{0, 1}: "3",
	[2]byte{0, 2}: "4",
	[2]byte{0, 3}: "5",
	[2]byte{1, 0}: "6",
	[2]byte{1, 1}: "7",
	[2]byte{1, 2}: "8",
	[2]byte{1, 3}: "A",
	[2]byte{2, 0}: "B",
	[2]byte{2, 1}: "C",
	[2]byte{2, 2}: "D",
	[2]byte{2, 3}: "E",
	[2]byte{3, 0}: "F",
	[2]byte{3, 1}: "G",
	[2]byte{3, 2}: "H",
	[2]byte{3, 3}: "J",
}

var sub_ids1_to_code map[byte]string = map[byte]string{
	0: "K",
	1: "M",
	2: "N",
	3: "P",
}

var sub_code_to_ids map[string][]byte = map[string][]byte{
	"2": []byte{0, 0},
	"3": []byte{0, 1},
	"4": []byte{0, 2},
	"5": []byte{0, 3},
	"6": []byte{1, 0},
	"7": []byte{1, 1},
	"8": []byte{1, 2},
	"A": []byte{1, 3},
	"B": []byte{2, 0},
	"C": []byte{2, 1},
	"D": []byte{2, 2},
	"E": []byte{2, 3},
	"F": []byte{3, 0},
	"G": []byte{3, 1},
	"H": []byte{3, 2},
	"J": []byte{3, 3},
	"K": []byte{0},
	"M": []byte{1},
	"N": []byte{2},
	"P": []byte{3},
}

func EncodeSubDelta(ids []byte) string {
	if length := len(ids); length > 2 {
		return EncodeSubDelta(ids[0:2]) + EncodeSubDelta(ids[2:])
	} else if length == 2 {
		key := [2]byte{ids[0], ids[1]}
		if code, ok := sub_ids2_to_code[key]; ok {
			return code
		}
	} else if length == 1 {
		key := ids[0]
		if code, ok := sub_ids1_to_code[key]; ok {
			return code
		}
	}
	panic("invalid sub delta ids")
}

func DecodeSubDelta(codes string) []byte {
	return DecodeSubDeltaArray(strings.Split(codes, ""))
}

func DecodeSubDeltaArray(codes []string) []byte {
	if length := len(codes); length > 1 {
		return append(DecodeSubDeltaOne(codes[0]), DecodeSubDeltaArray(codes[1:])...)
	} else if length == 1 {
		return DecodeSubDeltaOne(codes[0])
	} else {
		panic("sub delta codes is empty")
	}
}

func DecodeSubDeltaOne(code string) []byte {
	if ids, ok := sub_code_to_ids[code]; ok {
		return ids
	} else {
		panic("invalid sub delta code")
	}
}

func (ids DeltaIds) Encode() DeltaCode {
	if length := len(ids); length > 1 {
		return DeltaCode(EncodeWorldDelta(ids[0]) + EncodeSubDelta(ids[1:]))
	} else if length == 1 {
		return DeltaCode(EncodeWorldDelta(ids[0]))
	} else {
		panic("delta ids is empty")
	}
}

func (code DeltaCode) Decode() DeltaIds {
	code_s := string(code)
	if length := len(code_s); length > 1 {
		return append([]byte{DecodeWorldDelta(code_s[0:1])}, DecodeSubDelta(code_s[1:])...)
	} else if length == 1 {
		return []byte{DecodeWorldDelta(code_s[0:1])}
	} else {
		panic("delta codes is empty")
	}
}
