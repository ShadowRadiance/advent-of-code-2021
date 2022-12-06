# frozen_string_literal: true

module Pods
  class Move
    attr_reader :from, :to, :cost
    def initialize(from:, to:, cost:)
      @from = from
      @to = to
      @cost = cost
    end
  
    def to_h
      { from: from, to: to, cost: cost }
    end
  
    def to_a
      [from, to, cost]
    end
  
    def to_s
      "[#{encode(from)}->#{encode(to)}]:#{cost}"
    end

    DECODER = /\[(room|hall)(\d+|A|B|C|D)->(room|hall)(\d+|A|B|C|D)\]:(\d+)/

    def self.from_str(str, map)
      from_type, from_detail, to_type, to_detail, cost = DECODER.match(str).captures
      
      new(
        from: decode(from_type, from_detail, map),
        to: decode(to_type, to_detail, map),
        cost: cost.to_i
      )
    end
  
    def self.decode(type, detail, map)
      case type
      when "room" then map.door(detail).room
      when "hall" then map.locations[detail.to_i]
      end
    end

    def inspect
      "MOVE #{self}"
    end

    def encode(location)
      case location
      when Room then "room#{location.letter}"
      when Hall then "hall#{location.index}"
      end
    end
  end
end
