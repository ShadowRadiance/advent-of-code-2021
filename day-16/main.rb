require "./bitstream"
require "./packet_processor"

begin
  hex = File.readlines("./data/input.txt", chomp:true).first
  bitstream = BitStream.from_hexstring(hex)
  packet = PacketProcessor.from(bitstream)
  puts packet.version_sum
rescue => err
  puts err.message
  puts "#{hex}: ERROR}"
end