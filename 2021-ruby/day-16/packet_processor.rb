require "./packet/literal"
require "./packet/operator/sum"
require "./packet/operator/product"
require "./packet/operator/minimum"
require "./packet/operator/maximum"
require "./packet/operator/greater_than"
require "./packet/operator/less_than"
require "./packet/operator/equal_to"

class PacketProcessor
  def self.from(bitstream)
    return nil if bitstream.empty?

    version = bitstream.pop(3).to_i(2)
    type = bitstream.pop(3).to_i(2)
    klass = case type
    when 0 then Sum
    when 1 then Product
    when 2 then Minimum
    when 3 then Maximum
    when 4 then Literal
    when 5 then GreaterThan
    when 6 then LessThan
    when 7 then EqualTo
    end
    klass.new(version, type, bitstream)
  end
end

