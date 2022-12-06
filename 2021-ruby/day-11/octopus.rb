class Octopus
  attr_reader :energy_level
  attr_accessor :neighbors

  def initialize(energy_level)
    @energy_level = energy_level
    @neighbors = []
    @flashed = false
  end

  def power_up
    @energy_level += 1
  end

  def flash_cascade(initiator: false)
    # puts "Flash Cascade (#{initiator}) for #{@energy_level}"
    power_up unless initiator
    if @energy_level > 9 && !has_flashed?
      # puts "FLASH"
      @flashed = true
      @neighbors.each { |octo_neighbor| octo_neighbor.flash_cascade }
    end
  end

  def has_flashed?
    @flashed
  end

  def reset
    @flashed = false
    @energy_level = 0 if energy_level > 9
  end
end