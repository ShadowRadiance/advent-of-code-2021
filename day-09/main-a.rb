require "./height_map"

class App
  def initialize
    @height_map = HeightMap.load_string(load_data_from_input)
  end

  def run
    puts @height_map.low_point_risk_levels.sum
  end

  def load_data_from_input
    File.read("./data/input.txt")
  end
end

App.new.run
