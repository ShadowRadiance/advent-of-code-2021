# frozen_string_literal: true

require "./scanner"

class ScannerReportParser
  def parse(str)
    str.lines(chomp: true)
       .slice_when { |i, _j| i.length.zero? }
       .map { |scanner_data| scanner_data.delete_if(&:empty?) }
       .map { |scanner_data| Scanner.new(scanner_data) }
  end
end
