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
	assert.Equal(uint8(0), DecodeWorldDelta("Z"))
	assert.Equal(uint8(1), DecodeWorldDelta("Y"))
	assert.Equal(uint8(2), DecodeWorldDelta("X"))
	assert.Equal(uint8(3), DecodeWorldDelta("W"))
	assert.Equal(uint8(4), DecodeWorldDelta("V"))
	assert.Equal(uint8(5), DecodeWorldDelta("T"))
	assert.Equal(uint8(6), DecodeWorldDelta("S"))
	assert.Equal(uint8(7), DecodeWorldDelta("R"))
	// assert_raises(RuntimeError) { @mod.decode_world_delta("z") }
	// assert_raises(RuntimeError) { @mod.decode_world_delta("A") }
}

func TestEncodeAndDecodeWorldDelta(t *testing.T) {
	assert := assert.New(t)
	for id := uint8(0); id <= 7; id++ {
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
  func test_encode_sub_delta__4
    assert_raises(RuntimeError) { EncodeSubDelta([]) }
    assert_raises(RuntimeError) { EncodeSubDelta([-1]) }
    assert_raises(RuntimeError) { EncodeSubDelta([4]) }
  end

  func test_decode_sub_delta__1
    assert.Equal([0, 0], @mod.decode_sub_delta("2"))
    assert.Equal([0, 1], @mod.decode_sub_delta("3"))
    assert.Equal([0, 2], @mod.decode_sub_delta("4"))
    assert.Equal([0, 3], @mod.decode_sub_delta("5"))
    assert.Equal([1, 0], @mod.decode_sub_delta("6"))
    assert.Equal([1, 1], @mod.decode_sub_delta("7"))
    assert.Equal([1, 2], @mod.decode_sub_delta("8"))
    assert.Equal([1, 3], @mod.decode_sub_delta("A"))
    assert.Equal([2, 0], @mod.decode_sub_delta("B"))
    assert.Equal([2, 1], @mod.decode_sub_delta("C"))
    assert.Equal([2, 2], @mod.decode_sub_delta("D"))
    assert.Equal([2, 3], @mod.decode_sub_delta("E"))
    assert.Equal([3, 0], @mod.decode_sub_delta("F"))
    assert.Equal([3, 1], @mod.decode_sub_delta("G"))
    assert.Equal([3, 2], @mod.decode_sub_delta("H"))
    assert.Equal([3, 3], @mod.decode_sub_delta("J"))
  end

  func test_decode_sub_delta__2
    assert.Equal([0], @mod.decode_sub_delta("K"))
    assert.Equal([1], @mod.decode_sub_delta("M"))
    assert.Equal([2], @mod.decode_sub_delta("N"))
    assert.Equal([3], @mod.decode_sub_delta("P"))
  end

  func test_decode_sub_delta__3
    assert.Equal([0, 0, 0],    @mod.decode_sub_delta("2K"))
    assert.Equal([0, 0, 0, 0], @mod.decode_sub_delta("22"))
    assert.Equal([0, 1, 2],    @mod.decode_sub_delta("3N"))
    assert.Equal([0, 1, 2, 3], @mod.decode_sub_delta("3E"))
  end

  func test_decode_sub_delta__4
    assert_raises(RuntimeError) { @mod.decode_sub_delta("") }
    assert_raises(RuntimeError) { @mod.decode_sub_delta("a") }
    assert_raises(RuntimeError) { @mod.decode_sub_delta("Z") }
  end

  func test_encode_and_decode_sub_delta__1
    (0..3).each { |id1|
      encoded1 = EncodeSubDelta([id1])
      decoded1 = @mod.decode_sub_delta(encoded1)
      encoded2 = EncodeSubDelta(decoded1)
      assert.Equal([id1], decoded1)
      assert.Equal(encoded1, encoded2)
    }
  end

  func test_encode_and_decode_sub_delta__2
    (0..3).each { |id1|
      (0..3).each { |id2|
        encoded1 = EncodeSubDelta([id1, id2])
        decoded1 = @mod.decode_sub_delta(encoded1)
        encoded2 = EncodeSubDelta(decoded1)
        assert.Equal([id1, id2], decoded1)
        assert.Equal(encoded1, encoded2)
      }
    }
  end

  func test_encode
    assert.Equal("Z",   @mod.encode([0]))
    assert.Equal("ZM",  @mod.encode([0, 1]))
    assert.Equal("Z8",  @mod.encode([0, 1, 2]))
    assert.Equal("Z8P", @mod.encode([0, 1, 2, 3]))
    assert.Equal("R",   @mod.encode([7]))
    assert.Equal("RP",  @mod.encode([7, 3]))
    assert.Equal("RH",  @mod.encode([7, 3, 2]))
    assert.Equal("RHM", @mod.encode([7, 3, 2, 1]))
    assert_raises(RuntimeError) { @mod.encode([]) }
  end

  func test_decode
    assert.Equal([0],          @mod.decode("Z"))
    assert.Equal([0, 1],       @mod.decode("ZM"))
    assert.Equal([0, 1, 2],    @mod.decode("Z8"))
    assert.Equal([0, 1, 2, 3], @mod.decode("Z8P"))
    assert.Equal([7],          @mod.decode("R"))
    assert.Equal([7, 3],       @mod.decode("RP"))
    assert.Equal([7, 3, 2],    @mod.decode("RH"))
    assert.Equal([7, 3, 2, 1], @mod.decode("RHM"))
    assert_raises(RuntimeError) { @mod.encode("") }
  end

  func test_encode_and_decode__rush
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
  end
*/
