# frozen_string_literal: true

require "./stack"
require "./pods/move"

module Pods
  class Room
    attr_reader :door, :letter
    def initialize(door, letter:, capacity: 2)
      @door = door
      @letter = letter
      @contents = Stack.new(capacity: capacity)
    end

    def map
      door.map
    end
  
    def full?
      @contents.full?
    end

    def empty?
      @contents.empty?
    end

    def top
      @contents.top
    end

    def amphipods
      @contents.to_a.reverse
    end

    def to_s
      amphipods.map(&:letter).join
    end

    def inspect
      "ROOM #{letter}: #{to_s}"
    end
  
    def pop
      amphipod = @contents.pop
      amphipod.location = nil
      amphipod
    end
  
    def push(amphipod)
      @contents.push(amphipod)
      amphipod.location = self
      self
    end
  
    # Array<Move> [from: to: cost:]
    def accessible_halls_with_cost(per_move_cost)
      locations = door.map.locations
      accessible_halls = []
  
      door_index = door.index
      
      # look left
      (door_index - 1).downto(0) do |index|
        next unless locations[index].is_a? Hall
        break if locations[index].contains_amphipod?
        accessible_halls << locations[index]
      end
  
      # look right
      (door_index + 1).upto(door.map.corridor_size-1) do |index|
        next unless locations[index].is_a? Hall
        break if locations[index].contains_amphipod?
        accessible_halls << locations[index]
      end
  
      accessible_halls.map do |hall|
        Move.new(
          from: self, 
          to: hall,
          cost: per_move_cost * (
            (full? ? 1 : 2) +             # move to door
            (door_index - hall.index).abs # move to target hall
          ),
        )
      end
    end
  end
end




