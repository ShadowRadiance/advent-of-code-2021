class LanternfishSimulator
  attr_reader :fish_counts
  def initialize(remaining_days_per_fish)
    @fish_counts = Array.new(9, 0)
    remaining_days_per_fish.each do |days|
      @fish_counts[days] += 1
    end
  end

  def run(days: 80)
    days.times do |day|
      process_fish
    end
  end

  def count
    @fish_counts.sum
  end

  def process_fish
    # 1. remember how many 0s we have and move everything else down one place in the array
    mommas = @fish_counts.delete_at(0)
    # 1. add the mommas to position 6
    @fish_counts[6] += mommas
    @fish_counts << mommas # add a baby for each momma
  end
end
