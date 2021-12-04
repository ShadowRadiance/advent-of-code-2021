require "./bingo/card"

module Bingo
  class Game
    attr_reader :cards, :numbers, :winner, :turns

    def initialize(cards, numbers = [])
      raise ArgumentError if cards.empty?

      @cards = cards.map { |card_array| Card.new(card_array) }
      @numbers = numbers.dup
      @winner = nil
      @turns = 0
    end

    def play
      while winner.nil?
        call_a_number
        raise "Too many turns" if turns > 25
      end
    end

    def call_a_number
      raise "Called number after bingo" unless @winner.nil?

      number = numbers.first
      cards.each { |card| card.mark(number) }
      numbers.rotate!
      @turns += 1
      
      determine_winner
    end

    def determine_winner
      @winner = cards
        .select { |card| card.bingo? }
        .sort { |card| card.score }
        .last
    end

    def winning_score
      return nil if winner.nil?

      winner.unmarked.to_a.sum * numbers.last
    end

  end
end