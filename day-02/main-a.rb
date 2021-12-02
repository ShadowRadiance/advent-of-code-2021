require "./depth_scanner"

class App
  def initialize)
  end

  def run
  end

  def load_data_from_input
    File.readlines("./input.txt").map { |line| line.to_i }
  end
end

App.new.run
