# frozen_string_literal: true

require './snailfish_number'

numbers = File.readlines("./data/input.txt", chomp: true).map do |line|
  arr = eval(line)
  SnailfishNumber.new(arr)
end
sfn_sum = numbers.reduce { |memo, number| memo + number }
puts sfn_sum.magnitude
