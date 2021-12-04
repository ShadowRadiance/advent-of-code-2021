require "./bingo/game"

class App
  def initialize
    cards, numbers = load_data_from_input
    @game = Bingo::Game::new(cards, numbers)
  end

  def run
    @game.play
    puts "Winning Score: #{@game.winning_score}"
    puts "Winning Board:"
    puts @game.winner
  end

  def load_data_from_input
    lines = File.readlines("./data/input.txt", chomp: true)
    numbers = lines.shift.split(",").map(&:to_i)
    lines.shift # blank line

    cards = []
    loop do
      card_lines = lines.shift(5).map do |card_line|
        card_line.split(" ").map(&:to_i)
      end
      lines.shift # blank line
      
      break if card_lines.length != 5

      cards << card_lines.flatten
    end

    [cards, numbers]
  end
end

App.new.run
