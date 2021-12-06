require './vent/line'
require './vent/point'

module Vent
  class Parser
    def parse(data)
      # [
      #   "a,b -> c,d"
      #   "e,f -> g,h"
      # ]
      data.map do |line|
        a,b,c,d = *line.match(/(\d+),(\d+) -> (\d+),(\d+)/).values_at(1,2,3,4).map(&:to_i)
        Line.new(Point.new(a,b), Point.new(c,d))
      end
    end
  end
end