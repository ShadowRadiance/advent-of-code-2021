require "./depth_scanner"

class App
  attr_reader :scanner, :strategy
  def initialize(scanner = DepthScanner, strategy: SimpleDepthScanner.new)
    @scanner = scanner
    @strategy = strategy
  end

  def run
    puts scanner.new(load_data_from_input, strategy: strategy).increases
  end

  def load_data_from_input
    File.readlines("./data/input.txt").map { |line| line.to_i }
  end
end

App.new(strategy: SlidingWindowScanner.new(3)).run
