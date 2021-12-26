# frozen_string_literal: true

require "set"

class InfiniteImage
  attr_reader :min_x, :min_y, :max_x, :max_y, :width, :height

  # example data: [ "#..#.",
  #                 "#....",
  #                 "##..#",
  #                 "..#..",
  #                 "..###"]
  def initialize(data)
    data = preparse(data)
    @width = data.first&.length
    @height = data.length
    @min_x, @max_x = 0, @width - 1
    @min_y, @max_y = 0, @height - 1
    
    @orig_data = data # array of strings
    @data = Set.new   # array of on-lights
    
    # puts data.inspect

    data.each.with_index do |row, y|
      row.chars.each.with_index do |cell, x|
        @data << [x, y] if cell == "#"
      end
    end
  end

  def preparse(data)
    # throw away blank rows at the start and end
    # throw away blank cols at the start and end
    data
  end

  def nine_lights_around(x, y)
    ((y-1)..(y+1)).map { |b|
      ((x-1)..(x+1)).map { |a| 
        @data.include?([a, b]) ? "#" : "."
      }
    }
    # .tap do |arr|
    #   puts x
    #   puts y
    #   puts arr.inspect
    # end
  end

  def count_lit
    @data.size    
  end

  def to_s
    @orig_data.join("\n")
  end
end