require "./game"

class App
  def initialize
    @game = Game.new(load_data_from_input)
  end

  def run
    100.times { @game.step }
    puts @game.total_flashes
  end

  def load_data_from_input
    File.readlines("./data/input.txt", chomp: true)
      .map { |line| line.split(//).map(&:to_i) }
  end
end

App.new.run
