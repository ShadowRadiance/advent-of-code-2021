require "./cave/room"

module Cave
  class Map
    def initialize(input)
      build_graph(input)
    end

    def build_graph(input)
      @rooms = input.nodes.map { |name| [ name, Room.new(name) ] }.to_h
      
      input.connections.each do |connection| 
        connect(@rooms[connection.from], @rooms[connection.to])
      end
    end

    def rooms
      @rooms.values
    end

    def find_room(name)
      @rooms[name]
    end

    private

    def connect(a, b)
      a.connected_rooms << b
      b.connected_rooms << a
    end
  end
end