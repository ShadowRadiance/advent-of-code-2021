# frozen_string_literal: true

require "./scanner_report_parser"
require "./scanner"
require "./recombobulator"

scanners = ScannerReportParser.new.parse(File.read("./data/input.txt"))
reco = Recombobulator.new(scanners)
reco.recombobulate

puts reco.beacons.size

def manhattan_distance(a,b)
  (a[0]-b[0]).abs + (a[1]-b[1]).abs + (a[2]-b[2]).abs
end

scanners = reco.scanners.map(&:location)

max = -Float::INFINITY
scanners.combination(2) do |a,b|
  md = manhattan_distance(a,b)
  max = md if md > max
end
puts max
