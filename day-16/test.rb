require "./bitstream"
require "./packet_processor"

examples = {
  "8A004A801A8002F478" => 16,
  "620080001611562C8802118E34" => 12,
  "C0015000016115A2E0802F182340" => 23,
  "A0016C880162017C3686B18A3D4780" => 31
}
examples.each do |hex, expected|
  actual = PacketProcessor.from(BitStream.from_hexstring(hex)).version_sum
  puts "#{hex}: {expected: #{expected} actual: #{actual} --- #{ expected==actual ? "OK" : "XX" }}"
rescue => err
  puts err.message
  puts "#{hex}: {expected: #{expected} ERROR}"
end
