# frozen_string_literal: true

require './probe_launcher'

line = File.readlines("./data/input.txt", chomp: true).first
# target area: x=94..151, y=-156..-103
matches = line.match(/target area: x=(\d+)..(\d+), y=(-\d+)..(-\d+)/)
r1 = matches[1].to_i..matches[2].to_i
r2 = matches[3].to_i..matches[4].to_i
target = [r1, r2]
launcher = ProbeLauncher.new(target)
launcher.determine_highest_possible_shot
puts launcher.highest_accurate_shot

puts launcher.highest_accurate_shot[:height]
