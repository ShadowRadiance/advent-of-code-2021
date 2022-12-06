require "./polymerizer"
# require "byebug"

class App
  def initialize
    load_data_from_input
    @poly = Polymerizer.new(@rules)
  end

  def run
    result = @poly.steps(10, @template)

    counts = result.chars.group_by(&:itself).transform_values(&:size)
    min_count, max_count = counts.values.minmax

    puts max_count - min_count
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
