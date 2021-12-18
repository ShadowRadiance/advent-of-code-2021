require "./packet/literal"
require "./packet/operator"

class PacketProcessor
  def self.from(bitstream)
    return nil if bitstream.empty?

    version = bitstream.pop(3).to_i(2)
    type = bitstream.pop(3).to_i(2)
    case type
    when 4 then Literal.new(version, type, bitstream)
    else Operator.new(version, type, bitstream)
    end
  end
end
