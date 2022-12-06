# frozen_string_literal: true

require "set"
require "byebug"

class SlowReactorCore
  def initialize(instructions)
    @all_instructions = instructions.map { |str| Instruction.parse(str) }
    @instructions = @all_instructions
      .group_by { |inst| inst.constrained?(-50, 50) ? "init" : "rest" }
    @switches = Set.new
  end

  def on_count
    @switches.size
  end

  def reboot(init_only: false, max_steps: nil)
    instruction_set = if init_only
      @instructions["init"]
    else
      @all_instructions
    end

    instruction_set.each.with_index do |instruction, index|
      perform(instruction)
      return if max_steps && index + 1 >= max_steps
    end
  end

  def perform(instruction)
    if TurnOnInstruction===instruction
      # @switches = @switches.union(instruction_targets) (+ or |)
      instruction.combinations.each { |combo| @switches.add(combo) }
    else
      instruction.combinations.each { |combo| @switches.delete(combo) }
      # @switches = @switches.difference(instruction_targets) (-)
    end
  end

  def instructions
    @all_instructions
  end

  def initialization_instructions
    @instructions["init"]
  end

  def remaining_instructions
    @instructions["rest"]
  end

  class Instruction
    def self.parse(str)
      # on x=50..82,y=-60..-54,z=-401..-197
      onoff, ranges = str.split(" ")
      x_range, y_range, z_range = ranges.split(",")     # [ "x=50..82", ... ]
        .map { |eq| eq.split("=").last }                # [ "50..82", ... ]
        .map { |r_str| r_str.split("..").map(&:to_i) }  # [ [50,82], etc ]
        .map { |r1, r2| r1..r2 }                        # [ 50..82 , etc ]

      if onoff == "on"
        TurnOnInstruction.new(x_range, y_range, z_range)
      else
        TurnOffInstruction.new(x_range, y_range, z_range)
      end
    end

    attr_reader :x_range, :y_range, :z_range
    def initialize(x_range, y_range, z_range)
      @x_range, @y_range, @z_range = x_range, y_range, z_range
    end
    
    def constrained?(lo, hi)
      @constrained ||= [
        x_range.begin, x_range.end, 
        y_range.begin, y_range.end, 
        z_range.begin, z_range.end
      ].all? { |num| num.abs <= 50 }
    end

    def combinations
      x_range.to_a.product(y_range.to_a, z_range.to_a)
    end
  end

  class TurnOnInstruction < Instruction; end
  class TurnOffInstruction < Instruction; end
end