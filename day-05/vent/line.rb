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
      else
        raise NotImplementedError
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