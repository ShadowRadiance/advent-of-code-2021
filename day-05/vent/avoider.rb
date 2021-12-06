require './vent/parser'

module Vent
  class Avoider
    def initialize(data, parser: Parser.new)
      @data = case data
      when Array then data
      when String then parser.parse(data.lines(chomp: true))
      end
    end

    def overlapping_points
      grid.select { |x| x >= 2 }
    end

    def grid
      return @grid if defined?(@grid)

      largest_x = @data.flat_map {|line| [line.start_point.x, line.end_point.x] }.max
      largest_y = @data.flat_map {|line| [line.start_point.y, line.end_point.y] }.max
      width = largest_x + 1
      height = largest_y + 1
      @grid = Array.new(height * width, 0).tap do |grid|
        @data.select { |line| line.horizontal? || line.vertical? }
          .flat_map { |line| line.points }
          .each { |point| grid[point.y * width + point.x ] += 1 }
        # show_grid(grid, height, width)
      end
    end

    def show_grid(grid, height, width)
      puts
      puts (0...height).map { |y| grid.slice( y*width...(y+1)*width).join }.join("\n").gsub("0", ".")
      puts
    end
  end
end