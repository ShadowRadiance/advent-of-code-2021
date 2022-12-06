class SevenSegmentParser
  def parse(data)
    data.lines.map { |line| Line.from_string(line) }
  end

  class Line
    attr_reader :signal_patterns, :output_value

    def initialize(signal_patterns, output_value)
      @signal_patterns, @output_value = signal_patterns, output_value
    end

    def to_s
      "#{@signal_patterns.join(" ")} | #{@output_value.join(" ")}"
    end

    class << self
      def from_string(str)
        signal_patterns, output_value = str.split(' | ').map do |side|
          side.split(' ').map { |signal| signal.chars.sort.join }
        end  
        new(signal_patterns, output_value)
      end
    end
  end
end