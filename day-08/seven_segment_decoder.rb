class SevenSegmentDecoder
  attr_reader :original, :decoded

  def initialize(parsed_data)
    @original = parsed_data.dup
    @decoded = parsed_data.dup
  end

  def decode
    @decoded = @original.map { |line|
      SevenSegmentParser::Line.new(
        line.signal_patterns,
        line.output_value.map do |pattern|
          case pattern.length
          when 2 then 1 # 1 is the only one with 2 signals in the pattern
          when 3 then 7 # 7 is the only one with 3 signals in the pattern
          when 4 then 4 # 4 is the only one with 4 signals in the pattern
          when 7 then 8 # 8 is the only one with 7 signals in the pattern
          else pattern
          end
        end
      )

      # output_value = line.output_value.dup
      # line.signal_patterns.each do |pattern|
      #   case pattern.size
      #   when 2 then output_value.replace(pattern, 1)
      #   when 3 then output_value.replace(pattern, 7)
      #   when 4 then output_value.replace(pattern, 4)
      #   when 7 then output_value.replace(pattern, 8)
      #   else noop
      #   end
      # end
      # SevenSegmentParser::Line.new(line.signal_patterns, output_value)
    }
  end

  def count_all(digit)
    decoded.map(&:output_value).map { |item| item.count(digit) }.sum
  end

  def noop; end
end