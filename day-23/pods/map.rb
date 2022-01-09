# frozen_string_literal: true

require "./pods/hall"
require "./pods/door"
require "./pods/amphipod"
require "byebug"

module Pods
  class Map
    def initialize(string_representation)
      @corridor = [
        Hall.new(self,  0), 
        Hall.new(self,  1), 
        Door.new(self,  2, "A"),
        Hall.new(self,  3),
        Door.new(self,  4, "B"),
        Hall.new(self,  5),
        Door.new(self,  6, "C"),
        Hall.new(self,  7),
        Door.new(self,  8, "D"),
        Hall.new(self,  9),
        Hall.new(self, 10)
      ]
      @amphipods = []
  
      if string_representation[0]=="#"
        parse(string_representation)
      else
        parse_state(string_representation)
      end
    end
  
    def corridor_size
      @corridor.size
    end

    def locations
      @corridor
    end
  
    def doors
      @corridor.select { |location| location.is_a? Door }
    end
  
    def halls
      @corridor.select { |location| location.is_a? Hall }
    end
  
    def door(letter)
      case letter
      when "A" then @corridor[2]
      when "B" then @corridor[4]
      when "C" then @corridor[6]
      when "D" then @corridor[8]
      end
    end

    def apply_move(move)
      move.to.push(move.from.pop)
      move.cost
    end

    def apply_move_str(str)
      apply_move(Move.from_str(str, self))
    end

    def all_possible_moves
      @amphipods.map(&:valid_moves).flatten
    end

    def state
      @corridor.map { |cell| cell.to_s }.join("|")
    end

    def inspect
      "MAP: #{state}"
    end
  
    private
  
    def parse_state(str)
      # str => ".|.|BA|.|CD|.|BC|.|DA|.|."
      # str => ".|.|BA|.|CD|D|BC|.|A|.|."
      str.split("|").each_with_index do |s, index|
        next if s == "."
        location = @corridor[index]
        location = location.room if Door===location
        s.reverse.each_char do |ch|
          amphipod = Amphipod.new(ch, location: nil)
          @amphipods << amphipod
          location.push(amphipod)
        end
      end
    end

    def parse(str)
      # String Representation is expected to be something like
      # 
      # #############
      # #...........#
      # ###B#C#B#D###
      #   #A#D#C#A#
      #   #########
      #
      # where # is a wall, 
      #       . is an empty space, 
      #       ABCD are pods in otherwise empty spaces.
      # so, the top row of dots is a corridor, the other spots are rooms
      # we want to get all the ABCD into the "right" rooms like:
      # 
      # #############
      # #...........#
      # ###A#B#C#D###
      #   #A#B#C#D#
      #   #########
      #
      lines = str.lines(chomp: true).reverse
        # ["  #########",
        #  "  #A#D#C#A#",
        #  "###B#C#B#D###",
        #  "#...........#",
        #  "#############"]
        .slice(1..3)
        # ["  #A#D#C#A#",
        #  "###B#C#B#D###",
        #  "#...........#",]
        .map(&:strip)
        # ["#A#D#C#A#",
        #  "###B#C#B#D###",
        #  "#...........#",]
        .map { |line| line.split("#").reject(&:empty?) }
        # ["A", "D", "C", "A"]
        # ["B", "C", "B", "D"]
        # ["..........."]
  
      rooms = [
        @corridor[2].room,
        @corridor[4].room,
        @corridor[6].room,
        @corridor[8].room,
      ]

      lines[0..1].each do |line|
        line.each.with_index do |char, index|
          next if char=="."
          room = rooms[index]
          amphipod = Amphipod.new(char, location: room)
          @amphipods.push(amphipod)
          room.push(amphipod)
        end
      end
      lines[2][0].each_char.with_index do |char, index|
        hall = @corridor[index]
        next if char=="."
        amphipod = Amphipod.new(char, location: hall)
        @amphipods.push(amphipod)
        hall.push(amphipod)
      end
    end
  end
end
