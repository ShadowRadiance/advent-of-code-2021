class DiagnosticReport
  attr_reader :data
  def initialize(data)
    @data = data
  end

  def gamma_rate
    return 0 if counts.empty?
    counts.map { |count| count.common.to_s }.join.to_i(2)
  end
  
  def epsilon_rate
    return 0 if counts.empty?
    counts.map { |count| count.uncommon.to_s }.join.to_i(2)
  end

  def oxygen_generator_rating
    return 0 if counts.empty?
    
    filter(:common, 1)
  end

  def co2_scrubber_rating
    return 0 if counts.empty?

    filter(:uncommon, 0)
  end

  private

  Count = Struct.new(:zeros, :ones, keyword_init: true) do
    def equal?
      zeros == ones
    end
    def common
      zeros > ones ? 0 : 1
    end
    def uncommon
      zeros < ones ? 0 : 1
    end
  end

  def filter(method, default)
    filtered = data.dup
    filtered_counts = counts

    data.first.length.times do |index|  
      target = if filtered_counts[index].equal? 
                 default 
               else
                 filtered_counts[index].public_send(method)
               end.to_s
      filtered = filtered.filter { |row| row[index] == target }
      break if filtered.size == 1
      filtered_counts = determine_counts(filtered)
    end
    filtered.first.to_i(2)
  end

  def counts
    return [] if data.empty?

    @counts ||= determine_counts(data)
  end

  def determine_counts(arr)
    # arr
    # ["010", "111", ...]

    list_of_digit_arrays = arr.map { |row| row.split(//).map(&:to_i) }
    # [ [0,1,0], [1,1,1], ... ]
    
    numbers_of_ones = list_of_digit_arrays.reduce { |sums, digits| sums.zip(digits).map(&:sum) }
    # [ 15, 26, 12, ... ]
    
    total_rows = arr.length
    # 50

    numbers_of_zeros = numbers_of_ones.map { |ones| total_rows - ones }
    # [ 35, 24, 38, ... ]

    numbers_of_zeros.zip(numbers_of_ones).map { |zeros, ones| Count.new(zeros: zeros, ones: ones) }
    # [ Count(zeros: 35, ones: 15), Count(zeros: 24, ones: 26), ... ]
  end
end