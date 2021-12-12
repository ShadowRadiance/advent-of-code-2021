require "./cave/map"

module Cave
  class Solver
    def initialize(map)
      @map = map
    end

    def count_paths(room=nil, path_so_far=nil)
      room ||= @map.find_room("start")
      path_so_far ||= []

      path_so_far << room.name
      # puts path_so_far.join("-")
      
      return 1 if room.name == "end"

      room.connected_rooms
        .select { |r| r.big? || !path_so_far.include?(r.name) }
        .map { |r| count_paths(r, path_so_far.dup) }
        .sum
    end

  end
end