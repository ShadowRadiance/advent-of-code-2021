require "./heap"
require "byebug"

class CaveNavigator
  INFINITE = Float::INFINITY

  def initialize(cave)
    @cave = cave

    graph_the_cave
    update_home
    initialize_the_heap
    imitate_dijkstra
  end

  def graph_the_cave
    @vertices = @cave.grid.map.with_index do |row, y| 
      row.map.with_index do |cell, x| 
        Vertex.new(cell, x, y) 
      end
    end
  end

  def update_home
    @home = @vertices[0][0]
    @home.distance_from_home = 0
  end

  def initialize_the_heap
    @heap = Heap.new
    @vertices.flatten.each do |vertex|
      @heap.add(vertex)
    end
    # puts "INITIALIZED HEAP HOME IS FIRST: #{@heap.to_a.first.to_s}"
  end

  def imitate_dijkstra
    until @heap.empty?
      vertex = @heap.extract
      puts "Examining (#{vertex.x}, #{vertex.y}… (#{vertex.distance_from_home})"
      if vertex.x == width-1 && vertex.y == height-1
        break
      end
      
      vertex_neighbours(vertex).each do |neighbor|
        update_vertex(neighbor, from: vertex)
      end
    end
  end

  def update_vertex(vertex, from:)
    return unless @heap.include?(vertex)

    dist = vertex.cost_to_arrive + from.distance_from_home
    if dist < vertex.distance_from_home
      vertex.distance_from_home = dist
      vertex.previous = from
      @heap.rerank(vertex)
    end
  end

  def shortest_route
    @vertices[height-1][width-1].distance_from_home
  end

  def vertex_neighbours(vertex)
    x = vertex.x
    y = vertex.y

    neighbours = []
    neighbours << @vertices[y-1][x] if y-1 >= 0     # north
    neighbours << @vertices[y][x+1] if x+1 < width  # east
    neighbours << @vertices[y+1][x] if y+1 < height # south
    neighbours << @vertices[y][x-1] if x-1 >= 0     # west

    neighbours
  end

  def width; @cave.width; end
  def height; @cave.height; end

  class Vertex
    include Comparable

    attr_accessor :distance_from_home, :previous
    attr_reader :x, :y, :cost_to_arrive

    def initialize(cost_to_arrive, x, y, distance_from_home: INFINITE)
      @cost_to_arrive = cost_to_arrive
      @x = x
      @y = y
      @distance_from_home = distance_from_home
      @previous = nil
    end

    def ==(other)
      return false unless other.is_a? Vertex
      self.equal?(other)
    end

    def <=>(other)
      distance_from_home <=> other.distance_from_home
    end

    def to_s
      "#{[x, y, distance_from_home]}"
    end
  end

end
