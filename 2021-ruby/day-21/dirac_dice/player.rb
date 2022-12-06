# frozen_string_literal: true

module DiracDice
  class Player
    attr_reader :score, :name

    def initialize(name, space, game)
      @name = name
      @space = space
      @game = game
      @score = 0
    end

    def take_turn
      @space = (@space + roll(game.die)) % 10
      @space = 10 if @space.zero?

      @score += @space
      
      game.win if @score >= game.target
    end

    def roll(die)
      die.roll + die.roll + die.roll
    end

    private

    attr_reader :game
  end
end
