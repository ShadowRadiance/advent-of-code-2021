# frozen_string_literal: true

require "./pods/room"

module Pods
  class Door
    attr_reader :map, :index, :letter, :room
    def initialize(map, index, letter)
      @map = map
      @index = index
      @letter = letter
      @room = Room.new(self, letter: letter)
    end
    
    def to_s
      @room.to_s
    end

    def inspect
      "DOOR #{letter}: #{@room.inspect}"
    end
  end
end
