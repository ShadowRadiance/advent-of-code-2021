# frozen_string_literal: true

require "matrix"

class Probe
  attr_reader :velocity, :location
  def initialize(velocity_x = 0, velocity_y = 0)
    @location = Vector[0,0].freeze
    @velocity = Vector[velocity_x, velocity_y].freeze
  end

  def step
    @location = (location + velocity).freeze
    apply_force(drag)
    apply_force(gravity)

    @location
  end

  def apply_force(force)
    @velocity = (velocity + force).freeze
  end

  def drag
    return VECTORS[:zero] if @velocity[0].zero?

    if @velocity[0].negative?
      VECTORS[:drag_pos]
    else
      VECTORS[:drag_neg]
    end
  end

  def gravity
    VECTORS[:gravity]
  end

  VECTORS = {
    zero: Vector.zero(2).freeze,
    drag_pos: Vector[1, 0].freeze,
    drag_neg: Vector[-1, 0].freeze,
    gravity: Vector[0,-1].freeze,
  }.freeze
end
