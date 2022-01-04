# frozen_string_literal: true

require "./dirac_dice/player"
require "./dirac_dice/die"

module DiracDice
  class Game
    attr_reader :target, :die, :winner, :loser

    def initialize(p1_start, p2_start, die: Die.new, target: 1000)
      @active_player = Player.new("P1", p1_start, self)
      @inactive_player = Player.new("P2", p2_start, self)
      @die = die
      @target = target

      @game_over = false

      @winner = nil
      @loser = nil
    end

    DIRAC_DICE_DISTRIBUTION = {
      3 => 1 , # 111
      4 => 3 , # 112, 121, 211
      5 => 6 , # 113, 122, 131, 212, 221, 311
      6 => 7 , # 123, 132, 213, 222, 231, 312, 321
      7 => 6 , # 133, 223, 232, 313, 322, 331
      8 => 3 , # 233, 323, 332
      9 => 1 , # 333
    }.freeze

    def play_all
      # uuuuhhhhh????
    end

    def play
      until @game_over do
        take_turn
        swap_active
      end
    end

    def win
      @game_over = true
      @winner = @active_player
      @loser = @inactive_player
    end

    def score
      @loser.score * @die.times_rolled
    end

    def take_turn
      @active_player.take_turn
    end

    def swap_active
      @inactive_player, @active_player = @active_player, @inactive_player
    end
  end
end