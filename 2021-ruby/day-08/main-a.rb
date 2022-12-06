require "./seven_segment_parser"
require "./seven_segment_decoder"

class App
  def initialize
    parser = SevenSegmentParser.new
    @decoder = SevenSegmentDecoder.new(parser.parse(load_data_from_input))
  end

  def run
    @decoder.decode
    # puts @decoder.decoded.map(&:to_s).join("\n")
    puts [1,4,7,8].map { |i| @decoder.count_all(i) }.sum
  end

  def load_data_from_input
    File.read("./data/input.txt")
  end
end

App.new.run
