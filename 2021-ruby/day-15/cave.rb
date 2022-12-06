require "pp"

class Cave
  attr_reader :cells, :width, :height

  def initialize(data, x: 1)
    lines = data.lines(chomp: true)

    grid = lines.map { |line| line.split(//).map(&:to_i) }.to_a
    grid = embiggen(grid, x)

    @cells = grid.flatten
    @width = grid.first.length
    @height = grid.length
  end

  def grid
    @cells.each_slice(@width).to_a
  end

  def embiggen(grid, multiplier)
    return grid if multiplier == 1

    # duplicate and increase the rows
    grids = (0...multiplier).map do |i|
      grid.map do |row|
        row.map do |cell|
          val = cell + i
          val -= 9 if val >= 10
          val
        end
      end
    end.flatten(1)

    # duplicate and increase the columns
    grids = grids.map do |row|
      (0...multiplier).map do |i|
        row.map do |cell|
          val = cell + i
          val -= 9 if val >= 10
          val
        end
      end.flatten
    end
    
    grids
  end

  def at(x, y)
    @cells[xy_index(x,y)]
  end

  private

  def xy_index(x,y)
    y * @width + x
  end
end
