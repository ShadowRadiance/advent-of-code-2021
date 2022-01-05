# frozen_string_literal: true

require "./dirac_dice/game2"

class App
  def run
    input = File.readlines("./data/input.txt", chomp: true)
    p1_start, p2_start = input.map { |str| str[/(\d+)\z/,1].to_i }

    game = DiracDice::Game2.new(p1_start, p2_start)
    game.play
    puts game.score
  end
end

App.new.run
