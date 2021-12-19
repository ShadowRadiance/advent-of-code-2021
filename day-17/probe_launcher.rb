# frozen_string_literal: true

require './probe'

class ProbeLauncher
  attr_reader :target_x_range, :target_y_range, :shots_on_target, :shots_taken

  def initialize(target)
    @target_x_range, @target_y_range = target
    @shots_on_target = []
    @shots_taken = 0
  end

  def determine_highest_possible_shot
    # since each shot will lose x-vel at 1 per tick
    # and each shot will lose y-vel at 1 per tick
    # after x-vel ticks x-vel will be zero (furthest x-distance)
    # assuming a positive y-vel, after y_vel ticks, y_vel will be 0 (max-height)
    #                            after y_vel+1 ticks y_vel will be -1 (but still at max height)
    #                            after 2*y_vel+1 ticks y_vel will be -y_vel-1 (and returned to 0 height)
    #
    # # [20..30, -10..-5] ==> [6,1-9]  [7,1-9] [8,1]
    # # best 6,9 height 45 hit at 21,-10

    x_vel_edges = *triangles.map { |upto, sum| upto if target_x_range.cover?(sum) }.compact.minmax
    fudge = 1 # 1, 2, n?? we have to fudge the far edge of the x_vel range to account
              # for shots that get farther than the box, but might hit it on the way past
    x_vel_range = Range.new(x_vel_edges.first, x_vel_edges.last + fudge)
    y_vel_range = 1..-target_y_range.begin-1

    puts "will take #{x_vel_range.size * y_vel_range.size} shots…"
    x_vel_range.each do |x_vel|
      y_vel_range.each do |y_vel|
        shoot(x_vel, y_vel)
      end
    end
    puts "took #{shots_taken} shots of which #{shots_on_target.size} were on target"
    puts shots_on_target.map { |shot| "#{shot[:initial_velocity]}->#{shot[:final_location]}" }.join("\n")
  end

  def triangles
    (1..100).map do |upto|
      [upto, (upto*(upto+1))/2 ]
    end.to_h
  end

  def shoot(x_vel, y_vel)
    probe = Probe.new(x_vel, y_vel)
    initial_velocity = probe.velocity
    hit, height, final_location = run_ticks_until_probe_hits_or_cannot(probe)
    puts "Shot (#{x_vel}, #{y_vel}) #{hit ? "hit" : "missed"} at (#{final_location})"
    if hit
      @shots_on_target << {
        initial_velocity: initial_velocity,
        height: height,
        final_location: final_location
      }
    end
    @shots_taken += 1
    hit
  end

  def highest_accurate_shot
    @shots_on_target.max_by { |shot| shot[:height] }
  end

  def run_ticks_until_probe_hits_or_cannot(probe)
    max_height = -Float::INFINITY
    loop do
      probe.step

      x, y = probe.location.to_a
      max_height = [max_height, y].max
      if @target_x_range.cover?(x) && @target_y_range.cover?(y)
        return [true, max_height, probe.location]
      end
      if x > @target_x_range.end || y < @target_y_range.begin
        return [false, nil, probe.location]
      end
    end

    [hit, max_height]
  end

  def puzzle
    highest_accurate_shot[:height]
  end
end
