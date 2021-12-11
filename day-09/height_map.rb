class HeightMap
  attr_reader :data, :height, :width

  Point = Struct.new(:x, :y)
  class Basin
    attr_reader :members

    def initialize(members = [])
      @members = members
    end

    def size
      @members.size
    end
  end

  def initialize(height_map_2d)
    @data = height_map_2d
    @width = height_map_2d.first.length
    @height = height_map_2d.length
    @valid_x = (0...@width)
    @valid_y = (0...@height)
  end

  def self.load_string(str)
    rows = str.split("\n").map { |row| row.split(//).map(&:to_i) }
    new(rows)
  end

  def risk_level_at(location)
    1 + depth_at(location)
  end

  def low_points
    if @low_points.nil?
      @low_points = data.flat_map.with_index do |row, y|
        row.map.with_index do |col, x|
          location = Point.new(x,y)
          low_point?(location) ? location : nil
        end.compact
      end.compact
    end

    @low_points
  end

  def low_point_risk_levels
    low_points.map { |location| risk_level_at(location) }
  end

  def basins(top: nil)
    if @basins.nil?
      @basins = find_basins_by_target
        .sort_by(&:size)
        .reverse
    end

    if top.nil?
      @basins
    else
      @basins.first(top)
    end
  end

  private

  def low_point?(location)
    current_depth = depth_at(location)
    neighbors_of(location).all? { |neighbor| current_depth < depth_at(neighbor) }
  end

  def neighbors_of(location)
    x, y = location.x, location.y

    [
      Point.new(x,   y-1),
      Point.new(x+1, y  ),
      Point.new(x,   y+1),
      Point.new(x-1, y  ),
    ].select { |neighbor| valid?(neighbor) }
  end

  def depth_at(location)
    data[location.y][location.x]
  end

  def valid?(location)
    @valid_x.cover?(location.x) && @valid_y.cover?(location.y)
  end

  def find_basins_by_target
    targets = @valid_y.flat_map do |y|
      @valid_x.map do |x|
        location = Point.new(x,y)
        trickle_down(location, location)
      end.compact
    end

    targets
      .group_by { |hsh| hsh[:target] }
      .transform_values { |arr| arr.map { |hsh| hsh[:original] } }
      .map { |_target, origins| Basin.new(origins) }
  end

  def trickle_down(location, original)
    current_depth = depth_at(location)
    return nil if current_depth == 9
    deeper = neighbors_of(location).find { |neighbor| depth_at(neighbor) < current_depth }
    if deeper.nil?
      {
        original: original,
        target: location,
      }
    else
      trickle_down(deeper, original)
    end
  end
end