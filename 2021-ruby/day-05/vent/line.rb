require './vent/point'

module Vent
  class Line
    attr_reader :start_point, :end_point
    def initialize(start_point, end_point)
      @start_point = start_point
      @end_point = end_point
    end

    def points
      if vertical?
        min, max = [start_point.y, end_point.y].minmax
        (min..max).map { |y_val| Point.new(start_point.x, y_val) }
      elsif horizontal?
        min, max = [start_point.x, end_point.x].minmax
        (min..max).map { |x_val| Point.new(x_val, start_point.y) }
      else # pure diagonal
        x_distance = end_point.x - start_point.x # can be negative
        x_sign = x_distance.positive? ? 1 : -1
        y_distance = end_point.y - start_point.y # can be negative
        y_sign = y_distance.positive? ? 1 : -1
        (0..x_distance.abs).map do |index|
          Point.new(start_point.x + index * x_sign , start_point.y + index * y_sign)
        end
      end
    end

    def horizontal?
      start_point.y == end_point.y
    end

    def vertical?
      start_point.x == end_point.x
    end
  end
end