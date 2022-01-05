# frozen_string_literal: true

module DiracDice
  class Game2

    Player = Struct.new(:pos, :score, keyword_init: true)

    WINNING_SCORE = 21

    def initialize(initial_p1_pos, initial_p2_pos)
      @cache = {}
      @player1 = Player.new(pos: initial_p1_pos - 1, score: 0)
      @player2 = Player.new(pos: initial_p2_pos - 1, score: 0)
    end

    attr_reader :player_one_wins, :player_two_wins
    def play()
      @player_one_wins, @player_two_wins = count_wins(
        @player1,
        @player2,
      )
    end

    def score
      [player_one_wins, player_two_wins].max
    end

    private
    
    # returns [p1 win count, p2 win count] given the current state of p1, p2
    def count_wins(p1, p2)
      return [1, 0] if p1.score >= WINNING_SCORE
      return [0, 1] if p2.score >= WINNING_SCORE
      cached = @cache[[p1.pos, p2.pos, p1.score, p2.score]]
      return cached unless cached.nil?

      wins = [0,0]
      (1..3).each do |roll1|
        (1..3).each do |roll2|
          (1..3).each do |roll3|
            new_pos = (p1.pos + roll1 + roll2 + roll3) % 10
            new_p1 = Player.new(
              pos: new_pos, 
              score: (p1.score + new_pos + 1)
            )
            cascaded = count_wins(p2, new_p1)
            wins = wins.zip(cascaded.reverse).map(&:sum)
          end
        end
      end
      @cache[[p1.pos, p2.pos, p1.score, p2.score]] = wins

      wins
    end
  end
end
