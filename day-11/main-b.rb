require "./game"

class App
  def initialize
    @game = Game.new(load_data_from_input)
  end

  def run
    steps = 0
    begin
      steps += 1
      @game.step
    end until @game.last_flashes == @game.size
    puts steps
  end

  def load_data_from_input
    File.readlines("./data/input.txt", chomp: true)
      .map { |line| line.split(//).map(&:to_i) }
  end
end

App.new.run
