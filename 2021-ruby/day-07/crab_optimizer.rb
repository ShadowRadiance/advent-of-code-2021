class CrabOptimizer
  attr_reader :positions

  def initialize(positions)
    @positions = positions
  end

  def optimal_alignment
    @optimal_alignment ||= all_possible_alignments.min_by { |at| cost_of_alignment(at) }
  end

  def optimal_alignment_fuel
    cost_of_alignment(optimal_alignment)
  end

  def all_possible_alignments
    a,b = positions.minmax
    a..b
  end

  def cost_of_alignment(at)
    # positions.map { |pos| (pos-at).abs }.sum
    positions.map { |pos| (pos-at).zero? ? 0 : (1..(pos-at).abs).sum }.sum
  end
end