# frozen_string_literal: true

require "./stack"

module Pods
  class Hall
    attr_reader :map, :index, :amphipod
    def initialize(map, index)
      @map = map
      @index = index
      @amphipod = nil
    end
  
    def push(amphipod)
      raise Stack::Overflow unless @amphipod.nil?
      @amphipod = amphipod
      amphipod.location = self
    end
  
    def pop
      raise Stack::Underflow if @amphipod.nil?
      result = @amphipod
      @amphipod = nil
      result.location = nil
      result
    end
  
    def door(letter)
      map.door(letter)
    end
  
    def contains_amphipod?
      !@amphipod.nil?
    end

    def to_s
      @amphipod&.letter || "."
    end

    def inspect
      "HALL #{index}: #{to_s}"
    end

  end
end

