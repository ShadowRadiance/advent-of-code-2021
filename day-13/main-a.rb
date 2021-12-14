require "./origami"
require "./origami_parser"

class App
  def initialize
    @parser = OrigamiParser.new(load_data_from_input)
    @origami = Origami.new(@parser.dots)
  end

  def run
    @origami.process_one(@parser.instructions[0])
    puts @origami.display
    puts @origami.display.count("#")
  end

  def load_data_from_input
    File.read("./data/input.txt")
  end
end

App.new.run
