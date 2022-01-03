# frozen_string_literal: true

require "byebug"

class InfiniteImage
  attr_reader :min_x, :min_y, :max_x, :max_y, :width, :height, :background

  # example data: [ "#..#.",
  #                 "#....",
  #                 "##..#",
  #                 "..#..",
  #                 "..###"]
  def initialize(data, background: ".")
    raise ArgumentError unless data.length > 0 && data.first.length > 0

    @background = background
    
    data = preparse(data)
    @width = data.first&.length
    @height = data.length
    @min_x, @max_x = 0, @width - 1
    @min_y, @max_y = 0, @height - 1
    
    @data = data.join
  end

  def preparse(data)
    working = data.dup

    blank_row = background * working.first.length
    
    # throw away blank rows at the start and end

    working.shift while working.first == blank_row
    working.pop while working.last == blank_row

    # throw away blank cols at the start and end
    scsmb = shortest_common_start_matching_background(working)
    scfmb = shortest_common_final_matching_background(working)
    working.map! { |str| str[scsmb..-scfmb-1] } unless scsmb.zero? && scfmb.zero?

    working
  end

  def non_background
    @non_background ||= (background=="." ? "#" : ".")
  end

  def shortest_common_start_matching_background(arr)
    arr.map { |str| str.index(non_background) }.min
  end

  def shortest_common_final_matching_background(arr)
    arr.map { |str| str.reverse.index(non_background) }.min
  end

  def at(x, y)
    index = xy_index(x, y)
    return background if index.nil?

    @data[index]
  end

  def nine_lights_around(x, y)
    ((y-1)..(y+1)).map { |b| 
      ((x-1)..(x+1)).map { |a| 
        at(a, b) 
      } 
    }
  end

  def count_lit
    return Float::INFINITY if @background == "#"
    
    @data.count("#")
  end

  def count_unlit
    return Float::INFINITY if @background == "."
    
    @data.count(".")
  end

  def to_s
    @data.chars.each_slice(width).to_a.map(&:join).join("\n")
  end

  private
  
  def xy_index(x, y)
    return nil if x < 0
    return nil if y < 0
    return nil if x > max_x
    return nil if y > max_y

    y * width + x
  end
end