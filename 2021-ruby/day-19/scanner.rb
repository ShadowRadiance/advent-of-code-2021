# frozen_string_literal: true

require "matrix"

class Scanner
  # x increases to the right
  # y increases to the up
  # z increases toward the viewer
  #
  # positive rotation is clockwise (looking in the direction of the axis of rotation)
  # positive rotation is anticlockwise (looking "back down the barrel" of the axis of rotation)
  #
  # rotating the scanner in the CS of the world (alias)
  #   rotates the scanner's CS in the opposite direction (alibi)
  # i.e. the beacons don't move, but appear to from the POV of the scanner

  attr_reader :identifier, :report, :beacons, :orientation
  attr_accessor :location

  def initialize(report)
    @memo = {}

    @identifier = report.shift.match(/--- scanner (\d+) ---/)[1].to_i
    @beacons = report.map { |coord| coord.split(",").map(&:to_i) }
                     .map { |beacon| Vector[*beacon] }
    @location = Vector.zero(3)
    @orientation = 0
  end

  def orientation=(index)
    @orientation = index
    @beacons = beacons_in_orientation(index)
  end

  def beacons_in_orientation(index)
    @memo[index] ||= @beacons.map { |beacon| INVERSE_ORIENTATIONS[index] * beacon }
  end

  private

  # Imagine a flattened (d6) die (cube)
  #     5
  # 6 4 1 3
  #     2
  ORIENTATIONS = [
    # 1 up
    Matrix[ [ 1, 0, 0],  [ 0, 1, 0], [ 0, 0, 1] ],  # original orientation
    Matrix[ [ 0, 0, 1],  [ 0, 1, 0], [-1, 0, 0] ],  # rotate counterclockwise 90° around y-axis
    Matrix[ [-1, 0, 0],  [ 0, 1, 0], [ 0, 0,-1] ],  # rotate counterclockwise 180° around y-axis
    Matrix[ [ 0, 0,-1],  [ 0, 1, 0], [ 1, 0, 0] ],  # rotate counterclockwise 270° around y-axis

    # 2 up
    Matrix[ [ 1, 0, 0],  [ 0, 0, 1], [ 0,-1, 0] ],
    Matrix[ [ 0,-1, 0],  [ 0, 0, 1], [-1, 0, 0] ],
    Matrix[ [-1, 0, 0],  [ 0, 0, 1], [ 0, 1, 0] ],
    Matrix[ [ 0, 1, 0],  [ 0, 0, 1], [ 1, 0, 0] ],

    # 3 up
    Matrix[ [ 0,-1, 0],  [ 1, 0, 0], [ 0, 0, 1] ],
    Matrix[ [ 0, 0, 1],  [ 1, 0, 0], [ 0, 1, 0] ],
    Matrix[ [ 0, 1, 0],  [ 1, 0, 0], [ 0, 0,-1] ],
    Matrix[ [ 0, 0,-1],  [ 1, 0, 0], [ 0,-1, 0] ],

    # 4 up
    Matrix[ [ 0, 1, 0],  [-1, 0, 0], [ 0, 0, 1] ],
    Matrix[ [ 0, 0, 1],  [-1, 0, 0], [ 0,-1, 0] ],
    Matrix[ [ 0,-1, 0],  [-1, 0, 0], [ 0, 0,-1] ],
    Matrix[ [ 0, 0,-1],  [-1, 0, 0], [ 0, 1, 0] ],

    # 5 up
    Matrix[ [ 1, 0, 0],  [ 0, 0,-1], [ 0, 1, 0] ],
    Matrix[ [ 0, 1, 0],  [ 0, 0,-1], [-1, 0, 0] ],
    Matrix[ [-1, 0, 0],  [ 0, 0,-1], [ 0,-1, 0] ],
    Matrix[ [ 0,-1, 0],  [ 0, 0,-1], [ 1, 0, 0] ],

    # 6 up
    Matrix[ [ 1, 0, 0],  [ 0,-1, 0], [ 0, 0,-1] ],
    Matrix[ [ 0, 0,-1],  [ 0,-1, 0], [-1, 0, 0] ],
    Matrix[ [-1, 0, 0],  [ 0,-1, 0], [ 0, 0, 1] ],
    Matrix[ [ 0, 0, 1],  [ 0,-1, 0], [ 1, 0, 0] ],
  ].freeze

  INVERSE_ORIENTATIONS = ORIENTATIONS.map(&:inverse).map {|m| m.map(&:to_i)}.freeze
end
