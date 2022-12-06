require "./bingo/card"

module Bingo
  class Game
    attr_reader :cards, :numbers, :winner, :turns, :ranked_cards, :unranked_cards

    def initialize(cards, numbers = [])
      raise ArgumentError if cards.empty?

      @cards = cards.map { |card_array| Card.new(card_array) }
      @numbers = numbers.dup
      @winner = nil
      @turns = 0
      @unranked_cards = @cards.dup
      @ranked_cards = []
    end

    def play
      while winner.nil?
        call_a_number
        determine_winner
      end
    end

    def play_all_cards
      while unranked_cards.length > 0
        call_a_number
        rank_finished_cards
      end
    end

    def call_a_number
      raise "Too many turns" if turns > numbers.length
        
      number = numbers.first
      unranked_cards.each { |card| card.mark(number) }
      numbers.rotate!
      @turns += 1
      log("TURN: #{turns} CALL: #{number}")
    end

    def determine_winner
      @winner = cards
        .select { |card| card.bingo? }
        .sort { |card| score(card) }
        .last
    end

    def rank_finished_cards
      finished_cards = unranked_cards.select(&:bingo?)
      finished_cards.each do |card|
        ranked_cards << { card: card, score: score(card) }
        unranked_cards.delete(card)
      end
      log(
        "FINISHED:", ranked_cards, 
        "UNFINISHED:", unranked_cards.map { |card| card.to_s_marked }
      )
    end

    def score(card)
      card.unmarked.to_a.sum * numbers.last
    end

    def winning_score
      return nil if winner.nil?
      score(winner)
    end

    def log(*lines)
      if ENV['enable_logging']
        lines.each { |line| puts line }
      end
    end
  end
end