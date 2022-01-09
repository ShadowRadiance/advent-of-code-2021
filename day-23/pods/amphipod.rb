# frozen_string_literal: true

require "./pods/room"
require "./pods/hall"

module Pods
  class Amphipod
    attr_reader :letter, :location
    def initialize(letter, location:)
      @letter = letter
      @cost_to_move = case letter
      when "A" then 1
      when "B" then 10
      when "C" then 100
      when "D" then 1000
      end
      @location = location
    end
  
    def location=(location)
      @location = location
    end
  
    # Array<Move> [from: to: cost:]
    def valid_moves
      # no moves if we're in the right location!
      return [] if location==target_room && target_room.amphipods.map(&:letter).all?(letter)

      case location
      when Room
        return [] unless location.top == self

        moves = location.accessible_halls_with_cost(@cost_to_move)
        moves += [move_to_own_room] if can_move_to_own_room?
        moves
      when Hall
        return [] unless can_move_to_own_room?
        [move_to_own_room]
      end
    end

    def target_room
      @target_room ||= location.map.door(letter).room
    end

    def can_move_to_own_room?
      return false if target_room.full?
      return false if target_room.amphipods.map(&:letter).any? { |other| other != letter }
      return false if intervening_halls(target_room).any?(&:contains_amphipod?)
      true
    end

    def intervening_halls(room)
      current_index = location.is_a?(Room) ? location.door.index : location.index
      first, second = *([current_index, room.door.index].sort)
      (first+1..second-1).to_a
        .map { |index| location.map.locations[index] }
        .select { |location| Hall===location }
    end

    def move_to_own_room
      distance = 
        case location
        when Room
          [
            (location.full? ? 1 : 2),                           # move out of current room
            (location.door.index - target_room.door.index).abs, # move along the hall
            (target_room.empty? ? 2 : 1),                       # move into target room
          ].sum
        when Hall
          [
            (location.index - target_room.door.index).abs,  # move along the hall
            (target_room.empty? ? 2 : 1),                   # move into target room
          ].sum
        end

      Move.new(
        from: location, 
        to: target_room, 
        cost: @cost_to_move * distance,
      )
    end
  end
end
