# frozen_string_literal: true

class DeterministicDie
  attr_reader :times_rolled
  
  def initialize(sides = 100)
    @sides = sides
    @times_rolled = 0
  end

  def roll
    @times_rolled += 1

    modded = @times_rolled % 100
    if modded.zero?
      100
    else
      modded
    end
  end
end
