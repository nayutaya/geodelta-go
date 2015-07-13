package encoder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeWorldDelta(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("Z", EncodeWorldDelta(0))
	assert.Equal("Y", EncodeWorldDelta(1))
	assert.Equal("X", EncodeWorldDelta(2))
	assert.Equal("W", EncodeWorldDelta(3))
	assert.Equal("V", EncodeWorldDelta(4))
	assert.Equal("T", EncodeWorldDelta(5))
	assert.Equal("S", EncodeWorldDelta(6))
	assert.Equal("R", EncodeWorldDelta(7))
	// TODO:
	// assert_raises(RuntimeError) { @mod.encode_world_delta(-1) }
	// assert_raises(RuntimeError) { @mod.encode_world_delta(8) }
}

func TestDecodeWorldDelta(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(byte(0), DecodeWorldDelta("Z"))
	assert.Equal(byte(1), DecodeWorldDelta("Y"))
	assert.Equal(byte(2), DecodeWorldDelta("X"))
	assert.Equal(byte(3), DecodeWorldDelta("W"))
	assert.Equal(byte(4), DecodeWorldDelta("V"))
	assert.Equal(byte(5), DecodeWorldDelta("T"))
	assert.Equal(byte(6), DecodeWorldDelta("S"))
	assert.Equal(byte(7), DecodeWorldDelta("R"))
	// assert_raises(RuntimeError) { @mod.decode_world_delta("z") }
	// assert_raises(RuntimeError) { @mod.decode_world_delta("A") }
}

func TestEncodeAndDecodeWorldDelta(t *testing.T) {
	assert := assert.New(t)
	for id := byte(0); id <= 7; id++ {
		encoded1 := EncodeWorldDelta(id)
		decoded1 := DecodeWorldDelta(encoded1)
		encoded2 := EncodeWorldDelta(decoded1)
		assert.Equal(id, decoded1)
		assert.Equal(encoded1, encoded2)
	}
}

func TestEncodeSubDelta1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("2", EncodeSubDelta([]byte{0, 0}))
	assert.Equal("3", EncodeSubDelta([]byte{0, 1}))
	assert.Equal("4", EncodeSubDelta([]byte{0, 2}))
	assert.Equal("5", EncodeSubDelta([]byte{0, 3}))
	assert.Equal("6", EncodeSubDelta([]byte{1, 0}))
	assert.Equal("7", EncodeSubDelta([]byte{1, 1}))
	assert.Equal("8", EncodeSubDelta([]byte{1, 2}))
	assert.Equal("A", EncodeSubDelta([]byte{1, 3}))
	assert.Equal("B", EncodeSubDelta([]byte{2, 0}))
	assert.Equal("C", EncodeSubDelta([]byte{2, 1}))
	assert.Equal("D", EncodeSubDelta([]byte{2, 2}))
	assert.Equal("E", EncodeSubDelta([]byte{2, 3}))
	assert.Equal("F", EncodeSubDelta([]byte{3, 0}))
	assert.Equal("G", EncodeSubDelta([]byte{3, 1}))
	assert.Equal("H", EncodeSubDelta([]byte{3, 2}))
	assert.Equal("J", EncodeSubDelta([]byte{3, 3}))
}

func TestEncodeSubDelta2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("K", EncodeSubDelta([]byte{0}))
	assert.Equal("M", EncodeSubDelta([]byte{1}))
	assert.Equal("N", EncodeSubDelta([]byte{2}))
	assert.Equal("P", EncodeSubDelta([]byte{3}))
}

func TestEncodeSubDelta3(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("2K", EncodeSubDelta([]byte{0, 0, 0}))
	assert.Equal("22", EncodeSubDelta([]byte{0, 0, 0, 0}))
	assert.Equal("3N", EncodeSubDelta([]byte{0, 1, 2}))
	assert.Equal("3E", EncodeSubDelta([]byte{0, 1, 2, 3}))
}

/*
  func test_encode_sub_delta__4(t *testing.T) {
    assert_raises(RuntimeError) { EncodeSubDelta([]) }
    assert_raises(RuntimeError) { EncodeSubDelta([-1]) }
    assert_raises(RuntimeError) { EncodeSubDelta([4]) }
  end
*/

func TestDecodeSubDelta1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]byte{0, 0}, DecodeSubDelta("2"))
	assert.Equal([]byte{0, 1}, DecodeSubDelta("3"))
	assert.Equal([]byte{0, 2}, DecodeSubDelta("4"))
	assert.Equal([]byte{0, 3}, DecodeSubDelta("5"))
	assert.Equal([]byte{1, 0}, DecodeSubDelta("6"))
	assert.Equal([]byte{1, 1}, DecodeSubDelta("7"))
	assert.Equal([]byte{1, 2}, DecodeSubDelta("8"))
	assert.Equal([]byte{1, 3}, DecodeSubDelta("A"))
	assert.Equal([]byte{2, 0}, DecodeSubDelta("B"))
	assert.Equal([]byte{2, 1}, DecodeSubDelta("C"))
	assert.Equal([]byte{2, 2}, DecodeSubDelta("D"))
	assert.Equal([]byte{2, 3}, DecodeSubDelta("E"))
	assert.Equal([]byte{3, 0}, DecodeSubDelta("F"))
	assert.Equal([]byte{3, 1}, DecodeSubDelta("G"))
	assert.Equal([]byte{3, 2}, DecodeSubDelta("H"))
	assert.Equal([]byte{3, 3}, DecodeSubDelta("J"))
}

func TestDecodeSubDelta2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]byte{0}, DecodeSubDelta("K"))
	assert.Equal([]byte{1}, DecodeSubDelta("M"))
	assert.Equal([]byte{2}, DecodeSubDelta("N"))
	assert.Equal([]byte{3}, DecodeSubDelta("P"))
}

func TestDecodeSubDelta3(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]byte{0, 0, 0}, DecodeSubDelta("2K"))
	assert.Equal([]byte{0, 0, 0, 0}, DecodeSubDelta("22"))
	assert.Equal([]byte{0, 1, 2}, DecodeSubDelta("3N"))
	assert.Equal([]byte{0, 1, 2, 3}, DecodeSubDelta("3E"))
}

/*
  func test_decode_sub_delta__4(t *testing.T) {
	assert := assert.New(t)
    assert_raises(RuntimeError) { @mod.decode_sub_delta("") }
    assert_raises(RuntimeError) { @mod.decode_sub_delta("a") }
    assert_raises(RuntimeError) { @mod.decode_sub_delta("Z") }
  }
*/

func TestEncodeAndDecodeSubDelta1(t *testing.T) {
	assert := assert.New(t)
	for id1 := byte(0); id1 <= 3; id1++ {
		encoded1 := EncodeSubDelta([]byte{id1})
		decoded1 := DecodeSubDelta(encoded1)
		encoded2 := EncodeSubDelta(decoded1)
		assert.Equal([]byte{id1}, decoded1)
		assert.Equal(encoded1, encoded2)
	}
}

func TestEncodeAndDecodeSubDelta2(t *testing.T) {
	assert := assert.New(t)
	for id1 := byte(0); id1 <= 3; id1++ {
		for id2 := byte(0); id2 <= 3; id2++ {
			encoded1 := EncodeSubDelta([]byte{id1, id2})
			decoded1 := DecodeSubDelta(encoded1)
			encoded2 := EncodeSubDelta(decoded1)
			assert.Equal([]byte{id1, id2}, decoded1)
			assert.Equal(encoded1, encoded2)
		}
	}
}

func TestEncode(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("Z", Encode([]byte{0}))
	assert.Equal("ZM", Encode([]byte{0, 1}))
	assert.Equal("Z8", Encode([]byte{0, 1, 2}))
	assert.Equal("Z8P", Encode([]byte{0, 1, 2, 3}))
	assert.Equal("R", Encode([]byte{7}))
	assert.Equal("RP", Encode([]byte{7, 3}))
	assert.Equal("RH", Encode([]byte{7, 3, 2}))
	assert.Equal("RHM", Encode([]byte{7, 3, 2, 1}))
	// TODO:
	// assert_raises(RuntimeError) { @mod.encode([]) }
}

func TestDecode(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]byte{0}, Decode("Z"))
	assert.Equal([]byte{0, 1}, Decode("ZM"))
	assert.Equal([]byte{0, 1, 2}, Decode("Z8"))
	assert.Equal([]byte{0, 1, 2, 3}, Decode("Z8P"))
	assert.Equal([]byte{7}, Decode("R"))
	assert.Equal([]byte{7, 3}, Decode("RP"))
	assert.Equal([]byte{7, 3, 2}, Decode("RH"))
	assert.Equal([]byte{7, 3, 2, 1}, Decode("RHM"))
	// TODO:
	// assert_raises(RuntimeError) { @mod.encode("") }
}

/*
  func test_encode_and_decode__rush(t *testing.T) {
	assert := assert.New(t)
    world = (0..7).to_a
    sub   = (0..3).to_a
    1000.times {
      ids = [world[rand(world.size)]] + rand(20).times.map { sub[rand(sub.size)] }
      encoded1 = @mod.encode(ids)
      decoded1 = @mod.decode(encoded1)
      encoded2 = @mod.encode(decoded1)
      assert.Equal(ids, decoded1)
      assert.Equal(encoded1, encoded2)
    }
  }
*/
