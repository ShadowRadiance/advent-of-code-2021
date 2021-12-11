require "./seven_segment_parser"
require "./seven_segment_decoder"

class App
  def initialize
    parser = SevenSegmentParser.new
    @decoder = SevenSegmentDecoder.new(parser.parse(load_data_from_input))
  end

  def run
    @decoder.decode
    puts @decoder.decoded_integers.sum
  end

  def load_data_from_input
    File.read("./data/input.txt")
  end
end

App.new.run
