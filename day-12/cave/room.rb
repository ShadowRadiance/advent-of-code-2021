require 'set'

module Cave
  class Room
    attr_reader :name, :connected_rooms

    def initialize(name)
      @name = name
      @size = name.upcase==name ? :big : :small
      @connected_rooms = Set.new
    end

    def big?
      @size == :big
    end
    
    def small?
      @size == :small
    end
  end
end