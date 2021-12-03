require "./diagnostic_report"

class App
  def initialize
    @report = DiagnosticReport.new(load_data_from_input)
  end

  def run
    puts @report.oxygen_generator_rating * @report.co2_scrubber_rating
  end

  def load_data_from_input
    File.readlines("./data/input.txt", chomp: true)
  end
end

App.new.run
