require "./packet"

class Literal < Packet
  attr_reader :value
  def initialize(version, type, bitstream)
    super(version, type)
    @value = build_value(bitstream)
  end

  private

  def build_value(bitstream)
    value = ""
    continues = true
    while continues
      continues = slurp(bitstream,1).to_i(2) == 1
      value += slurp(bitstream,4)
    end
    value.to_i(2)
  end
end
