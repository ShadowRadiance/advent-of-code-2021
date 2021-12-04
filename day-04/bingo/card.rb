require "set"

module Bingo
  class Card
    attr_reader :data, :marked, :unmarked
    def initialize(data)
      raise ArgumentError unless data.size == 25

      @data = data.dup
      @marked = Set.new
      @unmarked = Set.new(data)
    end

    def mark(number)
      if data.include? number
        @marked << number
        @unmarked.delete(number)
      end
    end

    def bingo?
      lines.any? { |line| (line - marked.to_a).empty? }
    end

    def lines
      @lines ||= (0..4).map { |row| data[row*5...(row+1)*5] }
    end

    def columns
      lines.transpose
    end

    def to_s
      lines.map { |line| 
        line.map { |number| 
          number.to_s.rjust(2) 
        }.join(" ") 
      }.join("\n")
    end
  end
end