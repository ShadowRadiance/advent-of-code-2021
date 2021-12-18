require "./bitstream"
require "./packet"
require "./packet_processor"

class Operator < Packet
  def initialize(version, type, bitstream)
    super(version, type)
    length_type = slurp(bitstream,1).to_i(2)
    case length_type
    when 0
      subpacket_length = slurp(bitstream,15).to_i(2)
      process_subpackets_by_length(subpacket_length, bitstream)
    when 1
      number_of_direct_children = slurp(bitstream,11).to_i(2)
      process_subpackets_by_count(number_of_direct_children, bitstream)
    end
  end

  private

  def process_subpackets_by_length(length, bitstream)
    sub_bitstream = BitStream.from_bitstring(slurp(bitstream, length))
    loop do
      subpacket = PacketProcessor.from(sub_bitstream)
      break unless subpacket
      subpackets << subpacket
    end
  end

  def process_subpackets_by_count(count, bitstream)
    count.times do
      subpacket = PacketProcessor.from(bitstream)
      subpackets << subpacket
    end
    raise "Extracted #{subpackets.size}/#{count} subpackets from bitstream" if subpackets.size != subpackets.count
  end
end
