# frozen_string_literal: true

require "./scanner_report_parser"
require "./scanner"
require "./recombobulator"

scanners = ScannerReportParser.new.parse(File.read("./data/input.txt"))
reco = Recombobulator.new(scanners)
reco.recombobulate

puts reco.beacons.size
