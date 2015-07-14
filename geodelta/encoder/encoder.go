package encoder

import (
	"strings"
)

var WORLD_ID_TO_CODE map[byte]string = map[byte]string{
	0: "Z",
	1: "Y",
	2: "X",
	3: "W",
	4: "V",
	5: "T",
	6: "S",
	7: "R",
}

var WORLD_CHAR_TO_ID map[string]byte = map[string]byte{
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
	if code, ok := WORLD_ID_TO_CODE[id]; ok {
		return code
	} else {
		panic("invalid world delta id")
	}
}

func DecodeWorldDelta(code string) byte {
	if id, ok := WORLD_CHAR_TO_ID[code]; ok {
		return id
	} else {
		panic("invalid world delta code")
	}
}

var SUB_IDS2_TO_CODE map[[2]byte]string = map[[2]byte]string{
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

var SUB_IDS1_TO_CODE map[byte]string = map[byte]string{
	0: "K",
	1: "M",
	2: "N",
	3: "P",
}

/*
   SUB_DELTA_TABLE = [
   ].each(&:freeze).freeze
   SUB_IDS_TO_CHAR = SUB_DELTA_TABLE.inject({}) { |memo, (nums, char)| memo[nums] = char; memo }.freeze
   SUB_CHAR_TO_IDS = SUB_DELTA_TABLE.inject({}) { |memo, (nums, char)| memo[char] = nums; memo }.freeze
*/

func EncodeSubDelta(ids []byte) string {
	if length := len(ids); length > 2 {
		return EncodeSubDelta(ids[0:2]) + EncodeSubDelta(ids[2:])
	} else if length == 2 {
		key := [2]byte{ids[0], ids[1]}
		if code, ok := SUB_IDS2_TO_CODE[key]; ok {
			return code
		}
	} else if length == 1 {
		key := ids[0]
		if code, ok := SUB_IDS1_TO_CODE[key]; ok {
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
	// TODO: mapを使う
	switch code {
	case "2":
		return []byte{0, 0}
	case "3":
		return []byte{0, 1}
	case "4":
		return []byte{0, 2}
	case "5":
		return []byte{0, 3}
	case "6":
		return []byte{1, 0}
	case "7":
		return []byte{1, 1}
	case "8":
		return []byte{1, 2}
	case "A":
		return []byte{1, 3}
	case "B":
		return []byte{2, 0}
	case "C":
		return []byte{2, 1}
	case "D":
		return []byte{2, 2}
	case "E":
		return []byte{2, 3}
	case "F":
		return []byte{3, 0}
	case "G":
		return []byte{3, 1}
	case "H":
		return []byte{3, 2}
	case "J":
		return []byte{3, 3}
	case "K":
		return []byte{0}
	case "M":
		return []byte{1}
	case "N":
		return []byte{2}
	case "P":
		return []byte{3}
	default:
		panic("invalid sub delta code")
	}
}

func Encode(ids []byte) string {
	if length := len(ids); length > 1 {
		return EncodeWorldDelta(ids[0]) + EncodeSubDelta(ids[1:])
	} else if length == 1 {
		return EncodeWorldDelta(ids[0])
	} else {
		panic("delta ids is empty")
	}
}

func Decode(codes string) []byte {
	if length := len(codes); length > 1 {
		return append([]byte{DecodeWorldDelta(codes[0:1])}, DecodeSubDelta(codes[1:])...)
	} else if length == 1 {
		return []byte{DecodeWorldDelta(codes[0:1])}
	} else {
		panic("delta codes is empty")
	}
}
