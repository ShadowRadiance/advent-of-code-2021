# frozen_string_literal: true

require "./dirac_dice/game"
require "./dirac_dice/deterministic_die"

class App
  def run
    input = File.readlines("./data/input.txt", chomp: true)
    p1_start, p2_start = input.map { |str| str[/(\d+)\z/,1].to_i }

    game = DiracDice::Game.new(p1_start, p2_start, die: DeterministicDie.new(100))
    game.play
    puts game.score
  end
end

App.new.run
