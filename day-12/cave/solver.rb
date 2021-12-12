require "./cave/map"
require "byebug"

module Cave
  class Solver
    def initialize(map, repeats: 0)
      @map = map
      @repeats = repeats
    end

    def count_paths(room=nil, path_so_far=nil)
      room ||= @map.find_room("start")
      path_so_far ||= []

      path_so_far << room
      # puts path_so_far.map(&:name).join("-")
      
      return 1 if room.name == "end"

      # byebug if room.name=="d" && repeats==1

      room.connected_rooms
        .select { |r| room_allowed?(r, path_so_far) }
        .map { |r| count_paths(r, path_so_far.dup) }
        .sum
    end

    def room_allowed?(room, path_so_far)
      return true if room.big?
      return false if room.name == "start"

      path_names = path_so_far.select(&:small?).map(&:name)
      counts = path_names.group_by(&:itself).transform_values {|v| v.size-1 }
      duplicates = counts.values.sum
      
      duplicates < @repeats || !path_names.include?(room.name)
    end

  end
end