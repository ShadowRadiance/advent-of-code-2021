require "./polymerizer_pairs"
# require "byebug"

class App
  def initialize
    load_data_from_input
    @poly_pairs = PolymerizerPairs.new(@rules, @template)
  end

  def run
    @poly_pairs.steps(40)

    puts @poly_pairs.variance
  end

  def load_data_from_input
    lines = File.readlines("./data/input.txt", chomp: true)
    @template = lines.first

    lines.shift(2)

    @rules = lines.map { |line|
      # "OB -> C"
      line.split(" -> ")
    }.to_h
  end
end

App.new.run
