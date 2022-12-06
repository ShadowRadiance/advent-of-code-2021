require './vent/avoider'

class App
  def initialize
    @avoider = Vent::Avoider.new(load_data_from_input)
  end

  def run
    puts @avoider.overlapping_points.count
  end

  def load_data_from_input
    lines = File.read("./data/input.txt")
  end
end

App.new.run
