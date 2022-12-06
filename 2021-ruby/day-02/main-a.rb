require "./submarine"

class App
  def initialize
    @submarine = Submarine.new
  end

  def run
    @submarine.follow_instructions(load_data_from_input)
    puts @submarine.x_pos
    puts @submarine.depth
    puts @submarine.x_pos * @submarine.depth
  end

  def load_data_from_input
    File.readlines("./data/input.txt")
  end
end

App.new.run
