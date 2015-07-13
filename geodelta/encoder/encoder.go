package encoder

/*
   WORLD_DELTA_TABLE = [
     [0, "Z"],
     [1, "Y"],
     [2, "X"],
     [3, "W"],
     [4, "V"],
     [5, "T"],
     [6, "S"],
     [7, "R"],
   ].each(&:freeze).freeze
   WORLD_ID_TO_CHAR = WORLD_DELTA_TABLE.inject({}) { |memo, (num, char)| memo[num] = char; memo }.freeze
   WORLD_CHAR_TO_ID = WORLD_DELTA_TABLE.inject({}) { |memo, (num, char)| memo[char] = num; memo }.freeze
*/

func EncodeWorldDelta(id uint8) string {
	// return WORLD_ID_TO_CHAR[id] || raise("invalid world delta id -- #{id}")
	// TODO: mapに変更する。
	switch id {
	case 0:
		return "Z"
	case 1:
		return "Y"
	case 2:
		return "X"
	case 3:
		return "W"
	case 4:
		return "V"
	case 5:
		return "T"
	case 6:
		return "S"
	case 7:
		return "R"
	default:
		return ""
	}
}

func DecodeWorldDelta(code string) uint8 {
	// TODO: mapにする
	//    return WORLD_CHAR_TO_ID[code] || raise("invalid world delta code -- #{code}")
	switch code {
	case "Z":
		return 0
	case "Y":
		return 1
	case "X":
		return 2
	case "W":
		return 3
	case "V":
		return 4
	case "T":
		return 5
	case "S":
		return 6
	case "R":
		return 7
	default:
		return 8
	}
}

/*
   SUB_DELTA_TABLE = [
     [[0, 0], "2"],
     [[0, 1], "3"],
     [[0, 2], "4"],
     [[0, 3], "5"],
     [[1, 0], "6"],
     [[1, 1], "7"],
     [[1, 2], "8"],
     [[1, 3], "A"],
     [[2, 0], "B"],
     [[2, 1], "C"],
     [[2, 2], "D"],
     [[2, 3], "E"],
     [[3, 0], "F"],
     [[3, 1], "G"],
     [[3, 2], "H"],
     [[3, 3], "J"],
     [[0]   , "K"],
     [[1]   , "M"],
     [[2]   , "N"],
     [[3]   , "P"],
   ].each(&:freeze).freeze
   SUB_IDS_TO_CHAR = SUB_DELTA_TABLE.inject({}) { |memo, (nums, char)| memo[nums] = char; memo }.freeze
   SUB_CHAR_TO_IDS = SUB_DELTA_TABLE.inject({}) { |memo, (nums, char)| memo[char] = nums; memo }.freeze
*/

func EncodeSubDelta(ids []byte) string {
	// TODO: mapを使う
	// raise("sub delta ids is empty") if ids.empty?
	// return ids.each_slice(2).map { |part|
	//   SUB_IDS_TO_CHAR[part] || raise("invalid sub delta ids -- #{part}")
	// }.join("")
	if len(ids) > 2 {
		return EncodeSubDelta(ids[0:2]) + EncodeSubDelta(ids[2:])
	} else if len(ids) == 2 {
		switch {
		case ids[0] == 0 && ids[1] == 0:
			return "2"
		case ids[0] == 0 && ids[1] == 1:
			return "3"
		case ids[0] == 0 && ids[1] == 2:
			return "4"
		case ids[0] == 0 && ids[1] == 3:
			return "5"
		case ids[0] == 1 && ids[1] == 0:
			return "6"
		case ids[0] == 1 && ids[1] == 1:
			return "7"
		case ids[0] == 1 && ids[1] == 2:
			return "8"
		case ids[0] == 1 && ids[1] == 3:
			return "A"
		case ids[0] == 2 && ids[1] == 0:
			return "B"
		case ids[0] == 2 && ids[1] == 1:
			return "C"
		case ids[0] == 2 && ids[1] == 2:
			return "D"
		case ids[0] == 2 && ids[1] == 3:
			return "E"
		case ids[0] == 3 && ids[1] == 0:
			return "F"
		case ids[0] == 3 && ids[1] == 1:
			return "G"
		case ids[0] == 3 && ids[1] == 2:
			return "H"
		case ids[0] == 3 && ids[1] == 3:
			return "J"
		default:
			return ""
		}
	} else if len(ids) == 1 {
		switch ids[0] {
		case 0:
			return "K"
		case 1:
			return "M"
		case 2:
			return "N"
		case 3:
			return "P"
		default:
			return ""
		}
	} else {
		return ""
	}
}

/*
   def self.decode_sub_delta(codes)
     raise("sub delta codes is empty") if codes.empty?
     return codes.chars.inject([]) { |memo, char|
       memo + (SUB_CHAR_TO_IDS[char] || raise("invalid sub delta code -- #{char}"))
     }
   end

   def self.encode(ids)
     raise("delta ids is empty") if ids.empty?
     result  = self.encode_world_delta(ids[0])
     result += self.encode_sub_delta(ids[1..-1]) if ids.size >= 2
     return result
   end

   def self.decode(codes)
     raise("delta codes is empty") if codes.empty?
     result  = [self.decode_world_delta(codes[0])]
     result += self.decode_sub_delta(codes[1..-1]) if codes.size >= 2
     return result
   end
*/
