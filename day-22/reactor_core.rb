# frozen_string_literal: true

require "set"
require "byebug"

class Cuboid
  attr_reader :left, :top, :front, :right, :bottom, :back, :on
  def initialize(left, top, front, right, bottom, back, on: true)
    @left, @top, @front, @right, @bottom, @back = left, top, front, right, bottom, back
    @on = on
  end

  def to_a
    [left, top, front, right, bottom, back, on]
  end

  def coords
    [left, top, front, right, bottom, back]
  end

  def intersection(other)
    c = Cuboid.new(
      [left, other.left].max,
      [top, other.top].max,
      [front, other.front].max,
      [right, other.right].min,
      [bottom, other.bottom].min,
      [back, other.back].min,
      on: !other.on
    )
    c.valid? ? c : nil
  end

  def valid?
    left <= right && top <= bottom && front <= back
  end

  def to_s
    "#{on ? '+++' : '---'} (#{left},#{top},#{front})->(#{right},#{bottom},#{back})"
  end

  def value
    size * (on ? 1 : -1)
  end

  def size
    (right+1-left) * (bottom+1-top) * (back+1-front)
  end

  def constrained?
    @constrained ||= coords.all? { |num| num.abs <= 50 }
  end
end

class ReactorCore
  def initialize(instructions)
    @instructions = instructions.map { |str| cuboid_from_instruction(str) }
    @cuboids = []
  end

  def on_count
    @cuboids.sum(&:value)
  end

  def reboot(init_only: false, max_steps: nil)
    instruction_set = init_only ? initialization_instructions : instructions

    instruction_set.each.with_index do |instruction, index|
      perform(instruction)
      return if max_steps && index + 1 >= max_steps
    end
  end

  attr_reader :instructions

  def initialization_instructions
    @instructions.select(&:constrained?)
  end

  def remaining_instructions
    @instructions.reject(&:constrained?)
  end

  private

  def perform(instruction)
    puts "ADDING CUBOID #{instruction}"

    intersections = @cuboids.map { |existing| instruction.intersection(existing) }.compact
    @cuboids << instruction if instruction.on
    @cuboids += intersections
  end
  
  def cuboid_from_instruction(str)
    # on x=50..82,y=-60..-54,z=-401..-197
    onoff, ranges = str.split(" ")
    x_range, y_range, z_range = ranges.split(",")     # [ "x=50..82", ... ]
      .map { |eq| eq.split("=").last }                # [ "50..82", ... ]
      .map { |r_str| r_str.split("..").map(&:to_i) }  # [ [50,82], etc ]
      .map { |r1, r2| r1..r2 }                        # [ 50..82 , etc ]

    Cuboid.new(
      x_range.begin, y_range.begin, z_range.begin, 
      x_range.end, y_range.end, z_range.end,
      on: onoff == "on"
    )
  end
end
