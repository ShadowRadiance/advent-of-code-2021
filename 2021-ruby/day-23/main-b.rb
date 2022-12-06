# frozen_string_literal: true

require "./solver"
require "./pods/room"
require "./pods/map"

lines = File.readlines("./data/input.txt", chomp: true)

data = [
  lines[0],
  lines[1],
  lines[2],
  "  #D#C#B#A#",
  "  #D#B#A#C#",
  lines[3],
  lines[4]
].join("\n")

solver = Solver.new(data, ".|.|AAAA|.|BBBB|.|CCCC|.|DDDD|.|.", room_size: 4)
puts solver.solve
