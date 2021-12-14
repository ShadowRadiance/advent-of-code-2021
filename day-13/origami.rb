require "byebug"

class Origami
  def initialize(dots)
    initialize_grid(dots)
  end

  def process_one(instruction)
    case instruction.axis
    when "x" then fold_left(instruction.index)
    when "y" then fold_up(instruction.index)
    else raise "that doesn't make any sense"
    end
  end

  def process_all(instructions)
    instructions.each { |i| process_one(i) }
  end

  def display(space = ".", dot = "#")
    chars = [space, dot]
    @grid.flatten.map { |v| chars[v] }.each_slice(@width).map(&:join).join("\n")
  end

  private

  def initialize_grid(dots)
    @width = dots.map(&:x).max + 1
    @height = dots.map(&:y).max + 1
    @grid = Array.new(@height) { Array.new(@width, 0) }
    dots.each { |dot| @grid[dot.y][dot.x] = 1 }
  end

  def fold_up(at_y)
    grid_a = @grid.slice(0...at_y)
    grid_b = @grid.slice(at_y+1..-1).reverse

    pad_height(grid_a, grid_b)

    @height = grid_a.size
    @grid = grid_a.flatten.zip(grid_b.flatten)
                  .map { |a, b| a | b }
                  .each_slice(@width)
                  .to_a
  end

  def fold_left(at_x)
    grid_a = @grid.map { |row| row.slice(0...at_x)  }
    grid_b = @grid.map { |row| row.slice(at_x+1..-1).reverse }

    pad_width(grid_a, grid_b)

    @width = grid_a.flatten.size / @height
    @grid = grid_a.flatten.zip(grid_b.flatten)
                  .map { |a, b| a | b }
                  .each_slice(@width)
                  .to_a
  end

  def pad_height(grid_a, grid_b)
    height_a = grid_a.size
    height_b = grid_b.size

    return if height_a == height_b

    padding = Array.new((height_a-height_b).abs) { Array.new(@width, 0) }
    smaller_grid = height_a < height_b ? grid_a : grid_b
    smaller_grid.replace(padding + smaller_grid)
  end

  def pad_width(grid_a, grid_b)
    width_a = grid_a.flatten.size/@height
    width_b = grid_b.flatten.size/@height

    return if width_a == width_b

    padding = Array.new((width_a-width_b).abs, 0)
    smaller_grid = width_a < width_b ? grid_a : grid_b
    smaller_grid.map! { |row| row.replace(padding + row) }
  end

end