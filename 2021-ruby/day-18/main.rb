# frozen_string_literal: true

require './snailfish_number'

numbers = File.readlines("./data/input.txt", chomp: true).map do |line|
  arr = eval(line)
  SnailfishNumber.new(arr)
end

sfn_sum = numbers.reduce { |memo, number| memo + number }
puts "Part 1: #{sfn_sum.magnitude}"

biggest_magnitude = -Float::INFINITY
numbers.permutation(2) do |a, b|
  sum = a + b
  mag = sum.magnitude
  biggest_magnitude = mag if mag > biggest_magnitude
end
puts "Part 2: #{biggest_magnitude}"
