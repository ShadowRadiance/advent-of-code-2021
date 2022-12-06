class SevenSegmentDecoder
  attr_reader :original, :decoded

  def initialize(parsed_data)
    @original = parsed_data.dup
    @decoded = parsed_data.dup
  end

  def decode
    @decoded = @original.map { |line|

      mappings = Array.new(10, nil)
      
      # 1 is the only one with 2 segments in the pattern
      mappings[1] = line.signal_patterns.find { |p| p.length==2 }
      # 7 is the only one with 3 segments in the pattern
      mappings[7] = line.signal_patterns.find { |p| p.length==3 }
      # 4 is the only one with 4 segments in the pattern
      mappings[4] = line.signal_patterns.find { |p| p.length==4 }
      # 8 is the only one with 7 segments in the pattern
      mappings[8] = line.signal_patterns.find { |p| p.length==7 }
      
      pattern_4 = mappings[4].chars
      pattern_7 = mappings[7].chars
      
      # five segment patterns (2, 3, 5)
      five_segments = line.signal_patterns.select { |p| p.length==5 }
      # remove the segments for the 7 pattern
      # leaves the 3 pattern with 2s (and the 2,5 patterns with 3 segments each)
      five_segments_sub_7 = five_segments.map { |p| p.chars.difference(pattern_7).join }
      mappings[3] = five_segments[five_segments_sub_7.index { |p| p.length == 2 }]
      # remove the segments for the 4 pattern
      # leaves the 2 pattern with 2 segments(and the 3,5 patterns with 1 segment)
      five_segments_sub_47 = five_segments_sub_7.map { |p| p.chars.difference(pattern_4).join }
      mappings[2] = five_segments[five_segments_sub_47.index { |p| p.length == 2 }]
      five_segments.delete(mappings[3])
      five_segments.delete(mappings[2])
      mappings[5] = five_segments.first

      # six segment patterns (6,9,0)
      six_segments = line.signal_patterns.select { |p| p.length==6 }
      six_segments_sub_7 = six_segments.map { |p| p.chars.difference(pattern_7).join }
      mappings[6] = six_segments[six_segments_sub_7.index { |p| p.length == 4 }]
      six_segments_sub_47 = six_segments_sub_7.map { |p| p.chars.difference(pattern_4).join }
      mappings[9] = six_segments[six_segments_sub_47.index { |p| p.length == 1 }]
      six_segments.delete(mappings[6])
      six_segments.delete(mappings[9])
      mappings[0] = six_segments.first

      # puts mappings.map.with_index {|v, i| "#{i}: #{v}" }.join(" ")

      output_value = line.output_value.map { |p| mappings.index(p) || p }

      SevenSegmentParser::Line.new(
        line.signal_patterns,
        output_value
      )
    }
  end

  def count_all(digit)
    decoded.map(&:output_value).map { |item| item.count(digit) }.sum
  end

  def decoded_integers
    decoded.map(&:output_value).map { |ov|
      if ov.all? { |el| el.is_a?(Integer) }
        ov.join.to_i
      else
        ov
      end
    }
  end

  def noop; end
end