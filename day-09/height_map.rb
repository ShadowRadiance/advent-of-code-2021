class HeightMap
  attr_reader :data, :height, :width

  def initialize(height_map_2d)
    @data = height_map_2d
    @width = height_map_2d.first.length
    @height = height_map_2d.length
  end

  def self.load_string(str)
    rows = str.split("\n").map { |row| row.split(//).map(&:to_i) }
    new(rows)
  end

  def risk_level_at_point(x, y)
    1 + data[y][x]
  end

  def low_points
    data.flat_map.with_index do |row, y|
      row.map.with_index do |col, x|
        low_point?(x,y) ? [x, y] : nil
      end.compact
    end.compact
  end

  def low_point_risk_levels
    low_points.map { |a,b| risk_level_at_point(a,b) }
  end

  private

  def low_point?(x, y)
    neighbors = [
      [x,   y-1], 
      [x+1, y  ],
      [x,   y+1],
      [x-1, y  ],
    ].select { |a,b| (0...width).cover?(a) && (0...height).cover?(b) }

    neighbors.all? { |n_x, n_y| data[y][x] < data[n_y][n_x] } 
  end
end