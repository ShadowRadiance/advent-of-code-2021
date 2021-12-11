require "./octopus"
require "byebug"

class Game
  def initialize(board2d=[[]])
    @board = board2d.map do |row|
      row.map do |energy|
        Octopus.new(energy)
      end
    end
    @flattened = @board.flatten
    setup_neighbors
    @total_flashes = 0
    # output_board
  end

  def step
    # puts "STEP!"
    power_up                   #   .tap { output_board }
    cascade                    #   .tap { output_board }
    count_flashes              #   .tap { puts @total_flashes }
    reset                      #   .tap { output_board }
  end

  def octopus_board
    @board
  end

  def energy_board
    @board.map do |row|
      row.map do |octopus|
        octopus.energy_level
      end
    end
  end

  def setup_neighbors
    @board.each.with_index do |row, y|
      row.each.with_index do |octopus, x|
        # octopus at x,y
        octopus.neighbors = neighbors_of([x,y]).map { |a, b| @board[b][a] }
      end
    end
  end

  def neighbors_of(location)
    x, y = location

    [
      [x-1, y-1], [x,   y-1], [x+1, y-1],
      [x-1, y  ],             [x+1, y  ],
      [x-1, y+1], [x,   y+1], [x+1, y+1],
    ].select { |neighbor| valid_position?(neighbor) }
  end

  def valid_position?(pos)
    x, y = pos
    return false if x < 0
    return false if y < 0
    return false if y >= @board.length
    return false if x >= @board.first.length
    true
  end

  def power_up
    @flattened.each(&:power_up)
  end

  def cascade
    @flattened.each { |octopus| octopus.flash_cascade(initiator: true) }
  end

  def reset
    @flattened.each(&:reset)
  end

  def output_board
    puts
    puts "-" * energy_board.first.length
    puts energy_board.map { |row| row.join(" ") }.join("\n")
    puts "-" * energy_board.first.length
  end

  def total_flashes
    @total_flashes
  end

  def count_flashes
    flashes = @flattened.map do |octopus|
      octopus.has_flashed? ? 1 : 0
    end.sum
    @total_flashes += flashes
  end
end