require './crab_optimizer'

class App
  def initialize
    @crab_optimizer = CrabOptimizer.new(load_data_from_input)
  end

  def run
    puts @crab_optimizer.optimal_alignment_fuel
  end

  def load_data_from_input
    File.readlines("./data/input.txt").first.split(/,/).map(&:to_i)
  end
end

App.new.run
