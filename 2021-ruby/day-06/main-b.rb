require './lanternfish_simulator'

class App
  def initialize
    @lanternfish_simulator = LanternfishSimulator.new(load_data_from_input)
  end

  def run
    @lanternfish_simulator.run(days: 256)
    puts @lanternfish_simulator.count
  end

  def load_data_from_input
    File.readlines("./data/input.txt").first.split(/,/).map(&:to_i)
  end
end

App.new.run
