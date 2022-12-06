require 'set'

module Cave
  class Input
    attr_reader :nodes, :connections

    Connection = Struct.new(:from, :to)

    def initialize(raw_data)
      @nodes = Set.new
      @connections = []
      parse(raw_data)
    end

    def parse(raw_data)
      raw_data.lines(chomp: true).each do |line|
        from, to = line.split("-")
        @nodes << from
        @nodes << to
        @connections << Connection.new(from, to)
      end
    end
  end
end