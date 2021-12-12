require "./cave/input"
require "./cave/map"
require "./cave/solver"

class App
  def initialize
    @solver = 
      Cave::Solver.new(
        Cave::Map.new(
          Cave::Input.new(
            load_data_from_input
          )
        ),
        repeats: 1
      )
  end

  def run
    puts @solver.count_paths
  end

  def load_data_from_input
    File.read("./data/input.txt")
  end
end

App.new.run
