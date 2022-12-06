require "./cave"
require "./cave_navigator"

class App
  def initialize
    @cave_navigator = CaveNavigator.new(Cave.new(load_data_from_input, x: 5))
  end

  def run
    puts @cave_navigator.shortest_route
  end

  def load_data_from_input
    File.read("./data/input.txt")
  end
end

App.new.run
