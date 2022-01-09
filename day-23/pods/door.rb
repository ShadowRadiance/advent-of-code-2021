# frozen_string_literal: true

require "./pods/room"

module Pods
  class Door
    attr_reader :map, :index, :letter, :room
    def initialize(map, index, letter, room_size: 2)
      @map = map
      @index = index
      @letter = letter
      @room = Room.new(self, letter: letter, capacity: room_size)
    end
    
    def to_s
      @room.to_s
    end

    def inspect
      "DOOR #{letter}: #{@room.inspect}"
    end
  end
end
