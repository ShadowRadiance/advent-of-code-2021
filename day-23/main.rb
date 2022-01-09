# frozen_string_literal: true

require "./solver"

data = File.read("./data/input.txt")
solver = Solver.new(data)
puts solver.solve


