class LanternfishSimulator
  def initialize(fish_counters)
    @fish_counters = fish_counters
  end

  def run(days: 80)
    days.times do |day|
      process_fish
    end
  end

  def count
    @fish_counters.count
  end

  def all
    @fish_counters.dup
  end

  def process_fish
    new_fish = []
    @fish_counters.map! do |fish|
      if fish==0
        new_fish << 8
        6
      else 
        fish - 1
      end
    end
    @fish_counters += new_fish
  end
end
