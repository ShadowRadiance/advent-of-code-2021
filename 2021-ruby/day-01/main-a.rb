require "./depth_scanner"

class App
  attr_reader :scanner
  def initialize(scanner = DepthScanner)
    @scanner = scanner
  end

  def run
    puts scanner.new(load_data_from_input).increases
  end

  def load_data_from_input
    File.readlines("./data/input.txt").map { |line| line.to_i }
  end
end

App.new.run
