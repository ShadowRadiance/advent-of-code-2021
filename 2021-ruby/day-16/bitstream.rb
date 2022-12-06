class BitStream
  class EOF < RuntimeError; end

  def self.from_bitstring(bitstring)
    new(bitstring.chars)
  end

  def self.from_hexstring(hexstring)
    bitarray = hexstring.chars.map { |hex| hex.to_i(16).to_s(2).rjust(4, "0") }.join.chars
    new(bitarray)
  end

  def initialize(bitarray)
    @bitarray = bitarray.reverse # reverse since we want to pop/shorten
  end

  def state
    @bitarray.join.reverse
  end

  def empty?
    @bitarray.empty?
  end

  def size
    @bitarray.size
  end

  def pop(n)
    return EOF if @bitarray.empty?
    
    @bitarray.pop(n).join.reverse
  end
end
